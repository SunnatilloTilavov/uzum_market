package handler

import (
	us "backend_course/customer_api_gateway/genproto/user_service"
	"backend_course/customer_api_gateway/pkg/validator"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Router		/v1/seller/create [post]
// @Summary		Creates a seller
// @Description	This api creates a seller and returns its id
// @Tags		Seller
// @Accept		json
// @Produce		json
// @Param		seller body user_service.CreateSeller true "seller"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) CreateSeller(c *gin.Context) {
	req := &us.CreateSeller{}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	if !validator.ValidatePhone(req.Phone){
		handleGrpcErrWithDescription(c, h.log, errors.New("error while validating phone"), "wrong phone")
		return
	}

	if !validator.ValidateGmail(req.Gmail){
		handleGrpcErrWithDescription(c, h.log, errors.New("error while validating gmail"), "wrong gmail")
		return
	}

	resp, err := h.grpcClient.SellerService().Create(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while get create seller", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while create seller")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/seller/getbyid [post]
// @Summary		Get by id a seller
// @Description	This api get by id a seller
// @Tags		Seller
// Accept		json
// @Produce		json
// @Param		seller body user_service.SellerPrimaryKey true "seller"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetByIdSeller(c *gin.Context) {
	id := &us.SellerPrimaryKey{}

	if err := c.ShouldBindJSON(&id); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.SellerService().GetByID(c.Request.Context(), id)
	if err != nil {
		fmt.Errorf("error while get delete", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while getting by id")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/seller/getlist [post]
// @Summary		Get list a seller
// @Description	This api get list a seller
// @Tags		Seller
// Accept		json
// @Produce		json
// @Param		seller body user_service.GetListSellerRequest true "seller"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetListSeller(c *gin.Context) {
	req := &us.GetListSellerRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.SellerService().GetList(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while get list", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while get list")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/seller/update [PUT]
// @Summary		Update a seller
// @Description	This API updates a seller
// @Tags		Seller
// @Accept		json
// @Produce		json
// @Param		seller body user_service.UpdateSellerRequest true "Seller object to update"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) UpdateSeller(c *gin.Context) {
	req := &us.UpdateSellerRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while update seller")
		return
	}

	resp, err := h.grpcClient.SellerService().Update(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while update seller", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while ")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/seller/delete [delete]
// @Summary		delete a seller
// @Description	This api delete a seller
// @Tags		Seller
// Accept		json
// @Produce		json
// @Param		seller body user_service.SellerPrimaryKey true "seller"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) DeleteSeller(c *gin.Context) {
	id := &us.SellerPrimaryKey{}

	if err := c.ShouldBindJSON(&id); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.SellerService().Delete(c.Request.Context(), id)
	if err != nil {
		fmt.Errorf("error while delete", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while delete seller")
		return
	}

	c.JSON(http.StatusOK, resp)
}
