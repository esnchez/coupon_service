package types

import "fmt"

type Basket struct {
	Value                 int   `json:"value" binding:"required,numeric,gte=0"`
	AppliedDiscount       int   `json:"applied_discount"`
	ApplicationSuccessful *bool `json:"application_succesful"  binding:"required"`
}

type ApplicationRequest struct {
	Code   string  `json:"code" binding:"required"`
	Basket *Basket `json:"basket" binding:"required"`
}

func (b *Basket) Validate() error {
	if *b.ApplicationSuccessful {
		return fmt.Errorf("tried to apply discount to an already discounted basket")
	}
	return nil
}
