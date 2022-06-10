package handler

import (
	"klikdokter/authorization"
	"klikdokter/helper"
	"klikdokter/products"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productServices products.Services
	authServices    authorization.Services
}

func NewProductHandler(productServices products.Services, authServices authorization.Services) *productHandler {
	return &productHandler{productServices, authServices}
}

func (h *productHandler) SaveProduct(c *gin.Context) {
	var input products.ProductInput
	err := c.ShouldBind(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Create Product Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	checkSkuExisted, err := h.productServices.FindBySku(input.Sku)
	if err != nil {
		response := helper.APIResponse("Create Product Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if checkSkuExisted.Sku == input.Sku {
		response := helper.APIResponse("Create Product Failed! SKU sudah ada", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newProduct, err := h.productServices.SaveProduct(input)
	if err != nil {
		response := helper.APIResponse("Create Product Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := products.FormatProduct(newProduct)
	response := helper.APIResponse("Create Product Success", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}

func (h *productHandler) GetAllProduct(c *gin.Context) {
	product, err := h.productServices.FindAll()
	if err != nil {
		response := helper.APIResponse("Get All Product Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := products.FormatProducts(product)
	response := helper.APIResponse("Get All Product Success", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) GetBySku(c *gin.Context) {
	var input products.SkuInput
	err := c.ShouldBind(&input)
	product, err := h.productServices.FindBySku(input.Sku)
	if err != nil {
		response := helper.APIResponse("Get Product By SKU Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if product.Sku == "" {
		response := helper.APIResponse("Get Product By SKU Failed! SKU Tidak di temukan", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := products.FormatProduct(product)
	response := helper.APIResponse("Get Product By SKU Success", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) UpdateDataBySKU(c *gin.Context) {

	var input products.ProductInput
	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Update Product Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	product, err := h.productServices.FindAndUpdateDataBySku(input)
	if err != nil {
		response := helper.APIResponse("Update Product Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	result, _ := h.productServices.FindBySku(product.Sku)
	if result.Sku == "" {
		response := helper.APIResponse("Update Product Failed! SKU Tidak di temukan", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Update Product Success", http.StatusOK, "success", result)
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) RemoveDataBySku(c *gin.Context) {
	var input products.SkuInput
	err := c.ShouldBind(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Remove Product Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	remove, err := h.productServices.FindBySku(input.Sku)
	if err != nil {
		response := helper.APIResponse("Remove Product Failed!", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if remove.Sku == "" {
		response := helper.APIResponse("Remove Product Failed! SKU Tidak di temukan", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	h.productServices.RemoveDataBySku(input.Sku)
	if err != nil {
		response := helper.APIResponse("Remove Product Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Remove Product Success", http.StatusOK, "success", remove)
	c.JSON(http.StatusOK, response)
}
