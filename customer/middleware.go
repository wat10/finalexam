package customer

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func authMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")
	isValidToken := token == "token2019wrong_token"
	if !isValidToken {
		//c.JSON(http.StatusUnauthorized, "Unauthorized")
		c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("%s", "Unauthorized"))
	} else {
		c.Next()
	}

}
