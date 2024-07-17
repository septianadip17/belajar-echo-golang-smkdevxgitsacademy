package main

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "Hello, World!"})
	})

	e.POST("/user", func(c echo.Context) error {
		var request User
		if err := c.Bind(&request); err != nil {
			return err
		}

		newUser := new(User)
		newUser.ID = uuid.New()
		newUser.Name = request.Name
		newUser.Email = request.Email
		newUser.Password = request.Password

		return c.JSON(200, map[string]interface{}{
			"message": "successfully create a user",
			"data":    newUser,
		})
	})

	e.Logger.Fatal(e.Start(":8080"))
}
