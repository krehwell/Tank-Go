package users

import (
	"final-project/database"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRepository struct {
	database database.Database
}

func (u *UserRepository) registerUser(c *gin.Context) {
	fmt.Println("someone hit me")
	c.JSON(http.StatusOK, gin.H{
		"message": "you are good",
	})
}

func (u *UserRepository) loginUser(c *gin.Context) {

}
