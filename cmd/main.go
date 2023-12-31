package main

import (
	"go-htmx-fiber/handler"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	supa "github.com/nedpals/supabase-go"
)

func main() {

	app := echo.New()

	// Load environment variables and create Supabase client
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")
	supabaseClient := supa.CreateClient(supabaseUrl, supabaseKey)

	// Create dependencies
	deps := &handler.AppDependencies{
		Supabase: supabaseClient,
		// Add other dependencies as needed
	}
	userHandler := handler.NewUserHandler(deps)

	// userHandler := handler.UserHandler{}

	app.GET("/sign-up", userHandler.HandleUserSignUpShow)
	app.POST("/sign-up", userHandler.HandleUserSignUpNew)

	app.Logger.Fatal(app.Start("localhost:1323"))

}
