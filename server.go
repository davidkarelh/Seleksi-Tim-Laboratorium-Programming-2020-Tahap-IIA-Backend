package main

import (
	"os"

	"github.com/davidkarelh/Seleksi-Tim-Laboratorium-Programming-2020-Tahap-IIA-Backend/config"
	"github.com/davidkarelh/Seleksi-Tim-Laboratorium-Programming-2020-Tahap-IIA-Backend/controllers"
	"github.com/davidkarelh/Seleksi-Tim-Laboratorium-Programming-2020-Tahap-IIA-Backend/middlewares"
	"github.com/davidkarelh/Seleksi-Tim-Laboratorium-Programming-2020-Tahap-IIA-Backend/repository"
	"github.com/davidkarelh/Seleksi-Tim-Laboratorium-Programming-2020-Tahap-IIA-Backend/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()

	historyRepository              repository.IHistoryRepository              = repository.NewHistoryRepository(db)
	userRepository                 repository.IUserRepository                 = repository.NewUserRepository(db)
	registerVerificationRepository repository.IRegisterVerificationRepository = repository.NewRegisterVerificationRepository(db)

	jwtService                  services.IJWTService                  = services.NewJWTService()
	userService                 services.IUserService                 = services.NewUserService(userRepository)
	authService                 services.IAuthService                 = services.NewAuthService(userRepository, registerVerificationRepository)
	historyService              services.IHistoryService              = services.NewHistoryService(historyRepository)
	registerVerificationService services.IRegisterVerificationService = services.NewRegisterVerificationService(registerVerificationRepository)

	historyController              controllers.IHistoryController              = controllers.NewHistoryController(historyService)
	authController                 controllers.IAuthController                 = controllers.NewAuthController(jwtService, userService, authService)
	userController                 controllers.IUserController                 = controllers.NewUserController(userService, jwtService)
	registerVerificationController controllers.IRegisterVerificationController = controllers.NewRegisterVerificationController(registerVerificationService)
)

func main() {
	defer config.CloseDatabaseConnection(db)

	server := gin.New()
	server.Use(gin.Recovery(), gin.Logger(), middlewares.CORSMiddleware())

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "ok",
		})
	})

	server.POST("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "ok",
		})
	})

	authRoutes := server.Group("/api/auth")
	{
		authRoutes.POST("/login", func(ctx *gin.Context) { authController.Login(ctx) })
		authRoutes.POST("/register", func(ctx *gin.Context) { authController.Register(ctx) })
	}

	// JWT Authorization Middleware applies to non auth "/api" only.
	apiRoutes := server.Group("/api", middlewares.AuthorizeJWT(jwtService))
	{
		apiRoutes.GET("/profile", func(ctx *gin.Context) { userController.Profile(ctx) })
		apiRoutes.GET("/register-verification", func(ctx *gin.Context) { registerVerificationController.GetAllRegisterVerifications(ctx) })
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	server.Run(":" + port)
}
