package service

import (
	"coupon_service/internal/repository"
	"coupon_service/internal/types"
	"fmt"
)

type Service interface {
	ApplyCoupon(*types.Basket, string) (*types.Basket, error)
	CreateCoupon(*types.CreateCouponRequest) error
	GetCoupons([]string) ([]*types.Coupon, error)
}

type CouponService struct {
	repo repository.Repository
}

func New(repo repository.Repository) *CouponService {
	return &CouponService{
		repo: repo,
	}
}

func (s *CouponService) ApplyCoupon(b *types.Basket, code string) (*types.Basket, error) {

	if err := b.Validate(); err != nil {
		return nil, err
	}
	
	coupon, err := s.repo.FindByCode(code)
	if err != nil {
		return nil, err
	}

	if b.Value < coupon.MinBasketValue {
		return nil, fmt.Errorf("tried to apply discount to an invalid basket value")
	}
	
	b.Value -= coupon.Discount
	b.AppliedDiscount = coupon.Discount
	*b.ApplicationSuccessful = true
	return b, nil
}

func (s *CouponService) CreateCoupon(req *types.CreateCouponRequest) error {

	coupon, err := types.NewCoupon(req)
	if err != nil {
		return err
	}

	if err := s.repo.Save(coupon); err != nil {
		return err
	}
	return nil
}

func (s *CouponService) GetCoupons(codes []string) ([]*types.Coupon, error) {
	coupons := make([]*types.Coupon, 0, len(codes))
	var coupon *types.Coupon
	var err error

	for _, code := range codes {
		coupon, err = s.repo.FindByCode(code)
		if err != nil {
			return nil, err
		}
		coupons = append(coupons, coupon)
	}

	return coupons, nil
}
