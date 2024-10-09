package handlers

import (
	"github.com/emilpoppler134/api-template/internal/http"
	"gorm.io/gorm"
)

type HandlersInterface interface {
	RegisterRoutes(router *http.Router)
}

type Handlers struct {
	Product ProductInterface
	HandlersInterface
}

func Init(database *gorm.DB) HandlersInterface {
	return &Handlers{
		Product: &ProductHandler{db: database},
	}
}

func (handler Handlers) RegisterRoutes(router *http.Router) {
	router.GET("/products", handler.Product.List)
	router.GET("/products/:id", handler.Product.Find)
}
