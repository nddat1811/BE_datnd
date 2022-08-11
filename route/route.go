package route

import (
	"BE_datnd/middleware"
	"BE_datnd/model"
	"BE_datnd/service"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"BE_datnd/controller"
	"gorm.io/gorm"
)

func SetupRouter(conn *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("CORS_ALLOWED_ORIGINS")},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// m := model.NewUserModel(conn)
	client := router.Group("/api")
	{
		auth := controller.NewAuthController(model.NewUserModel(conn), service.NewAuthService())
		client.POST("/login", auth.Login)
	}

	var authMiddleware middleware.Authorized
	client.Use(authMiddleware.TokenAuthMiddleware())
	{
		// user := controller.NewUserController(m)

	}
	return router
}
