package handler

import (
	"backend_course/customer_api_gateway/genproto/catalog_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /create/product_review [post]
// @Summary Create product review
// @Description API for creating a product review
// @Tags product_review
// @Accept json
// @Produce json
// @Param review body catalog_service.CreateProductReviewRequest true "ProductReview"
// @Success 200 {object} catalog_service.ProductReview
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) CreateProductReview(c *gin.Context) {
	var req catalog_service.CreateProductReviewRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	resp, err := h.grpcClient.ProductReviewService().CreateProductReview(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /product_review/{id} [get]
// @Summary Get a single product review by ID
// @Description API for getting a single product review by ID
// @Tags product_review
// @Accept json
// @Produce json
// @Param id path string true "ProductReview ID"
// @Success 200 {object} catalog_service.ProductReview
// @Failure 404 {object} models.ResponseError "Product review not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) GetProductReviewByID(c *gin.Context) {
	id := c.Param("id")
	req := &catalog_service.ProductReviewPrimaryKey{Id: id}

	resp, err := h.grpcClient.ProductReviewService().GetProductReviewByID(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /product_reviews/product/{product_id} [get]
// @Summary Get product reviews by product ID
// @Description API for getting product reviews by product ID
// @Tags product_review
// @Accept json
// @Produce json
// @Param product_id path string true "Product ID"
// @Success 200 {object} catalog_service.GetProductReviewsByProductIDResponse
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) GetProductReviewsByProductID(c *gin.Context) {
	productID := c.Param("product_id")
	req := &catalog_service.GetProductReviewsByProductIDRequest{ProductId: productID}

	resp, err := h.grpcClient.ProductReviewService().GetProductReviewsByProductID(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /product_reviews/customer/{customer_id} [get]
// @Summary Get product reviews by customer ID
// @Description API for getting product reviews by customer ID
// @Tags product_review
// @Accept json
// @Produce json
// @Param customer_id path string true "Customer ID"
// @Success 200 {object} catalog_service.GetProductReviewsByCustomerIDResponse
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) GetProductReviewsByCustomerID(c *gin.Context) {
	customerID := c.Param("customer_id")
	req := &catalog_service.GetProductReviewsByCustomerIDRequest{CustomerId: customerID}

	resp, err := h.grpcClient.ProductReviewService().GetProductReviewsByCustomerID(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /product_review/{id} [put]
// @Summary Update a product review by ID
// @Description API for updating a product review by ID
// @Tags product_review
// @Accept json
// @Produce json
// @Param id path string true "ProductReview ID"
// @Param review body catalog_service.UpdateProductReviewRequest true "ProductReview"
// @Success 200 {object} catalog_service.ProductReview
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 404 {object} models.ResponseError "Product review not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) UpdateProductReview(c *gin.Context) {
	id := c.Param("id")
	var req catalog_service.UpdateProductReviewRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	req.Id = id
	resp, err := h.grpcClient.ProductReviewService().UpdateProductReview(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /product_review/{id} [delete]
// @Summary Delete a product review by ID
// @Description API for deleting a product review by ID
// @Tags product_review
// @Accept json
// @Produce json
// @Param id path string true "ProductReview ID"
// @Success 200 {object} catalog_service.Empty4
// @Failure 404 {object} models.ResponseError "Product review not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) DeleteProductReview(c *gin.Context) {
	id := c.Param("id")
	req := &catalog_service.ProductReviewPrimaryKey{Id: id}

	resp, err := h.grpcClient.ProductReviewService().DeleteProductReview(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}
