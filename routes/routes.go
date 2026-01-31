package routes

import (
	"fenmo-ai-assignment/database"
	"fenmo-ai-assignment/handler"
	"fenmo-ai-assignment/middleware"
	"fenmo-ai-assignment/repository"
	"fenmo-ai-assignment/service"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all routes
func SetupRoutes() *gin.Engine {
	// Create repository
	expenseRepo := repository.NewExpenseRepository(database.DB)

	// Create service
	expenseService := service.NewExpenseService(expenseRepo)

	// Create handler
	expenseHandler := handler.NewExpenseHandler(expenseService)

	// Setup router
	router := gin.Default()

	// Apply middleware
	router.Use(middleware.Logger())
	router.Use(middleware.CORS())
	router.Use(middleware.ErrorHandler())

	// API routes
	api := router.Group("/api")
	{
		api.POST("/expenses", expenseHandler.CreateExpense)
		api.GET("/expenses", expenseHandler.GetExpenses)
	}

	// Serve frontend
	router.StaticFile("/", "./frontend/index.html")
	router.Static("/static", "./frontend/static")

	return router
}
