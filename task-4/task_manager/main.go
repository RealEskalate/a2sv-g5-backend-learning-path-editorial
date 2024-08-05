package main

import (
	router "task-management-api/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.SetupRouter(r)
}
