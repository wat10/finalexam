package customer

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getCustomerByIDHandler(c *gin.Context) {
	id := c.Param("id")
	fmt.Printf("get customer by id : % #v\n", id)
	customer, err := getCustomerByID(id)
	if err != nil {
		fmt.Printf("get customer by id error: % #v\n", err.OriginError)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Code, "Message": err.Message})
		return
	}
	c.JSON(http.StatusOK, customer)
}

func getCustomersHandler(c *gin.Context) {
	customers, err := getCustomers()
	if err != nil {
		fmt.Printf("get customers error: % #v\n", err.OriginError)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Code, "Message": err.Message})
		return
	}
	c.JSON(http.StatusOK, customers)
}
