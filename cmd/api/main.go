package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	http_delivery "github.com/devdavidalonso/ecs-manager-bff/internal/delivery/http"
	firestore_repo "github.com/devdavidalonso/ecs-manager-bff/internal/repository/firestore"
	"github.com/devdavidalonso/ecs-manager-bff/internal/usecase"
)

func main() {
	ctx := context.Background()

	// Initialize Firebase application
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Fatalf("failed to initialize firebase app: %v", err)
	}

	// Initialize Firestore client
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("failed to initialize firestore client: %v", err)
	}
	defer client.Close()

	// Dependency injection wiring
	studentRepo := firestore_repo.NewStudentRepository(client)
	studentUC := usecase.NewStudentUsecase(studentRepo)

	mux := http.NewServeMux()

	// Register HTTP delivery handlers
	http_delivery.NewStudentHandler(mux, studentUC)

	// Health check endpoint
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	log.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
