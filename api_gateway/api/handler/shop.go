package handler

import (
	us "backend_course/customer_api_gateway/genproto/user_service"
	"backend_course/customer_api_gateway/pkg/validator"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Router		/v1/shop/create [post]
// @Summary		Creates a shop
// @Description	This api creates a shop and returns its id
// @Tags		Shop
// @Accept		json
// @Produce		json
// @Param		shop body user_service.CreateShop true "shop"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) CreateShop(c *gin.Context) {
	req := &us.CreateShop{}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	if !validator.ValidatePhone(req.Phone){
		handleGrpcErrWithDescription(c, h.log, errors.New("error while validating phone"), "wrong phone")
		return
	}


	resp, err := h.grpcClient.ShopService().Create(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while get create shop", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while create shop")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/shop/getbyid [post]
// @Summary		Get by id a shop
// @Description	This api get by id a shop
// @Tags		Shop
// Accept		json
// @Produce		json
// @Param		shop body user_service.ShopPrimaryKey true "shop"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetByIdShop(c *gin.Context) {
	id := &us.ShopPrimaryKey{}

	if err := c.ShouldBindJSON(&id); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.ShopService().GetById(c.Request.Context(), id)
	if err != nil {
		fmt.Errorf("error while get delete", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while delete")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/shop/getlist [post]
// @Summary		Get list a shop
// @Description	This api get list a shop
// @Tags		Shop
// Accept		json
// @Produce		json
// @Param		shop body user_service.GetListShopRequest true "shop"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetListShop(c *gin.Context) {
	req := &us.GetListShopRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.ShopService().GetList(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while get list", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while get list")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/shop/update [PUT]
// @Summary		Update a shop
// @Description	This API updates a shop
// @Tags		Shop
// @Accept		json
// @Produce		json
// @Param		shop body user_service.UpdateShopRequest true "Shop object to update"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) UpdateShop(c *gin.Context) {
	req := &us.UpdateShopRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while update shop")
		return
	}

	resp, err := h.grpcClient.ShopService().Update(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while update shop", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while ")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/shop/delete [delete]
// @Summary		delete a shop
// @Description	This api delete a shop
// @Tags		Shop
// Accept		json
// @Produce		json
// @Param		shop body user_service.ShopPrimaryKey true "shop"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) DeleteShop(c *gin.Context) {
	id := &us.ShopPrimaryKey{}

	if err := c.ShouldBindJSON(&id); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.ShopService().Delete(c.Request.Context(), id)
	if err != nil {
		fmt.Errorf("error while delete", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while delete shop")
		return
	}

	c.JSON(http.StatusOK, resp)
}
