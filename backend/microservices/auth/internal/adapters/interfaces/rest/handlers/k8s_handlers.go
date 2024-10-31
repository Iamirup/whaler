package handlers

import (
	"net/http"

	"github.com/Iamirup/whaler/backend/microservices/auth/internal/adapters/interfaces/rest"
	"github.com/gofiber/fiber/v2"
)

type KubernetesHandler struct {
	server *rest.Server
}

func (h *KubernetesHandler) liveness(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}

func (h *KubernetesHandler) readiness(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}
