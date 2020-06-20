package customer

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func postCustomerHandler(c *gin.Context) {
	var json Customer
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("%#v", json)

	if err := insertCustomer(&json); err != nil {
		fmt.Printf("insert error: % #v\n", err.OriginError)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Code, "Message": err.Message})
		return
	}

	c.JSON(http.StatusCreated, json)
}
