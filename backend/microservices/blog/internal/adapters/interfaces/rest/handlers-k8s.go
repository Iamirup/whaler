package rest

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type KubernetesHandler struct {
	server *restServer
}

func NewKubernetesHandler(server *restServer) *KubernetesHandler {
	return &KubernetesHandler{server: server}
}

func (h *KubernetesHandler) Liveness(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}

func (h *KubernetesHandler) Readiness(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}
