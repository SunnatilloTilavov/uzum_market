package handler

import (
	"backend_course/customer_api_gateway/genproto/catalog_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /create/product [post]
// @Summary Create product
// @Description API for creating a product
// @Tags product
// @Accept json
// @Produce json
// @Param product body catalog_service.CreateProduct true "Product"
// @Success 200 {object} catalog_service.Product
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) CreateProduct(c *gin.Context) {
	var req catalog_service.CreateProduct

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	resp, err := h.grpcClient.ProductService().Create(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /product/{id} [get]
// @Summary Get a single product by ID
// @Description API for getting a single product by ID
// @Tags product
// @Accept json
// @Produce json
// @Param id path string true "product ID"
// @Success 200 {object} catalog_service.Product
// @Failure 404 {object} models.ResponseError "Product not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) GetProductByID(c *gin.Context) {
	id := c.Param("id")
	req := &catalog_service.ProductPrimaryKey{Id: id}

	resp, err := h.grpcClient.ProductService().GetByID(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /product/{id} [put]
// @Summary Update a product by ID
// @Description API for updating a product by ID
// @Tags product
// @Accept json
// @Produce json
// @Param id path string true "product ID"
// @Param product body catalog_service.UpdateProduct true "Product"
// @Success 200 {object} catalog_service.Product
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 404 {object} models.ResponseError "Product not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var req catalog_service.UpdateProduct

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	req.Id = id
	resp, err := h.grpcClient.ProductService().Update(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /getlistProduct [get]
// @Summary Get List of Products
// @Description API for getting a list of products
// @Tags product
// @Accept json
// @Produce json
// @Param limit query string true "Limit"
// @Param page query string true "Page"
// @Param search query string false "Search term"
// @Success 200 {object} catalog_service.GetListProductResponse
// @Failure 400 {object} models.ResponseError "Invalid query parameters"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) GetListProduct(c *gin.Context) {
	limitStr := c.Query("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid limit parameter")
		return
	}

	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid page parameter")
		return
	}

	req := catalog_service.GetListProductRequest{
		Limit:  int64(limit),
		Page:   int64(page),
		Search: c.Query("search"),
	}

	resp, err := h.grpcClient.ProductService().GetList(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /product/{id} [delete]
// @Summary Delete a product by ID
// @Description API for deleting a product by ID
// @Tags product
// @Accept json
// @Produce json
// @Param id path string true "product ID"
// @Success 200 {object} catalog_service.Empty1
// @Failure 404 {object} models.ResponseError "Product not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	req := &catalog_service.ProductPrimaryKey{Id: id}

	resp, err := h.grpcClient.ProductService().Delete(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}
