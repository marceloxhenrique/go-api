package controller

import (
	"go-products/model"
	"go-products/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUsecase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		productUsecase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {

	products, err := p.productUsecase.GetAllProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *productController) GetProductById(ctx *gin.Context) {
	productId := ctx.Param("id")
	id, err := strconv.Atoi(productId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	product, err := p.productUsecase.GetProductById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func (p *productController) AddProduct(ctx *gin.Context) {

	var product model.Product
	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := p.productUsecase.AddProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}
