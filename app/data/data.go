package data

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// favourite trees struct.
type tree struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Height float64 `json:"height"`
}

var trees = []tree{
	{ID: "1", Title: "avocado tree", Height: 5.67},
	{ID: "2", Title: "walnut", Height: 4.55},
	{ID: "3", Title: "almond tree", Height: 4.01},
}

// getTrees responds with the list of all trees as JSON.
func GetTrees(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, trees)
}

//health status
func GetHealth(c *gin.Context) {
	//Print Result in case of 200 status
	c.JSON(http.StatusOK, gin.H{"status": "UP"})
}
