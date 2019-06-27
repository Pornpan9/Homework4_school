package todo

import(
	"github.com/gin-gonic/gin"
	"net/http"
)

type Todo struct {
	ID int64 `json:"id"`
	Title string `json:"title"`
	Status string `json:"status"`
}

func GetHandler( c *gin.Context)  {

	c.JSON(http.StatusOK, gin.H{"message":"Hello GetHandler"})
}