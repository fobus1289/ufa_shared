package main

import (
	"errors"

	"github.com/fobus1289/ufa_shared/jwtService"
	"github.com/fobus1289/ufa_shared/middleware"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type User struct {
	Id   int64
	Name string
	Age  int
	Role string
}

func (u *User) ID() int64 {
	return u.Id
}

func (u *User) Pre(_ *gorm.DB, _ echo.Context) error {
	return errors.New("Forbidden")
}

func (u *User) PreWithPermission(_ *gorm.DB, _ echo.Context, _ ...string) error {
	return errors.New("Forbidden")
}

func main() {
	jwtService := jwtService.NewJwtService[*User](
		jwtService.JwtConfig{
			Secret:  "1234",
			Expired: 15,
		},
	)

	authMiddleware := middleware.NewAuthMiddleware(jwtService, nil, nil)

	e := echo.New()

	e.POST("/", func(c echo.Context) error {
		token, _ := jwtService.Token(&User{
			Id:   1,
			Name: "user 1",
			Age:  18,
			Role: "ADMIN",
		})
		return c.JSON(200, echo.Map{"token": token})
	})

	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, echo.Map{"OK": 1})
	}, authMiddleware.JwtPermissionMiddleware("ADMIN"))

	e.Start(":8080")
}
