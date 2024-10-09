package handlers

import (
	"github.com/emilpoppler134/api-template/internal/http"
	"github.com/emilpoppler134/api-template/internal/models"
	"github.com/emilpoppler134/api-template/internal/utils"
	"gorm.io/gorm"
)

type ProductInterface interface {
	List(params http.Params, body http.Body) http.Response
	Find(params http.Params, body http.Body) http.Response
}

type ProductHandler struct {
	db *gorm.DB
	ProductInterface
}

func (handler ProductHandler) List(params http.Params, body http.Body) http.Response {
	products := []models.Product{}
	result := handler.db.Model(&models.Product{}).Find(&products)

	if result.Error != nil {
		return http.InternalServerError()
	}

	return http.Ok(products)
}

func (handler ProductHandler) Find(params http.Params, body http.Body) http.Response {
	id, err := utils.ParseInt(params["id"])
	if err != nil {
		return http.NotFound("Invalid parameters")
	}

	product := models.Product{}
	result := handler.db.Model(&models.Product{}).Where("id = ?", id).Limit(1).Find(&product)

	if result.Error != nil {
		return http.InternalServerError()
	}
	if result.RowsAffected != 1 {
		return http.NotFound("There is no product with that id")
	}

	return http.Ok(product)
}
