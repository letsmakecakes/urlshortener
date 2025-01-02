package main

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"urlshortener/internal/api/handlers"
	"urlshortener/internal/api/middleware"
	"urlshortener/internal/api/routes"
	"urlshortener/internal/config"
	"urlshortener/internal/pkg/database"
	"urlshortener/internal/pkg/service"
	"urlshortener/pkg/logger"
)

func main() {
	// Load configuration
	cfg, err := loadConfiguration()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize logger
	if err := initializeLogger(); err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	zapLogger := logger.GetLogger()

	// Setup database connection
	db, err := setupDatabase(cfg)
	if err != nil {
		zapLogger.Fatal("Failed to connect database", zap.Error(err))
	}

	// Initialize dependencies
	urlHandler := initializeHandlers(db)

	// Setup and start the server
	startServer(cfg, urlHandler, zapLogger)
}

// loadConfiguration loads the application configuration
func loadConfiguration() (*config.Config, error) {
	return config.LoadConfig()
}

// initializeLogger sets up the logger based on the environment
func initializeLogger() error {
	return logger.Initialize(os.Getenv("ENVIRONMENT"))
}

// setupDatabase initializes the database connection
func setupDatabase(cfg *config.Config) (*database.MongoDB, error) {
	return database.NewMongoDB(cfg.MongoURI, cfg.MongoDB)
}

// initializeHandlers sets up the repository, service and handlers
func initializeHandlers(db *database.MongoDB) *handlers.URLHandler {
	urlRepo := database.NewMongoURLRepository(db)
	urlService := service.NewURLService(urlRepo)
	return handlers.NewURLHandler(urlService)
}

// startServer configures the router and starts the HTTP server
func startServer(cfg *config.Config, urlHandler *handlers.URLHandler, zapLogger *zap.Logger) {
	router := mux.NewRouter()

	// Add logging middleware
	router.Use(middleware.LoggingMiddleware)

	// Setup routes
	routes.SetupRoutes(router, urlHandler)

	// Start server
	zapLogger.Info("Server starting", zap.String("address", cfg.ServerAddress))
	if err := http.ListenAndServe(cfg.ServerAddress, router); err != nil {
		zapLogger.Fatal("Server failed to start", zap.Error(err))
	}
}
