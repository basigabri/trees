package routes

import (
	d "tree-webservice/data"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()

	router.GET("/trees", d.GetTrees)
	router.GET("/healthz", d.GetHealth)
	router.Run(":8080")
}
