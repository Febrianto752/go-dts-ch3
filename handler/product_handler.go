package handler

import (
	"net/http"

	"github.com/Febrianto752/go-dts-ch3/entity"
	"github.com/Febrianto752/go-dts-ch3/helper"
	"github.com/Febrianto752/go-dts-ch3/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type ProductHandler interface {
	CreateProductHandler(ctx *gin.Context)
}

type productHandler struct {
	productService service.ProductService
}

func (h *productHandler) CreateProductHandler(ctx *gin.Context) {
	var payload entity.ProductRequest

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		helper.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	product, err := h.productService.Create(payload, userID)
	if err != nil {
		helper.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.SuccessResponse(ctx, http.StatusCreated, product)

}

func NewProductHandler(productService service.ProductService) ProductHandler {
	return &productHandler{productService: productService}
}
