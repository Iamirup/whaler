package http

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/Iamirup/whaler/backend/eventor/internal/models"
	"github.com/Iamirup/whaler/backend/eventor/pkg/rdbms"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

type Content struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type Event struct {
	Contents []Content `json:"contents"`
	Time     string    `json:"time"`
}

func generateContents() []Content {
	var contents []Content
	for i := 1; i <= 10; i++ {
		contents = append(contents, Content{
			ID:    i,
			Title: fmt.Sprintf("Content Title %d", rand.Intn(100)),
		})
	}
	return contents
}

func (handler *Server) liveness(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}

func (handler *Server) readiness(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}

func (handler *Server) events(c *fiber.Ctx) error {

	ctx := c.Context()

	ctx.SetContentType("text/event-stream")
	ctx.Response.Header.Set("Cache-Control", "no-cache")
	ctx.Response.Header.Set("Connection", "keep-alive")
	ctx.Response.Header.Set("Transfer-Encoding", "chunked")
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
	ctx.Response.Header.Set("Access-Control-Allow-Headers", "Cache-Control")
	ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
	ctx.SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
		for {
			event := Event{
				Contents: generateContents(),
				Time:     time.Now().Format(time.RFC3339),
			}
			data, _ := json.Marshal(event)
			msg := fmt.Sprintf("data: %s\n\n", data)
			fmt.Fprintf(w, "%s", msg)
			w.Flush()
			time.Sleep(1 * time.Second)
		}
	}))
	return nil
}

func (handler *Server) getConfig(c *fiber.Ctx) error {
	userId, ok := c.Locals("user-id").(string)
	if !ok || userId == "" {
		handler.logger.Error("Invalid user-id local")
		return c.SendStatus(http.StatusInternalServerError)
	}

	configId, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil || configId == 0 {
		handler.logger.Error("Invalid token header", zap.Error(err))
		response := map[string]string{"error": "Invalid config id in path parameters"}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	config, err := handler.repository.GetConfigById(userId, configId)
	if err != nil {
		if err.Error() == rdbms.ErrNotFound {
			response := map[string]string{"error": fmt.Sprintf("The given config id (%d) doesn't exists", configId)}
			return c.Status(http.StatusBadRequest).JSON(response)
		}

		handler.logger.Error("Error happened while getting the config", zap.Any("config", config), zap.Error(err))
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.Status(http.StatusOK).JSON(config)
}

func (handler *Server) updateConfig(c *fiber.Ctx) error {
	userId, ok := c.Locals("user-id").(string)
	if !ok || userId == "" {
		handler.logger.Error("Invalid user-id local")
		return c.SendStatus(http.StatusInternalServerError)
	}

	configId, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil || configId == 0 {
		handler.logger.Error("Invalid token header", zap.Error(err))
		response := map[string]string{"error": "Invalid config id in path parameters"}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	oldConfig, err := handler.repository.GetConfigById(userId, configId)
	if err != nil {
		if err.Error() == rdbms.ErrNotFound {
			response := map[string]string{"error": fmt.Sprintf("The given config id (%d) doesn't exists", configId)}
			return c.Status(http.StatusBadRequest).JSON(response)
		}

		errString := "Error happened while getting the config"
		handler.logger.Error(errString, zap.String("user-id", userId), zap.Uint64("config-id", configId), zap.Error(err))
		return c.SendStatus(http.StatusInternalServerError)
	}

	newConfig := &models.UserConfig{}
	if err := c.BodyParser(newConfig); err != nil {
		handler.logger.Error("Error parsing request body", zap.Any("config", newConfig), zap.Error(err))
		response := map[string]string{"error": "Error parsing request body"}
		return c.Status(http.StatusBadRequest).JSON(response)
	}
	newConfig.Update(oldConfig)

	if err := handler.repository.UpdateConfig(userId, newConfig); err != nil {
		handler.logger.Error("Error happened while creating the config", zap.Any("config", newConfig), zap.Error(err))
		response := map[string]string{"error": "Error happened while creating the config"}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	response := map[string]string{"message": "Config has been updated successfully"}
	return c.Status(http.StatusOK).JSON(response)
}
