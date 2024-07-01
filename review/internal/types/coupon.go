package types

import (
	"fmt"

	"github.com/google/uuid"
)

type Coupon struct {
	ID             string `json:"id"`
	Code           string `json:"code"`
	Discount       int    `json:"discount"`
	MinBasketValue int    `json:"min_basket_value"`
}

type CreateCouponRequest struct {
	Code           string `json:"code" binding:"required"`
	Discount       int    `json:"discount" binding:"required,numeric,gte=0"`
	MinBasketValue int    `json:"min_basket_value" binding:"required,numeric,gte=0"`
}

type GetCouponRequest struct {
	Codes []string `json:"codes" binding:"required"`
}

func NewCoupon(req *CreateCouponRequest) (*Coupon, error) {

	if err := validateCreateCouponRequest(req); err != nil {
		return nil, err
	}
	return &Coupon{
		ID:             uuid.NewString(),
		Code:           req.Code,
		Discount:       req.Discount,
		MinBasketValue: req.MinBasketValue,
	}, nil
}

func validateCreateCouponRequest(req *CreateCouponRequest) error {
	if req.MinBasketValue <= req.Discount {
		return fmt.Errorf("tried to create coupon with invalid min basket value / discount relation")
	}
	return nil
}
