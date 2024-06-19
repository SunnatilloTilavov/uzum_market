package handler

import (
	us "backend_course/customer_api_gateway/genproto/user_service"
	"backend_course/customer_api_gateway/pkg/validator"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Router		/v1/system-user/create [post]
// @Summary		Creates a system-user
// @Description	This api creates a system-user and returns its id
// @Tags		SystemUser
// @Accept		json
// @Produce		json
// @Param		system-user body user_service.CreateSystemUser true "system-user"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) CreateSystemUser(c *gin.Context) {
	req := &us.CreateSystemUser{}

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

	resp, err := h.grpcClient.SystemUserService().Create(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while get create system-user", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while creating system-user")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/system-user/getbyid [post]
// @Summary		Get by id a system-user
// @Description	This api get by id a system-user
// @Tags		SystemUser
// Accept		json
// @Produce		json
// @Param		system-user body user_service.SystemUserPrimaryKey true "system-user"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetByIdSystemUser(c *gin.Context) {
	id := &us.SystemUserPrimaryKey{}

	if err := c.ShouldBindJSON(&id); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.SystemUserService().GetByID(c.Request.Context(), id)
	if err != nil {
		fmt.Errorf("error while get delete", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while getting by id")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/system-user/getlist [post]
// @Summary		Get list a system-user
// @Description	This api get list a system-user
// @Tags		SystemUser
// Accept		json
// @Produce		json
// @Param		system-user body user_service.GetListSystemUserRequest true "system-user"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetListSystemUser(c *gin.Context) {
	req := &us.GetListSystemUserRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.SystemUserService().GetList(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while get list", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while get list")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/system-user/update [PUT]
// @Summary		Update a system-user
// @Description	This API updates a system-user
// @Tags		SystemUser
// @Accept		json
// @Produce		json
// @Param		system-user body user_service.UpdateSystemUserRequest true "SystemUser object to update"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) UpdateSystemUser(c *gin.Context) {
	req := &us.UpdateSystemUserRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while update system-user")
		return
	}

	resp, err := h.grpcClient.SystemUserService().Update(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while update system-user", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while ")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/system-user/delete [delete]
// @Summary		delete a system-user
// @Description	This api delete a system-user
// @Tags		SystemUser
// Accept		json
// @Produce		json
// @Param		system-user body user_service.SystemUserPrimaryKey true "system-user"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) DeleteSystemUser(c *gin.Context) {
	id := &us.SystemUserPrimaryKey{}

	if err := c.ShouldBindJSON(&id); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.SystemUserService().Delete(c.Request.Context(), id)
	if err != nil {
		fmt.Errorf("error while delete", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while delete system-user")
		return
	}

	c.JSON(http.StatusOK, resp)
}
