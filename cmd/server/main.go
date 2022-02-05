package main

import (
	"net/http"

	"github.com/IBearSmile2319/go-rest-api/internal/comment"
	"github.com/IBearSmile2319/go-rest-api/internal/database"
	transportHTTP "github.com/IBearSmile2319/go-rest-api/internal/transport/http"

	log "github.com/sirupsen/logrus"
)

// App - contain application information
type App struct {
	Name    string
	Version string
}

// Run - sets up our application
func (app *App) Run() error {
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(
		log.Fields{
			"AppName": app.Name,
			"Version": app.Version,
		}).Info("Starting up our server")
	// fmt.Println("Setting UP our server")
	var err error
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}
	err = database.MigrateDB(db)
	if err != nil {
		return err
	}
	commentService := comment.NewService(db)

	// Setup our routes
	handler := transportHTTP.NewHandler(commentService)
	handler.SetupRoutes()
	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		log.Error("Failed to set up server")
		return err
	}
	return nil
}

func main() {
	app := App{
		Name:    "Go Rest Api Server",
		Version: "0.0.1",
	}
	if err := app.Run(); err != nil {
		log.Error("Error starting up our server")
		log.Fatal(err)
	}
}
