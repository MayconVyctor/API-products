package controler

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productControler struct {
	productUseCase usecase.ProductUsecase
}

func NewProductControler(usecase usecase.ProductUsecase) productControler {
	return productControler{
		productUseCase: usecase,
	}
}

func (p *productControler) GetProducts(ctx *gin.Context) {

	products, err := p.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, products)
}

func (p *productControler) CreateProduct(ctx *gin.Context) {

	var product model.Product
	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := p.productUseCase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (p *productControler) GetProductById(ctx *gin.Context) {

	id := ctx.Param("productId")
	if id == "" {
		response := model.Response{
			Message: "Id from product cannot null",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "The product id must be a number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.productUseCase.GetProductById(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		response := model.Response{
			Message: "The product not found in database",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (p *productControler) UpdateProduct(ctx *gin.Context) {
	id := ctx.Param("productId")
	if id == "" {
		response := model.Response{
			Message: "Id from product cannot null",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "The product id must be a number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var productRequest model.Product
	if err := ctx.ShouldBindJSON(&productRequest); err != nil {
		response := model.Response{
			Message: "Error reading product data",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	updatedProduct, err := p.productUseCase.UpdateProduct(productId, productRequest)
	if err != nil {
		response := model.Response{
			Message: "Error to updated product",
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	if updatedProduct.ID == 0 {
		response := model.Response{
			Message: "The product not found in database",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, updatedProduct)
}

func (p *productControler) UpdateStock(ctx *gin.Context) {

	id := ctx.Param("productId")
	if id == "" {
		response := model.Response{
			Message: "Id from product cannot null",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "The product id must be a number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var StockRequest model.Product
	if err := ctx.ShouldBindJSON(&StockRequest); err != nil {
		response := model.Response{
			Message: "Error reading product data",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	updatedStock, err := p.productUseCase.UpdateStock(productId, StockRequest)
	if err != nil {
		response := model.Response{
			Message: "Error to updated product",
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	if updatedStock == nil {
		response := model.Response{
			Message: "The product not found in database",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, updatedStock)
}

func (p *productControler) DeleteProduct(ctx *gin.Context) {

	id := ctx.Param("productId")
	if id == "" {
		response := model.Response{
			Message: "Id from product cannot null",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "The product id must be a number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.productUseCase.DeleteProduct(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		response := model.Response{
			Message: "The product not found in database",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	deleteStudent, err := p.productUseCase.DeleteProduct(productId)

	ctx.JSON(http.StatusOK, deleteStudent)
}
