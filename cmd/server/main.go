package main

import (
	"fmt"
)

// App - the struct which contains things like the pinters
// to database connections
type App struct {
}

// Run - sets up our application
func (app *App) Run() error {
	fmt.Println("Setting UP our server")
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
