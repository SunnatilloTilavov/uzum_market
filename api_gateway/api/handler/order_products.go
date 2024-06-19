package handler

import (
	"net/http"
	"backend_course/customer_api_gateway/genproto/order_product_service"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateOrderProduct godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/orderProduct [POST]
// @Summary 	Create an orderProduct
// @Description API for creating an order products
// @Tags 		order-products
// @Accept  	json
// @Produce  	json
// @Param		order body order_product_service.CreateOrderProduct true "order-products"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) CreateOrderProduct(c *gin.Context) {
	orderProduct := order_product_service.CreateOrderProduct{}

	if err := c.ShouldBindJSON(&orderProduct); err != nil {
		handleGrpcErrWithDescription(c, h.log,err, "error while reading request body")
		return
	}

	resp, err := h.grpcClient.OrderProductsService().Create(c.Request.Context(), &orderProduct)

	if err != nil {
		handleGrpcErrWithDescription(c, h.log,err, "error while creating order product")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetByIdOrderProduct godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/orderProduct/{id} [GET]
// @Summary 	Get an order-product
// @Description API for getting an order
// @Tags 		order-products
// @Accept  	json
// @Produce  	json
// @Param		id path string true "id"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) GetByIdOrderProduct(c *gin.Context) {
	id := c.Param("id")

	if err := uuid.Validate(id); err != nil {
		handleGrpcErrWithDescription(c, h.log,err, "error while validating order productId")
		return
	}

	orderProductId := order_product_service.OrderProductPrimaryKey{
		Id: id,
	}

	resp, err := h.grpcClient.OrderProductsService().GetById(c.Request.Context(), &orderProductId)

	if err != nil {
		handleGrpcErrWithDescription(c, h.log,err, "error while getting an order product")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateOrderProduct godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/orderProduct [PUT]
// @Summary 	Update an order product
// @Description API for update an order product
// @Tags 		order-products
// @Accept  	json
// @Produce  	json
// @Param		order body order_product_service.UpdateOrderProduct true "order-products"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) UpdateOrderProduct(c *gin.Context) {
	orderProduct := order_product_service.UpdateOrderProduct{}

	if err := c.ShouldBindJSON(&orderProduct); err != nil {
		handleGrpcErrWithDescription(c, h.log,err, "error while reading request body")
		return
	}

	resp, err := h.grpcClient.OrderProductsService().Update(c.Request.Context(), &orderProduct)

	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err,"error while updating an order product")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteOrderProduct godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/orderProduct/{id} [DELETE]
// @Summary 	Delete an order product
// @Description API for delete an order product
// @Tags 		order-products
// @Accept  	json
// @Produce  	json
// @Param		id path string true "id"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) DeleteOrderProduct(c *gin.Context) {
	id := c.Param("id")

	if err := uuid.Validate(id); err != nil {
		handleGrpcErrWithDescription(c, h.log,err, "error while validating order productId")
		return
	}

	order := order_product_service.OrderProductPrimaryKey{
		Id: id,
	}

	resp, err := h.grpcClient.OrderProductsService().Delete(c.Request.Context(), &order)

	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err,"error while deleting an order product")
		return
	}
	
	c.JSON(http.StatusOK, resp)
}

// GetAllOrderProducts godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/orderProducts [GET]
// @Summary 	Get all orders
// @Description API for Get all order products
// @Tags 		order-products
// @Accept  	json
// @Produce  	json
// @Param    	page query int false "page"
// @Param    	limit query int false "limit"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) GetAllOrderProducts(c *gin.Context) {
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

	orderProduct := order_product_service.GetListOrderProductRequest{
		Offset:int64((page - 1) * limit),
		Limit:  int64(limit),
	}

	resp, err := h.grpcClient.OrderProductsService().GetAll(c.Request.Context(), &orderProduct)

	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err,"error while getting all order products")
		return
	}

	
	c.JSON(http.StatusOK, resp)
}
