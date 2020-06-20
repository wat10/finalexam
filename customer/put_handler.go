package customer

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func updateCustomerByIDHandler(c *gin.Context) {
	id := c.Param("id")
	var json Customer
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	json.ID = cid
	if err := updateCustomerByID(&json); err != nil {
		fmt.Printf("update error: % #v\n", err.OriginError)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Code, "Message": err.Message})
		return
	}
	c.JSON(http.StatusOK, json)
}
