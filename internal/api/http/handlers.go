package http

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/Iamirup/whaler/internal/models"
	"github.com/Iamirup/whaler/pkg/rdbms"
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

func (handler *Server) register(c *fiber.Ctx) error {
	request := struct{ Username, Password string }{}
	if err := c.BodyParser(&request); err != nil {
		errString := "Error parsing request body"
		handler.logger.Error(errString, zap.Any("request", request), zap.Error(err))
		return c.Status(http.StatusBadRequest).SendString(errString)
	}

	user, err := handler.repository.GetUserByUsername(request.Username)
	if err != nil && err.Error() != rdbms.ErrNotFound {
		errString := "Error while retrieving data from database"
		handler.logger.Error(errString, zap.Error(err))
		return c.Status(http.StatusInternalServerError).SendString(errString)
	} else if err == nil || (user != nil && user.Id != 0) {
		errString := "User with given username already exists"
		handler.logger.Error(errString, zap.String("username", request.Username))
		return c.Status(http.StatusBadRequest).SendString(errString)
	}

	user = &models.User{Username: request.Username, Password: request.Password}
	if err := handler.repository.CreateUser(user); err != nil {
		errString := "Error happened while creating the user"
		handler.logger.Error(errString, zap.Error(err))
		return c.Status(http.StatusInternalServerError).SendString(errString)
	} else if user.Id == 0 {
		errString := "Error invalid user id created"
		handler.logger.Error(errString, zap.Any("user", user))
		return c.Status(http.StatusInternalServerError).SendString(errString)
	}

	accessToken, err := handler.token.CreateTokenString(user.Id)
	if err != nil {
		errString := "Error creating JWT access token for user"
		handler.logger.Error(errString, zap.Any("user", user), zap.Error(err))
		return c.Status(http.StatusInternalServerError).SendString(errString)
	}

	refreshToken, err := handler.token.CreateRefreshTokenString(user.Id)
	if err != nil {
		errString := "Error creating JWT refresh token for user"
		handler.logger.Error(errString, zap.Any("user", user), zap.Error(err))
		return c.Status(http.StatusInternalServerError).SendString(errString)
	}

	// Set refresh token as HTTP-only cookie
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(handler.token.GetRefreshTokenExpiration()),
		HTTPOnly: true,
	})

	response := map[string]string{"AccessToken": accessToken}
	return c.Status(http.StatusCreated).JSON(&response)
}

func (handler *Server) login(c *fiber.Ctx) error {
	request := struct{ Username, Password string }{}

	if err := c.BodyParser(&request); err != nil {
		errString := "Error parsing request body"
		handler.logger.Error(errString, zap.Error(err))
		return c.Status(http.StatusBadRequest).SendString(errString)
	}

	user, err := handler.repository.GetUserByUsernameAndPassword(request.Username, request.Password)
	if err != nil {
		errString := "Wrong username or password has been given"
		handler.logger.Error(errString, zap.Error(err))
		return c.Status(http.StatusBadRequest).SendString(errString)
	} else if user == nil {
		errString := "Error invalid user returned"
		handler.logger.Error(errString, zap.Any("request", request))
		return c.Status(http.StatusInternalServerError).SendString(errString)
	}

	accessToken, err := handler.token.CreateTokenString(user.Id)
	if err != nil {
		errString := "Error creating JWT access token for user"
		handler.logger.Error(errString, zap.Any("user", user), zap.Error(err))
		return c.Status(http.StatusInternalServerError).SendString(errString)
	}

	refreshToken, err := handler.token.CreateRefreshTokenString(user.Id)
	if err != nil {
		errString := "Error creating JWT refresh token for user"
		handler.logger.Error(errString, zap.Any("user", user), zap.Error(err))
		return c.Status(http.StatusInternalServerError).SendString(errString)
	}

	// Set refresh token as HTTP-only cookie
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(handler.token.GetRefreshTokenExpiration()),
		HTTPOnly: true,
	})

	response := map[string]string{"AccessToken": accessToken}
	return c.Status(http.StatusOK).JSON(&response)
}

func (handler *Server) getConfig(c *fiber.Ctx) error {
	userId, ok := c.Locals("user-id").(uint64)
	if !ok || userId == 0 {
		handler.logger.Error("Invalid user-id local")
		return c.SendStatus(http.StatusInternalServerError)
	}

	configId, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil || configId == 0 {
		handler.logger.Error("Invalid token header", zap.Error(err))
		response := "Invalid config id in path parameters"
		return c.Status(http.StatusBadRequest).SendString(response)
	}

	config, err := handler.repository.GetConfigById(userId, configId)
	if err != nil {
		if err.Error() == rdbms.ErrNotFound {
			response := fmt.Sprintf("The given config id (%d) doesn't exists", configId)
			return c.Status(http.StatusBadRequest).SendString(response)
		}

		errString := "Error happened while getting the config"
		handler.logger.Error(errString, zap.Any("config", config), zap.Error(err))
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.Status(http.StatusOK).JSON(config)
}

func (handler *Server) updateConfig(c *fiber.Ctx) error {
	userId, ok := c.Locals("user-id").(uint64)
	if !ok || userId == 0 {
		handler.logger.Error("Invalid user-id local")
		return c.SendStatus(http.StatusInternalServerError)
	}

	configId, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil || configId == 0 {
		handler.logger.Error("Invalid token header", zap.Error(err))
		response := "Invalid config id in path parameters"
		return c.Status(http.StatusBadRequest).SendString(response)
	}

	oldConfig, err := handler.repository.GetConfigById(userId, configId)
	if err != nil {
		if err.Error() == rdbms.ErrNotFound {
			response := fmt.Sprintf("The given config id (%d) doesn't exists", configId)
			return c.Status(http.StatusBadRequest).SendString(response)
		}

		errString := "Error happened while getting the config"
		handler.logger.Error(errString, zap.Uint64("user-id", userId), zap.Uint64("config-id", configId), zap.Error(err))
		return c.SendStatus(http.StatusInternalServerError)
	}

	newConfig := &models.UserConfig{}
	if err := c.BodyParser(newConfig); err != nil {
		errString := "Error parsing request body"
		handler.logger.Error(errString, zap.Any("config", newConfig), zap.Error(err))
		return c.Status(http.StatusBadRequest).SendString(errString)
	}
	newConfig.Update(oldConfig)

	if err := handler.repository.UpdateConfig(userId, newConfig); err != nil {
		errString := "Error happened while creating the config"
		handler.logger.Error(errString, zap.Any("config", newConfig), zap.Error(err))
		return c.SendStatus(http.StatusInternalServerError)
	}

	response := "Config has been updated successfully"
	return c.Status(http.StatusOK).SendString(response)
}
