package routers

import (
	"net/http"
	"os"
	"project2-microservice-go/internal/auth-service/config"
	"project2-microservice-go/internal/auth-service/service"
	"project2-microservice-go/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type IRouter interface {
	RegisterRoutes() http.Handler
}

type Router struct {
}

func NewRouter() IRouter {
	return &Router{}
}

func (r *Router) RegisterRoutes() http.Handler {
	router := gin.Default()

	// Middleware
	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		panic("FRONTEND_URL environment variable is not set")
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{frontendURL},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	jwtConfig := config.NewJWTConfig()
	jwtService := service.NewJWTService(jwtConfig)
	jwtMiddleware := middleware.NewJWTAuthMiddleware(jwtService)
	// Group routes by module
	api := router.Group("/api")
	RegisterHealthRoutes(api)              // Health module
	RegisterUserRoutes(api, jwtMiddleware) // User module (example)
	// RegisterProductRoutes(api) // Product module (example)

	return router
}
