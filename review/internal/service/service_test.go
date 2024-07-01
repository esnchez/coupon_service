package service

import (
	"coupon_service/internal/repository"
	"coupon_service/internal/repository/memdb"
	"coupon_service/internal/types"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		repo repository.Repository
	}
	tests := []struct {
		name string
		args args
		want CouponService
	}{
		{"initialize service", args{repo: nil}, CouponService{repo: nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_ApplyCoupon(t *testing.T) {
	type fields struct {
		repo repository.Repository
	}
	type args struct {
		basket *types.Basket
		code   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantB   *types.Basket
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := CouponService{
				repo: tt.fields.repo,
			}
			gotB, err := s.ApplyCoupon(tt.args.basket, tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApplyCoupon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotB, tt.wantB) {
				t.Errorf("ApplyCoupon() gotB = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func TestService_CreateCoupon(t *testing.T) {
	type fields struct {
		repo repository.Repository
	}
	type args struct {
		discount       int
		code           string
		minBasketValue int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   any
	}{
		{"Apply 10%", fields{memdb.New()}, args{10, "Superdiscount", 55}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := CouponService{
				repo: tt.fields.repo,
			}
			req := &types.CreateCouponRequest{
				Discount:       tt.args.discount,
				Code:           tt.args.code,
				MinBasketValue: tt.args.minBasketValue,
			}
			s.CreateCoupon(req)
		})
	}
}
