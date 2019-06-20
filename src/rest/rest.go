package rest

import (
	"github.com/gin-gonic/gin"
)

//RunAPI Get request
func RunAPI(address string) error {
	r := gin.Default()

	//get products

	r.GET("/products", func(c *gin.Context) {
		// TODO: return a list of all products
	})

	//get promos

	r.GET("/promos", func(c *gin.Context) {
		// TODO: return a list of all promotions to the client
	})

	//post user sign in

	r.POST("/users/signin", func(c *gin.Context) {
		//TODO: sign in a user
	})

	//add user

	r.POST("/users", func(c *gin.Context) {
		//TODO: add a user
	})

	r.POST("/users/:id/signout", func(c *gin.Context) {
		//TODO: signout
	})

	r.GET("/users/:id/orders", func(c *gin.Context) {
		// TODO: get user id orders
	})

	r.POST("/users/charge", func(c *gin.Context) {
		//TODO: charge credit card
	})

}
