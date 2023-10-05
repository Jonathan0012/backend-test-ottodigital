package main

import (
	middleware "github.com/Jonathan0012/backend-test-ottodigital/Middleware"
	"github.com/Jonathan0012/backend-test-ottodigital/app"
	"github.com/Jonathan0012/backend-test-ottodigital/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	app.NewDB()

	r.GET("/users", controllers.IndexUsers)
	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:id", controllers.ShowUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	r.Use(middleware.Authenticate)

	//Authentication code is abc as it's written in authentication.go middleware
	r.GET("/tasks", controllers.IndexTasks)
	r.POST("/tasks", controllers.CreateTask)
	r.GET("/tasks/:id", controllers.ShowTask)
	r.PUT("/tasks/:id", controllers.UpdateTask)
	r.DELETE("/tasks/:id", controllers.DeleteTask)

	r.Run()
}
