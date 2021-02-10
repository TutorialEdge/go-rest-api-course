package main

import (
	"fmt"
	"net/http"

	"github.com/TutorialEdge/go-rest-api-course/internal/comment"
	"github.com/TutorialEdge/go-rest-api-course/internal/database"
	transportHTTP "github.com/TutorialEdge/go-rest-api-course/internal/transport/http"
)

// App -
type App struct{}

// Run - sets up our application
func (app *App) Run() error {
	fmt.Println("Setting Up Our APP")

	// set up our new database connection
	var err error
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}

	if err = database.MigrateDB(db); err != nil {
		return err
	}

	// create out comment service which needs a connection
	// to the database
	commentService := comment.NewService(db)

	// expose our comment service by wrapping it in a http router
	handler := transportHTTP.NewHandler(commentService)
	handler.SetupRoutes()

	// finally, start listening on port 8080 and handle the errors
	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}
	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}
}
