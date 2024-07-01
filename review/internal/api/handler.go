package api

import (
	"coupon_service/internal/types"
	"net/http"

	"github.com/gin-gonic/gin"
)


func (a *API) Apply(c *gin.Context) {
	req := &types.ApplicationRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.Error(&CustomError{
			StatusCode: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	basket, err := a.svc.ApplyCoupon(req.Basket, req.Code)
	if err != nil {
		c.Error(&CustomError{
			StatusCode: http.StatusUnprocessableEntity,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, basket)
}

func (a *API) Create(c *gin.Context) {
	req := &types.CreateCouponRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.Error(&CustomError{
			StatusCode: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	if err := a.svc.CreateCoupon(req); err != nil {
		c.Error(&CustomError{
			StatusCode: http.StatusUnprocessableEntity,
			Message: err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}

func (a *API) Get(c *gin.Context) {
	req := &types.GetCouponRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.Error(&CustomError{
			StatusCode: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	coupons, err := a.svc.GetCoupons(req.Codes)
	if err != nil {
		c.Error(&CustomError{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, coupons)
}
