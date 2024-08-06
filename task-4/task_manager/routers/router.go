package router

import (
	controller "task-management-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/tasks", controller.GetTasks)
	r.GET("/tasks/:id", controller.GetTaskById)
	r.POST("/tasks", controller.CreateTask)
	r.PUT("/tasks", controller.UpdateTask)
	r.DELETE("/tasks/:id", controller.DeleteTask)
	r.Run(":8080")
}
