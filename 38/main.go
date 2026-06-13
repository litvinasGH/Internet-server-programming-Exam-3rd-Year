package main

import (
	"github.com/kataras/iris/v12"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "38/docs"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{1, "Ivan"},
	{2, "Alex"},
	{3, "Maria"},
}

// @title Iris Swagger API
// @version 1.0
// @description REST API on Iris
// @host localhost:8080
// @BasePath /
func main() {

	app := iris.New()

	app.Get("/users", getUsers)
	app.Get("/users/{id:int}", getUser)
	app.Post("/users", createUser)
	app.Put("/users/{id:int}", updateUser)
	app.Delete("/users/{id:int}", deleteUser)

	app.Get("/swagger/{any:path}",
		iris.FromStd(httpSwagger.WrapHandler),
	)

	app.Listen(":8080")
}

// @Summary Get all users
// @Tags Users
// @Produce json
// @Success 200 {array} User
// @Router /users [get]
func getUsers(ctx iris.Context) {
	ctx.JSON(users)
}

// @Summary Get user by id
// @Tags Users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Router /users/{id} [get]
func getUser(ctx iris.Context) {

	id, _ := ctx.Params().GetInt("id")

	for _, user := range users {
		if user.ID == id {
			ctx.JSON(user)
			return
		}
	}

	ctx.StatusCode(404)
}

// @Summary Create user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body User true "User"
// @Success 201 {object} User
// @Router /users [post]
func createUser(ctx iris.Context) {

	var user User

	if err := ctx.ReadJSON(&user); err != nil {
		ctx.StatusCode(400)
		return
	}

	users = append(users, user)

	ctx.StatusCode(201)
	ctx.JSON(user)
}

// @Summary Update user
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body User true "User"
// @Success 200 {object} User
// @Router /users/{id} [put]
func updateUser(ctx iris.Context) {

	id, _ := ctx.Params().GetInt("id")

	var updated User

	if err := ctx.ReadJSON(&updated); err != nil {
		ctx.StatusCode(400)
		return
	}

	for i := range users {

		if users[i].ID == id {

			users[i].Name = updated.Name

			ctx.JSON(users[i])
			return
		}
	}

	ctx.StatusCode(404)
}

// @Summary Delete user
// @Tags Users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]string
// @Router /users/{id} [delete]
func deleteUser(ctx iris.Context) {

	id, _ := ctx.Params().GetInt("id")

	for i, user := range users {

		if user.ID == id {

			users = append(
				users[:i],
				users[i+1:]...,
			)

			ctx.JSON(iris.Map{
				"message": "deleted",
			})

			return
		}
	}

	ctx.StatusCode(404)
}
