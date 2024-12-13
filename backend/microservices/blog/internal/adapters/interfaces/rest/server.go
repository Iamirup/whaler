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

type restServer struct {
	Logger *zap.Logger
	Token  token.Token

	managmentApp *fiber.App
	clientApp    *fiber.App
}

func NewRestServer(log *zap.Logger, articleRepo ports.ArticlePersistencePort, token token.Token) *restServer {
	restServer := &restServer{Logger: log, Token: token}

	restServer.managmentApp = fiber.New(fiber.Config{JSONEncoder: json.Marshal, JSONDecoder: json.Unmarshal})
	restServer.clientApp = fiber.New(fiber.Config{JSONEncoder: json.Marshal, JSONDecoder: json.Unmarshal})

	kubernetesHandler := NewKubernetesHandler(restServer)

	restServer.managmentApp.Get("/healthz/liveness", kubernetesHandler.Liveness)
	restServer.managmentApp.Get("/healthz/readiness", kubernetesHandler.Readiness)

	articleService := domainService.NewArticleService(articleRepo, log, token)
	articleHandler := NewArticleHandler(restServer, appService.NewArticleApplicationService(articleService, log))

	blogV1 := restServer.clientApp.Group("/api/blog/v1", articleHandler.fetchUserDataMiddleware)

	blogV1.Get("/article/:articleId", articleHandler.GetAnArticle)
	blogV1.Get("/all-articles", articleHandler.GetAllArticles)
	blogV1.Get("/my-articles", articleHandler.GetMyArticles)
	blogV1.Post("/article", articleHandler.NewArticle)
	blogV1.Patch("/article", articleHandler.UpdateArticle)
	blogV1.Delete("/article", articleHandler.DeleteArticle)

	return restServer
}

func (handler *restServer) Serve() {
	go func() {
		err := handler.managmentApp.Listen(":8081")
		handler.Logger.Fatal("error resolving managment server", zap.Error(err))
	}()

	go func() {
		err := handler.clientApp.Listen(":8080")
		handler.Logger.Fatal("error resolving client server", zap.Error(err))
	}()
}
