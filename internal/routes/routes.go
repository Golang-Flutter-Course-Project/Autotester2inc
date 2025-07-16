package routes

import (
	"Autotester/configs"
	_ "Autotester/docs" // путь к сгенерированным swagger docs
	"Autotester/internal/auth"
	"Autotester/internal/handlers"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/gorm"
)

type RoutesHandlerDeps struct {
	Config *configs.Config
	DB     *gorm.DB
}

func SetupAuthRoutes(router *mux.Router, deps *RoutesHandlerDeps) {
	// Initialize auth handler
	authHandler := auth.NewAuthHandler(deps.DB, deps.Config.JWTSecret)

	// Auth routes
	router.HandleFunc("/auth/register", authHandler.Register).Methods("POST")
	router.HandleFunc("/auth/login", authHandler.Login).Methods("POST")
	router.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	// Initialize handlers
	scanHandler := handlers.NewCheckUrlHandler(deps.Config)
	testsHandler := handlers.NewTestsHandler(deps.Config)
	resultsHandler := handlers.NewResultHandler(deps.Config)

	// Protected routes with middleware
	protected := router.PathPrefix("/api").Subrouter()
	protected.Use(auth.NewJWTMiddleware(deps.Config.JWTSecret))
	protected.HandleFunc("/checkurl", scanHandler.Check).Methods("POST")
	protected.HandleFunc("/tests", testsHandler.Tests).Methods("POST")
	protected.HandleFunc("/results", resultsHandler.Results).Methods("POST")

	// Public route
	router.HandleFunc("/api/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	}).Methods("GET")
}
