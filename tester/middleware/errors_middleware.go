package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NoMethodHandler() gin.HandlerFunc {
	return func(request *gin.Context) {
		request.JSON(http.StatusMethodNotAllowed, utils.MethodNotAllowed)
	}
}

func NoRouteHandler() gin.HandlerFunc {
	return func(request *gin.Context) {
		request.JSON(http.StatusNotFound, utils.RouteNotDefined)
	}
}
