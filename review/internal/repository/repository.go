package repository

import "coupon_service/internal/types"

type Repository interface {
	FindByCode(string) (*types.Coupon, error)
	Save(*types.Coupon) error
}