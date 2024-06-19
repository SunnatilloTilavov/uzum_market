package handler

import (
	as "backend_course/customer_api_gateway/genproto/auth_service"
	"backend_course/customer_api_gateway/pkg/validator"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Router		/v1/customer/register [post]
// @Summary		Register a customer
// @Description	This api register a customer
// @Tags		Customer_Auth
// @Accept		json
// @Produce		json
// @Param		customer body auth_service.GmailCheckRequest true "customer"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) RegisterCustomer(c *gin.Context) {
	req := &as.GmailCheckRequest{}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	if !validator.ValidateGmail(req.Gmail) {
		handleGrpcErrWithDescription(c, h.log,errors.New("wrong gmail"), "error while validating gmail")
		return
	}

	resp, err := h.grpcClient.CustomerAuthService().RegisterByMail(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while register customer", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while register customer")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/customer/registercomfirm [post]
// @Summary		Register Confirm a customer
// @Description	This api register confirm a customer
// @Tags		Customer_Auth
// @Accept		json
// @Produce		json
// @Param		customer body auth_service.RConfirm true "customer"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) RegisterConfirmCustomer(c *gin.Context) {
	req := &as.RConfirm{}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.CustomerAuthService().RegisterByMailConfirm(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while register confirm customer", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while register confirm customer")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/customer/login [post]
// @Summary		Login  a customer
// @Description	This api login a customer
// @Tags		Customer_Auth
// @Accept		json
// @Produce		json
// @Param		customer body auth_service.GmailCheckRequest true "customer"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) LoginCustomer(c *gin.Context) {
	req := &as.GmailCheckRequest{}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.CustomerAuthService().LoginByGmail(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while login customer", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while login customer")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/customer/loginconfirm [post]
// @Summary		Login confirm a customer
// @Description	This api login confirm a customer
// @Tags		Customer_Auth
// @Accept		json
// @Produce		json
// @Param		customer body auth_service.LoginByGmailRequest true "customer"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) LoginConfirm(c *gin.Context) {
	req := &as.LoginByGmailRequest{}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.CustomerAuthService().LoginByGmailComfirm(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while login confirm customer", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while login confirm customer")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/customer/resetpassword [put]
// @Summary		Reset a password a customer
// @Description	This api reset a password a customer
// @Tags		Customer_Auth
// @Accept		json
// @Produce		json
// @Param		customer body auth_service.GmailCheckRequest true "customer"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) ResetPasswordCustomer(c *gin.Context) {
	req := &as.GmailCheckRequest{}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.CustomerAuthService().ResetPassword(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while reset password customer", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while login customer")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/customer/resetconfirm [put]
// @Summary		Confirm reset password a customer
// @Description	This api reset password confirm a customer
// @Tags		Customer_Auth
// @Accept		json
// @Produce		json
// @Param		customer body auth_service.CustomerPasswordConfirm true "customer"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) ResetPasswordConfirm(c *gin.Context) {
	req := &as.CustomerPasswordConfirm{}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.CustomerAuthService().ResetPasswordConfirm(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while reset password confirm customer", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while reset password confirm customer")
		return
	}

	c.JSON(http.StatusOK, resp)
}