package handler

import (
	us "backend_course/customer_api_gateway/genproto/user_service"
	"backend_course/customer_api_gateway/pkg/validator"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Router		/v1/branch/create [post]
// @Summary		Creates a Branch
// @Description	This api creates a Branch and returns its id
// @Tags		Branch
// @Accept		json
// @Produce		json
// @Param		Branch body user_service.CreateBranch true "Branch"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) CreateBranch(c *gin.Context) {
	var req us.CreateBranch

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	if !validator.ValidatePhone(req.Phone){
		handleGrpcErrWithDescription(c, h.log, errors.New("error while validating phone"), "error while validating phone body")
		return
	}

	resp, err := h.grpcClient.BranchBranch().Create(c.Request.Context(), &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while create branch")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/branch/getbyid [post]
// @Summary		Get by id a Branch
// @Description	This api get by id a Branch
// @Tags		Branch
// Accept		json
// @Produce		json
// @Param		Branch body user_service.BranchPrimaryKey true "Branch"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetByID(c *gin.Context) {
	id := &us.BranchPrimaryKey{}

	if err := c.ShouldBindJSON(&id); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.BranchBranch().GetByID(c.Request.Context(), id)
	if err != nil {
		log.Fatal("error while get by id", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while get by id branch")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/branch/getlist [post]
// @Summary		Get list a Branch
// @Description	This api get list a Branch
// @Tags		Branch
// Accept		json
// @Produce		json
// @Param		Branch body user_service.GetListBranchRequest true "Branch"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetList(c *gin.Context) {
	req := &us.GetListBranchRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.BranchBranch().GetList(c.Request.Context(), req)
	if err != nil {
		log.Fatal("error while get list ", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while get list branch")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/branch/updatebranch [PUT]
// @Summary		Update a Branch
// @Description	This API updates a Branch
// @Tags		Branch
// @Accept		json
// @Produce		json
// @Param		Branch body user_service.UpdateBranchRequest true "Branch object to update"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) Update(c *gin.Context) {
	req := &us.UpdateBranchRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding binding body")
		return
	}

	resp, err := h.grpcClient.BranchBranch().Update(c.Request.Context(), req)
	if err != nil {
		log.Fatal("error while update Branch", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while updating branch")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/branch/delete [delete]
// @Summary		delete a Branch
// @Description	This api delete a Branch
// @Tags		Branch
// Accept		json
// @Produce		json
// @Param		Branch body user_service.BranchPrimaryKey true "Branch"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) Delete(c *gin.Context) {
	id := &us.BranchPrimaryKey{}

	if err := c.ShouldBindJSON(&id); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.BranchBranch().Delete(c.Request.Context(), id)
	if err != nil {
		log.Fatal("error while get delete", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while deleting branch")
		return
	}

	c.JSON(http.StatusOK, resp)
}
