package rest

import (
	"encoding/json"

	"github.com/Iamirup/whaler/backend/microservices/blog/internal/core/application/ports"
	appService "github.com/Iamirup/whaler/backend/microservices/blog/internal/core/application/services"
	domainService "github.com/Iamirup/whaler/backend/microservices/blog/internal/core/domain/services"
	"github.com/Iamirup/whaler/backend/microservices/blog/pkg/token"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Server struct {
	Logger *zap.Logger
	Token  token.Token

	managmentApp *fiber.App
	clientApp    *fiber.App
}

func New(log *zap.Logger, articleRepo ports.ArticlePersistencePort, token token.Token) *Server {
	server := &Server{Logger: log, Token: token}

	server.managmentApp = fiber.New(fiber.Config{JSONEncoder: json.Marshal, JSONDecoder: json.Unmarshal})
	server.clientApp = fiber.New(fiber.Config{JSONEncoder: json.Marshal, JSONDecoder: json.Unmarshal})

	kubernetesHandler := NewKubernetesHandler(server)

	server.managmentApp.Get("/healthz/liveness", kubernetesHandler.Liveness)
	server.managmentApp.Get("/healthz/readiness", kubernetesHandler.Readiness)

	articleService := domainService.NewArticleService(articleRepo, log, token)
	articleHandler := NewArticleHandler(server, appService.NewArticleApplicationService(articleService, log))

	blogV1 := server.clientApp.Group("/api/blog/v1", articleHandler.fetchUserDataMiddleware)

	blogV1.Get("/article/:url_path", articleHandler.GetAnArticle)
	blogV1.Get("/articles", articleHandler.GetArticles)
	blogV1.Post("/article", articleHandler.NewArticle)
	blogV1.Patch("/article", articleHandler.UpdateArticle)
	blogV1.Delete("/article", articleHandler.DeleteArticle)

	return server
}

func (handler *Server) Serve() {
	go func() {
		err := handler.managmentApp.Listen(":8081")
		handler.Logger.Fatal("error resolving managment server", zap.Error(err))
	}()

	go func() {
		err := handler.clientApp.Listen(":8080")
		handler.Logger.Fatal("error resolving client server", zap.Error(err))
	}()
}
