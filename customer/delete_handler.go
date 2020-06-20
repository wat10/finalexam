package customer

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func deleteCustomerByIDHandler(c *gin.Context) {
	id := c.Param("id")
	if err := deleteCustomerByID(id); err != nil {
		fmt.Printf("delete customer by id error: % #v\n", err.OriginError)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Code, "Message": err.Message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "customer deleted"})
}
