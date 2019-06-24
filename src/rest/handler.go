package rest

import (
	"fmt"
	"github.com/corganfuzz/ginapi/src/models/dblayer"
	"github.com/corganfuzz/ginapi/src/models/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// Allowing extensibility
type HandlerInterface interface {
	GetProducts(c *.gin.Context)
	GetPromos(c *.gin.Context)
	AddUser(c *.gin.Context)
	SignIn(c *.gin.Context)
	SignOut(c *.gin.Context)
	GetOrders(c *.gin.Context)
	Charge(c *.gin.Context)
}

type Handler struct {
	db dblayer.DBLayer
}

// Constructor

func NewHandler() (*Handler, error) {

	// this creates a new pointer to the handler object

	return new(Handler), nil
}

func (h *Handler) GetProducts(c *gin.Context) {
	if h.db == nil {
		return
	}
	products, err := h.db.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (h *Handler) GetPromos(c *.gin.Context) {
	if h.db == nil {
		return
	}
	promos, err := h.db.GetPromos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, 
		gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, promos)
}

