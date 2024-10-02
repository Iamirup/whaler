package routes

import (
	"main/api"
	"main/db"
	JWT "main/jwt"
	"main/middleware"
	"main/repository"
	"main/service"
	"main/validators"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	//db
	postDb  *gorm.DB      = db.ConnectPostgres()
	redisDb *redis.Client = db.ConnectRedis()

	//auth
	authRepository repository.AuthRepository = repository.NewAuthRepository(postDb, redisDb)
	authService    service.AuthService       = service.NewAuthService(authRepository)
	authAPI        api.AuthAPI               = api.NewAuthAPI(authService)

	//user
	userRepository repository.UserRepository = repository.NewUserRepository(postDb, redisDb)
	userService    service.UserService       = service.NewUserService(userRepository)
	userAPI        api.UserAPI               = api.NewUserAPI(userService)

	//jwt
	jwtAuth JWT.Jwt
)

func Urls() *gin.Engine {
	router := gin.Default()
	//middlewares
	router.Use(middleware.CORSMiddleware())
	router.NoRoute(middleware.NoRouteHandler())
	router.HandleMethodNotAllowed = true
	router.NoMethod(middleware.NoMethodHandler())
	//register validations
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("username_validator", validators.UsernameValidate); err != nil {
			panic("validator issue in URLS")
		}
	}

	apiV1 := router.Group("api/v1")

	//Authentication API
	auth := apiV1.Group("/auth", middleware.NotAuthorization())

	auth.POST("", authAPI.Register)
	auth.PUT("", authAPI.Login, middleware.LoginAttemptCheck())

	//Config API
	user := apiV1.Group("/configs", middleware.AuthorizationJWT(jwtAuth))

	user.GET("/get-config/:id", userAPI.GetAllUsers)
	user.PUT("/update-config/:id", userAPI.ProfileUpdate)

	return router
}
