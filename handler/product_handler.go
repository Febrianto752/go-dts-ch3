package handler

import (
	"net/http"
	"strconv"

	"github.com/Febrianto752/go-dts-ch3/entity"
	"github.com/Febrianto752/go-dts-ch3/helper"
	"github.com/Febrianto752/go-dts-ch3/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type ProductHandler interface {
	CreateProductHandler(ctx *gin.Context)
	GetProductsHandler(ctx *gin.Context)
	GetProductHandler(ctx *gin.Context)
	UpdateProductHandler(ctx *gin.Context)
	DeleteProductHandler(ctx *gin.Context)
}

type productHandler struct {
	productService service.ProductService
}

func (h *productHandler) CreateProductHandler(ctx *gin.Context) {
	var payload entity.ProductRequest

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		helper.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if userData["role"] != "admin" {
		if userId != payload.UserId {
			helper.FailResponse(ctx, http.StatusUnauthorized, "you cannot access this data")
			return
		}
	}

	product, err := h.productService.Create(payload)
	if err != nil {
		helper.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.SuccessResponse(ctx, http.StatusCreated, product)

}

func (h *productHandler) GetProductsHandler(ctx *gin.Context) {
	products, err := h.productService.GetAll()

	if err != nil {
		helper.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.SuccessResponse(ctx, http.StatusOK, products)
}

func (h *productHandler) GetProductHandler(ctx *gin.Context) {
	requestParam := ctx.Param("id")
	productId, _ := strconv.Atoi(requestParam)

	product, err := h.productService.GetById(uint(productId))
	if err != nil {
		helper.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.SuccessResponse(ctx, http.StatusOK, product)

}

func (h *productHandler) UpdateProductHandler(ctx *gin.Context) {
	requestParam := ctx.Param("id")
	productId, _ := strconv.Atoi(requestParam)

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	var payload entity.ProductRequest
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		helper.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	product, err := h.productService.Update(payload, uint(productId), userID)
	if err != nil {
		helper.FailResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helper.SuccessResponse(ctx, http.StatusOK, product)
}

func (h *productHandler) DeleteProductHandler(ctx *gin.Context) {
	// var product entity.Product
	requestParam := ctx.Param("id")
	productId, _ := strconv.Atoi(requestParam)
	// product.Id = uint(productId)

	h.productService.Delete(uint(productId))
	helper.SuccessResponse(ctx, http.StatusOK, nil)
}

func NewProductHandler(productService service.ProductService) ProductHandler {
	return &productHandler{productService: productService}
}
