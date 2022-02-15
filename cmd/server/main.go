package main

import (
	"net/http"

	"github.com/TutorialEdge/go-rest-api-course/internal/comment"
	"github.com/TutorialEdge/go-rest-api-course/internal/database"
	transportHTTP "github.com/TutorialEdge/go-rest-api-course/internal/transport/http"

	log "github.com/sirupsen/logrus"
)

// Run - sets up our application
func Run() error {
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("Setting Up Our APP")

	var err error
	store, err := database.NewDatabase()
	if err != nil {
		return err
	}
	err = store.MigrateDB()
	if err != nil {
		log.Error("failed to setup database")
		return err
	}

	commentService := comment.NewService(store)
	handler := transportHTTP.NewHandler(commentService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		log.Error("Failed to set up server")
		return err
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Error(err)
		log.Fatal("Error starting up our REST API")
	}
}
