package customer

import "github.com/gin-gonic/gin"

// SetupRouter is a SetupRouter
func SetupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/customers")
	api.Use(authMiddleware)
	api.POST("/", postCustomerHandler)
	api.GET("/:id", getCustomerByIDHandler)
	api.GET("/", getCustomersHandler)
	api.PUT("/:id", updateCustomerByIDHandler)
	api.DELETE("/:id", deleteCustomerByIDHandler)
	return r
}
