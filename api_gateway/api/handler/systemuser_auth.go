package handler

import (
	as "backend_course/customer_api_gateway/genproto/auth_service"
	"backend_course/customer_api_gateway/pkg/validator"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Router		/v1/systemuser/register [post]
// @Summary		Register a systemuser
// @Description	This api register a systemuser
// @Tags		SystemUser_Auth
// @Accept		json
// @Produce		json
// @Param		systemuser body auth_service.SystemUserGmailCheckRequest true "systemuser"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) RegisterSystemUser(c *gin.Context) {
	req := &as.SystemUserGmailCheckRequest{}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	if !validator.ValidateGmail(req.Gmail) {
		handleGrpcErrWithDescription(c, h.log,errors.New("wrong gmail"), "error while validating gmail")
		return
	}

	resp, err := h.grpcClient.SystemUserAuthService().SystemUserRegisterByMail(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while register systemuser", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while register systemuser")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/systemuser/registercomfirm [post]
// @Summary		Register Confirm a systemuser
// @Description	This api register confirm a systemuser
// @Tags		SystemUser_Auth
// @Accept		json
// @Produce		json
// @Param		systemuser body auth_service.SystemUserRConfirm true "systemuser"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) RegisterConfirmSystemUser(c *gin.Context) {
	req := &as.SystemUserRConfirm{}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.SystemUserAuthService().SystemUserRegisterByMailConfirm(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while register confirm systemuser", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while register confirm systemuser")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/systemuser/login [post]
// @Summary		Login  a systemuser
// @Description	This api login a systemuser
// @Tags		SystemUser_Auth
// @Accept		json
// @Produce		json
// @Param		systemuser body auth_service.SystemUserGmailCheckRequest true "systemuser"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) LoginSystemUser(c *gin.Context) {
	req := &as.SystemUserGmailCheckRequest{}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.SystemUserAuthService().SystemUserLoginByGmail(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while login systemuser", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while login systemuser")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/systemuser/loginconfirm [post]
// @Summary		Login confirm a systemuser
// @Description	This api login confirm a systemuser
// @Tags		SystemUser_Auth
// @Accept		json
// @Produce		json
// @Param		systemuser body auth_service.SystemUserRConfirm true "systemuser"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) SystemUserLoginConfirm(c *gin.Context) {
	req := &as.SystemUserRConfirm{}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.SystemUserAuthService().SystemUserRegisterByMailConfirm(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while login confirm systemuser", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while login confirm systemuser")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/systemuser/resetpassword [put]
// @Summary		Reset a password a systemuser
// @Description	This api reset a password a systemuser
// @Tags		SystemUser_Auth
// @Accept		json
// @Produce		json
// @Param		systemuser body auth_service.SystemUserGmailCheckRequest true "systemuser"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) ResetPasswordSystemUser(c *gin.Context) {
	req := &as.SystemUserGmailCheckRequest{}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.SystemUserAuthService().SystemUserResetPassword(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while reset password systemuser", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while login systemuser")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/systemuser/resetconfirm [put]
// @Summary		Confirm reset password a systemuser
// @Description	This api reset password confirm a systemuser
// @Tags		SystemUser_Auth
// @Accept		json
// @Produce		json
// @Param		systemuser body auth_service.SystemUserPasswordConfirm true "systemuser"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) SystemUserResetPasswordConfirm(c *gin.Context) {
	req := &as.SystemUserPasswordConfirm{}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.SystemUserAuthService().SystemUserResetPasswordConfirm(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while reset password confirm systemuser", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while reset password confirm systemuser")
		return
	}

	c.JSON(http.StatusOK, resp)
}