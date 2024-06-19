package handler

import (
	"backend_course/customer_api_gateway/genproto/catalog_service"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /create/category [post]
// @Summary Create category
// @Description API for creating a category
// @Tags category
// @Accept json
// @Produce json
// @Param category body catalog_service.CreateCategory true "Category"
// @Success 200 {object} catalog_service.Category
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) CreateCategory(c *gin.Context) {
	var req catalog_service.CreateCategory

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	resp, err := h.grpcClient.CategoryService().Create(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /category/{id} [get]
// @Summary Get a single category by ID
// @Description API for getting a single category by ID
// @Tags category
// @Accept json
// @Produce json
// @Param id path string true "category ID"
// @Success 200 {object} catalog_service.GetCategory
// @Failure 404 {object} models.ResponseError "Category not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) GetCategoryByID(c *gin.Context) {
	id := c.Param("id")
	req := &catalog_service.CategoryPrimaryKey{Id: id}

	resp, err := h.grpcClient.CategoryService().GetByID(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /category/{id} [put]
// @Summary Update a category by ID
// @Description API for updating a category by ID
// @Tags category
// @Accept json
// @Produce json
// @Param id path string true "category ID"
// @Param category body catalog_service.UpdateCategory true "Category"
// @Success 200 {object} catalog_service.GetCategory
// @Failure 400 {object} models.ResponseError "Invalid request body"
// @Failure 404 {object} models.ResponseError "Category not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var req catalog_service.UpdateCategory

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	req.Id = id
	resp, err := h.grpcClient.CategoryService().Update(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /getlistCategory [get]
// @Summary Get List of Categories
// @Description API for getting a list of categories
// @Tags category
// @Accept json
// @Produce json
// @Param limit query string true "Limit"
// @Param page query string true "Page"
// @Param search query string false "Search term"
// @Success 200 {object} catalog_service.GetListCategoryResponse
// @Failure 400 {object} models.ResponseError "Invalid query parameters"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) GetListCategory(c *gin.Context) {
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

	req := catalog_service.GetListCategoryRequest{
		Limit:  int64(limit),
		Page:   int64(page),
		Search: c.Query("search"),
	}

	resp, err := h.grpcClient.CategoryService().GetList(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /category/{id} [delete]
// @Summary Delete a category by ID
// @Description API for deleting a category by ID
// @Tags category
// @Accept json
// @Produce json
// @Param id path string true "category ID"
// @Success 200 {object} catalog_service.Empty
// @Failure 404 {object} models.ResponseError "Category not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	req := &catalog_service.CategoryPrimaryKey{Id: id}

	resp, err := h.grpcClient.CategoryService().Delete(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}


// @Security ApiKeyAuth
// @Router /getcategorywith/{id} [get]
// @Summary Get a single category by ID
// @Description API for getting a single category by ID
// @Tags category
// @Accept json
// @Produce json
// @Param id path string true "category ID"
// @Success 200 {object} catalog_service.GetCategory
// @Failure 404 {object} models.ResponseError "Category not found"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *handler) GetCategoryWithProductId(c *gin.Context) {
	id := c.Param("id")
	resp:=&catalog_service.GetCategory{}
	req := &catalog_service.CategoryPrimaryKey{Id: id}

	resp, err := h.grpcClient.CategoryService().GetByID(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}