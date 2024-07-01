package memdb

import (
	"coupon_service/internal/types"
	"fmt"
	"sync"
)

type MemRepository struct {
	mu      sync.RWMutex
	entries map[string]*types.Coupon
}

func New() *MemRepository {
	return &MemRepository{
		entries: make(map[string]*types.Coupon),
	}
}

func (r *MemRepository) FindByCode(code string) (*types.Coupon, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	coupon, ok := r.entries[code]
	if !ok {
		return nil, fmt.Errorf("coupon with code %s does not exist", code)
	}
	return coupon, nil
}

func (r *MemRepository) Save(coupon *types.Coupon) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.entries[coupon.Code] = coupon
	return nil
}
