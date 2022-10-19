package main

import (
	"FP1hacktiv8/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	routes.TodoRouters(router)
	router.Run(":8080")

}
