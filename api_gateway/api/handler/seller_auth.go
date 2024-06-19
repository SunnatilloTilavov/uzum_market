package handler

import (
	as "backend_course/customer_api_gateway/genproto/auth_service"
	"backend_course/customer_api_gateway/pkg/validator"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Router		/v1/seller/register [post]
// @Summary		Register a seller
// @Description	This api register a seller
// @Tags		Seller_Auth
// @Accept		json
// @Produce		json
// @Param		seller body auth_service.SellerGmailCheckRequest true "seller"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) RegisterSeller(c *gin.Context) {
	req := &as.SellerGmailCheckRequest{}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	if !validator.ValidateGmail(req.Gmail) {
		handleGrpcErrWithDescription(c, h.log,errors.New("wrong gmail"), "error while validating gmail")
		return
	}

	resp, err := h.grpcClient.SellerAuthService().SellerRegisterByMail(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while register seller", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while register seller")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/seller/registercomfirm [post]
// @Summary		Register Confirm a seller
// @Description	This api register confirm a seller
// @Tags		Seller_Auth
// @Accept		json
// @Produce		json
// @Param		seller body auth_service.SellerRConfirm true "seller"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) RegisterConfirmSeller(c *gin.Context) {
	req := &as.SellerRConfirm{}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.SellerAuthService().SellerRegisterByMailConfirm(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while register confirm seller", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while register confirm seller")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/seller/login [post]
// @Summary		Login  a seller
// @Description	This api login a seller
// @Tags		Seller_Auth
// @Accept		json
// @Produce		json
// @Param		seller body auth_service.SellerGmailCheckRequest true "seller"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) LoginSeller(c *gin.Context) {
	req := &as.SellerGmailCheckRequest{}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.SellerAuthService().SellerLoginByGmail(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while login seller", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while login seller")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/seller/loginconfirm [post]
// @Summary		Login confirm a seller
// @Description	This api login confirm a seller
// @Tags		Seller_Auth
// @Accept		json
// @Produce		json
// @Param		seller body auth_service.SellerRConfirm true "seller"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) SellerLoginConfirm(c *gin.Context) {
	req := &as.SellerRConfirm{}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.SellerAuthService().SellerRegisterByMailConfirm(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while login confirm seller", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while login confirm seller")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/seller/resetpassword [put]
// @Summary		Reset a password a seller
// @Description	This api reset a password a seller
// @Tags		Seller_Auth
// @Accept		json
// @Produce		json
// @Param		seller body auth_service.SellerGmailCheckRequest true "seller"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) ResetPasswordSeller(c *gin.Context) {
	req := &as.SellerGmailCheckRequest{}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.SellerAuthService().SellerResetPassword(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while reset password seller", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while login seller")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/seller/resetconfirm [put]
// @Summary		Confirm reset password a seller
// @Description	This api reset password confirm a seller
// @Tags		Seller_Auth
// @Accept		json
// @Produce		json
// @Param		seller body auth_service.SellerPasswordConfirm true "seller"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) SellerResetPasswordConfirm(c *gin.Context) {
	req := &as.SellerPasswordConfirm{}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.SellerAuthService().SellerResetPasswordConfirm(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while reset password confirm seller", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while reset password confirm seller")
		return
	}

	c.JSON(http.StatusOK, resp)
}