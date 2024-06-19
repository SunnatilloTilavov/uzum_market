package handler

import (
    "net/http"
    "strconv"
    "backend_course/customer_api_gateway/genproto/order_service"

    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
)

// CreateOrder godoc
// @Security ApiKeyAuth
// @Router      /api/v1/order [POST]
// @Summary     Create an order
// @Description API for creating an order
// @Tags        orders
// @Accept      json
// @Produce     json
// @Param       order body order_service.CreateOrder true "order"
// @Success     200 {object} models.Response
// @Failure     400 {object} models.Response
// @Failure     404 {object} models.Response
// @Failure     500 {object} models.Response
func (h *handler) CreateOrder(c *gin.Context) {
    order := order_service.CreateOrder{}

    if err := c.ShouldBindJSON(&order); err != nil {
        handleGrpcErrWithDescription(c, h.log, err, "error while creating an order")
        return
    }

    resp, err := h.grpcClient.OrderService().Create(c.Request.Context(), &order)
    if err != nil {
        handleGrpcErrWithDescription(c, h.log, err, "error while creating order")
        return
    }
    c.JSON(http.StatusOK, resp)
}

// GetByIdOrder godoc
// @Security ApiKeyAuth
// @Router      /api/v1/order/{id} [GET]
// @Summary     Get an order
// @Description API for getting an order
// @Tags        orders
// @Accept      json
// @Produce     json
// @Param       id path string true "id"
// @Success     200 {object} models.Response
// @Failure     400 {object} models.Response
// @Failure     404 {object} models.Response
// @Failure     500 {object} models.Response
func (h *handler) GetByIdOrder(c *gin.Context) {
    id := c.Param("id")

    if err := uuid.Validate(id); err != nil {
        handleGrpcErrWithDescription(c, h.log, err, "error while validating orderId")
        return
    }

    order := order_service.OrderPrimaryKey{
        Id: id,
    }

    resp, err := h.grpcClient.OrderService().GetById(c.Request.Context(), &order)
    if err != nil {
        handleGrpcErrWithDescription(c, h.log, err, "error while getting an order")
        return
    }

    c.JSON(http.StatusOK, resp)
}

// UpdateOrder godoc
// @Security ApiKeyAuth
// @Router      /api/v1/order [PUT]
// @Summary     Update an order
// @Description API for update an order
// @Tags        orders
// @Accept      json
// @Produce     json
// @Param       order body order_service.UpdateOrder true "order"
// @Success     200 {object} models.Response
// @Failure     400 {object} models.Response
// @Failure     404 {object} models.Response
// @Failure     500 {object} models.Response
func (h *handler) UpdateOrder(c *gin.Context) {
    order := order_service.UpdateOrder{}

    if err := c.ShouldBindJSON(&order); err != nil {
        handleGrpcErrWithDescription(c, h.log, err, "error while reading request body")
        return
    }

    resp, err := h.grpcClient.OrderService().Update(c.Request.Context(), &order)
    if err != nil {
        handleGrpcErrWithDescription(c, h.log, err, "error while updating an order")
        return
    }
    c.JSON(http.StatusOK, resp)
}

// DeleteOrder godoc
// @Security ApiKeyAuth
// @Router      /api/v1/order/{id} [DELETE]
// @Summary     Delete an order
// @Description API for delete an order
// @Tags        orders
// @Accept      json
// @Produce     json
// @Param       id path string true "id"
// @Success     200 {object} models.Response
// @Failure     400 {object} models.Response
// @Failure     404 {object} models.Response
// @Failure     500 {object} models.Response
func (h *handler) DeleteOrder(c *gin.Context) {
    id := c.Param("id")

    if err := uuid.Validate(id); err != nil {
        handleGrpcErrWithDescription(c, h.log, err, "error while validating orderId")
        return
    }

    order := order_service.OrderPrimaryKey{
        Id: id,
    }

    resp, err := h.grpcClient.OrderService().Delete(c.Request.Context(), &order)
    if err != nil {
        handleGrpcErrWithDescription(c, h.log, err, "error while deleting an order")
        return
    }
    c.JSON(http.StatusOK, resp)
}

// GetAllOrder godoc
// @Security ApiKeyAuth
// @Router      /api/v1/orders [GET]
// @Summary     Get all orders
// @Description API for Get all orders
// @Tags        orders
// @Accept      json
// @Produce     json
// @Param       search query string false "search"
// @Param       page query int false "page"
// @Param       limit query int false "limit"
// @Success     200 {object} models.Response
// @Failure     400 {object} models.Response
// @Failure     404 {object} models.Response
// @Failure     500 {object} models.Response
func (h *handler) GetAllOrder(c *gin.Context) {
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

    search := c.Query("search")

    order := order_service.GetListOrderRequest{
        Offset: int64((page - 1) * limit),
        Limit:  int64(limit),
        Search: search,
    }

    resp, err := h.grpcClient.OrderService().GetAll(c.Request.Context(), &order)
    if err != nil {
        handleGrpcErrWithDescription(c, h.log, err, "error while getting all orders")
        return
    }

    c.JSON(http.StatusOK, resp)
}
