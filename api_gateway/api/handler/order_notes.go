package handler

import (
    "net/http"
    "backend_course/customer_api_gateway/genproto/order_notes"
	"strconv"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
)

// CreateOrderNotes godoc
// @Security ApiKeyAuth
// @Router      /api/v1/orderNote [POST]
// @Summary     Create an order notes
// @Description API for creating an order notes
// @Tags        order-notes
// @Accept      json
// @Produce     json
// @Param       order body order_notes.CreateOrderNotes true "order-notes"
// @Success     200 {object} models.Response
// @Failure     400 {object} models.Response
// @Failure     404 {object} models.Response
// @Failure     500 {object} models.Response
func (h *handler) CreateOrderNotes(c *gin.Context) {
    orderNotes := order_notes.CreateOrderNotes{}

    if err := c.ShouldBindJSON(&orderNotes); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid request body")
        return
    }

    resp, err := h.grpcClient.OrderProductNotesService().Create(c.Request.Context(), &orderNotes)
    if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid request body")
        return
    }
	c.JSON(http.StatusOK, resp)
   
}

// GetByIdOrderNotes godoc
// @Security ApiKeyAuth
// @Router      /api/v1/orderNote/{id} [GET]
// @Summary     Get an order notes
// @Description API for getting an order notes
// @Tags        order-notes
// @Accept      json
// @Produce     json
// @Param       id path string true "id"
// @Success     200 {object} models.Response
// @Failure     400 {object} models.Response
// @Failure     404 {object} models.Response
// @Failure     500 {object} models.Response
func (h *handler) GetByIdOrderNotes(c *gin.Context) {
    id := c.Param("id")

    if _, err := uuid.Parse(id); err != nil {
        handleGrpcErrWithDescription(c, h.log,err, "error while validating order notes id")
        return
    }

    orderNotes := order_notes.OrderNotesPrimaryKey{
        Id: id,
    }

    resp, err := h.grpcClient.OrderProductNotesService().GetById(c.Request.Context(), &orderNotes)
    if err != nil {
        handleGrpcErrWithDescription(c, h.log,err, "error while getting an order notes")
        return
    }

	c.JSON(http.StatusOK, resp)
}

// UpdateOrderNotes godoc
// @Security ApiKeyAuth
// @Router      /api/v1/orderNote [PUT]
// @Summary     Update an order notes
// @Description API for update an order notes
// @Tags        order-notes
// @Accept      json
// @Produce     json
// @Param       order body order_notes.UpdateOrderNotes true "order-notes"
// @Success     200 {object} models.Response
// @Failure     400 {object} models.Response
// @Failure     404 {object} models.Response
// @Failure     500 {object} models.Response
func (h *handler) UpdateOrderNotes(c *gin.Context) {
    orderNotes := order_notes.UpdateOrderNotes{}

    if err := c.ShouldBindJSON(&orderNotes); err != nil {
        handleGrpcErrWithDescription(c, h.log,err, "error while reading request body")
        return
    }

    resp, err := h.grpcClient.OrderProductNotesService().Update(c.Request.Context(), &orderNotes)
    if err != nil {
        handleGrpcErrWithDescription(c, h.log,err, "error while updating an order notes")
        return
    }

	c.JSON(http.StatusOK, resp)
}

// DeleteOrderNotes godoc
// @Security ApiKeyAuth
// @Router      /api/v1/orderNote/{id} [DELETE]
// @Summary     Delete an order notes
// @Description API for delete an order notes
// @Tags        order-notes
// @Accept      json
// @Produce     json
// @Param       id path string true "id"
// @Success     200 {object} models.Response
// @Failure     400 {object} models.Response
// @Failure     404 {object} models.Response
// @Failure     500 {object} models.Response
func (h *handler) DeleteOrderNotes(c *gin.Context) {
    id := c.Param("id")

    if _, err := uuid.Parse(id); err != nil {
        handleGrpcErrWithDescription(c, h.log,err, "error while validating order notes id")
        return
    }

    orderNotesPk := order_notes.OrderNotesPrimaryKey{
        Id: id,
    }

    resp, err := h.grpcClient.OrderProductNotesService().Delete(c.Request.Context(), &orderNotesPk)
    if err != nil {
        handleGrpcErrWithDescription(c, h.log,err, "error while deleting an order notes")
        return
    }

	c.JSON(http.StatusOK, resp)
}

// GetAllOrderNotes godoc
// @Security ApiKeyAuth
// @Router      /api/v1/orderNotes [GET]
// @Summary     Get all order notes
// @Description API for Get all order notes
// @Tags        order-notes
// @Accept      json
// @Produce     json
// @Param       page query int false "page"
// @Param       limit query int false "limit"
// @Success     200 {object} models.Response
// @Failure     400 {object} models.Response
// @Failure     404 {object} models.Response
// @Failure     500 {object} models.Response
func (h *handler) GetAllOrderNotes(c *gin.Context) {
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
    orderNotes := order_notes.GetListOrderNotesRequest{
        Offset: int64((page - 1) * limit),
        Limit: int64(limit),

    }

    resp, err := h.grpcClient.OrderProductNotesService().GetAll(c.Request.Context(), &orderNotes)
    if err != nil {
        handleGrpcErrWithDescription(c, h.log, err,"error while getting all order notes")
        return
    }

	c.JSON(http.StatusOK, resp)
}
