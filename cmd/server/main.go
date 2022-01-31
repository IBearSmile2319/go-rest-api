package main

import (
	"fmt"
	"net/http"

	"github.com/IBearSmile2319/go-rest-api/internal/comment"
	"github.com/IBearSmile2319/go-rest-api/internal/database"
	transportHTTP "github.com/IBearSmile2319/go-rest-api/internal/transport/http"
)

// App - the struct which contains things like the pinters
// to database connections
type App struct{}

// Run - sets up our application
func (app *App) Run() error {
	fmt.Println("Setting UP our server")
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
		fmt.Println("Failed to set up server")
		return err
	}
	return nil
}

func main() {
	fmt.Println("Go Rest Api Server")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our server")
		fmt.Println(err)
	}
}
