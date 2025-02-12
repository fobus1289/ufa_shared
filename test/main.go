package main

import (
	"errors"

	"github.com/fobus1289/ufa_shared/jwtService"
	"github.com/fobus1289/ufa_shared/middleware"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type User struct {
	Id   int64  `form:"id"`
	Name string `form:"name"`
	Age  int    `form:"age"`
	Role string `form:"role"`
}

func (u *User) ID() int64 {
	return u.Id
}

func (u *User) Pre(_ echo.Context, _ *gorm.DB, _ ...string) error {
	return errors.New("Forbidden")
}

// ID() int64
// Pre(T, E, ...K) error

func FormValue(key string) string {
	m := map[string]string{
		"id":   "1",
		"name": "user 1",
		"age":  "18",
		"role": "admin",
	}

	return m[key]
}

func main() {

	c := jwtService.JwtConfig{
		Secret:  "1234",
		Expired: 15,
	}

	authMiddleware := middleware.NewAuthEchoMiddleware[*User](
		c,
		nil,
		nil,
	)

	e := echo.New()

	e.POST("/", func(c echo.Context) error {
		token, _ := authMiddleware.Token(&User{
			Id:   1,
			Name: "user 1",
			Age:  18,
			Role: "ADMIN",
		})

		return c.JSON(200, echo.Map{"token": token})
	})

	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, echo.Map{"OK": 1})
	}, authMiddleware.BuildMiddleware("ADMIN"))

	e.Start(":8080")
}
