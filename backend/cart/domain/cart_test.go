package domain_test

import (
	"testing"

	"github.com/bkielbasa/go-ecommerce/backend/cart/domain"
	"github.com/matryer/is"
)

var pID = domain.NewProduct("productID", 2.0)
var pID2 = domain.NewProduct("productID2", 10.0)

func TestCart_Adding_Products_To_Cart_Should_Change_Quantity(t *testing.T) {
	is := is.New(t)
	// given
	c := domain.NewCart(newMockPrice())

	// when
	_ = c.Add(pID, 3)
	_ = c.Add(pID, 1)

	// then
	is.Equal(4, c.Quantity(pID.ID()))
}

func TestCart_Adding_Products_To_Cart_Should_Change_Quantity_Of_Single_Product(t *testing.T) {
	is := is.New(t)
	// given
	c := domain.NewCart(newMockPrice())
	_ = c.Add(pID, 1)

	// when
	_ = c.Add(pID2, 3)

	// then
	is.Equal(1, c.Quantity(pID.ID()))
	is.Equal(3, c.Quantity(pID2.ID()))
}

func TestCart_CalculateTotalPrice(t *testing.T) {
	is := is.New(t)
	// given
	c := domain.NewCart(newMockPrice())

	// when
	_ = c.Add(pID, 1)
	_ = c.Add(pID2, 3)

	// then
	is.Equal(4, c.TotalQuantity())
}

func TestPriceService_CalculatingTotalPrice(t *testing.T) {
	is := is.New(t)
	// given
	serv := domain.NewCart(newMockPrice())

	// when
	_ = serv.Add(pID, 3)
	total, err := serv.TotalPrice()

	// then
	is.NoErr(err)
	is.Equal(6.0, total)
}

func newMockPrice() mockPriceService {
	return mockPriceService{prices: map[string]float64{
		pID.ID(): 2.0,
	}}
}

type mockPriceService struct {
	prices map[string]float64
	err    error
}

func (ps mockPriceService) PriceFor(producID string) (float64, error) {
	return ps.prices[producID], ps.err
}
