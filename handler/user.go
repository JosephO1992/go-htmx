package handler

import (
	"context"
	"go-htmx-fiber/model"
	"go-htmx-fiber/views/user"

	"github.com/labstack/echo/v4"
	supa "github.com/nedpals/supabase-go"
)

type AppDependencies struct {
	Supabase *supa.Client
	// Add other dependencies as needed
}

type UserHandler struct {
	Dependencies *AppDependencies
}

func NewUserHandler(deps *AppDependencies) *UserHandler {
	return &UserHandler{Dependencies: deps}
}

func (h UserHandler) HandleUserSignUpShow(c echo.Context) error {
	return render(c, user.SignUpForm())
}

// func (h UserHandler) HandleUserSignUpNew(c echo.Context) error {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	supabaseUrl := os.Getenv("SUPABASE_URL")
// 	supabaseKey := os.Getenv("SUPABASE_KEY")
// 	supabase := supa.CreateClient(supabaseUrl, supabaseKey)

// 	ctx := context.Background()
// 	email := c.Request().PostFormValue("email")
// 	password := c.Request().PostFormValue("password")
// 	u, err := supabase.Auth.SignUp(ctx, supa.UserCredentials{
// 		Email:    email,
// 		Password: password,
// 	})
// 	if err != nil {
// 		panic(err)
// 	}

// 	newUser := model.User{Email: u.Email}

// 	return render(c, user.Show(newUser))

// }

func (h UserHandler) HandleUserSignUpNew(c echo.Context) error {
	ctx := context.Background()
	email := c.Request().PostFormValue("email")
	password := c.Request().PostFormValue("password")

	u, err := h.Dependencies.Supabase.Auth.SignUp(ctx, supa.UserCredentials{
		Email:    email,
		Password: password,
	})
	if err != nil {
		panic(err)
	}

	newUser := model.User{Email: u.Email}

	return render(c, user.Show(newUser))

}
