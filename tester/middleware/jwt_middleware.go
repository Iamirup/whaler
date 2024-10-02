package middleware

import (
	"GinChat/pkg/JWT"
	"GinChat/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthorizationJWT(jwtService JWT.JwtService) gin.HandlerFunc {
	return func(request *gin.Context) {
		var authToken string = request.GetHeader("Authorization")
		if authToken == "" {
			request.AbortWithStatusJSON(http.StatusUnauthorized, utils.AuthenticationRequired)
			return
		}
		if strings.HasPrefix(authToken, "AccessToken") {
			authToken = authToken[11:]
		}
		user, err := jwtService.ValidateToken(authToken)
		if err != nil {
			request.AbortWithStatusJSON(http.StatusBadRequest, utils.TokenIsExpiredOrInvalid)
			return
		}
		request.Set("phoneNo", user.Phone.PhoneNo)
		return
	}
}

func NotAuthorization() gin.HandlerFunc {
	return func(request *gin.Context) {
		var authToken string = request.GetHeader("Authorization")
		if authToken != "" {
			request.AbortWithStatusJSON(http.StatusBadRequest, utils.MustNotAuthenticated)
		}
	}
}
