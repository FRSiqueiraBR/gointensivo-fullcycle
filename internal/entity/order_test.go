package entity

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_If_It_Get_An_Error_If_ID_Is_Blank(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "id is required")
}

func Test_If_It_Gets_An_Error_If_Price_Is_Blank(t *testing.T) {
	order := Order{ID: "123"}
	assert.Error(t, order.Validate(), "price must be greater than zero")
}

func Test_If_It_Gets_An_Error_If_Tax_Is_Blank(t *testing.T) {
	order := Order{ID: "123", Price: 10.0, Tax: 0}
	assert.Error(t, order.Validate(), "tax must be greater tha nor equal to zero")
}

func Test_Final_Price(t *testing.T) {
	order := Order{ID: "123", Price: 10.0, Tax: 5.50}
	assert.NoError(t, order.Validate())
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 5.50, order.Tax)
	order.CalculateFinalPrice()
	assert.Equal(t, 15.5, order.FinalPrice)
}