package services

import (
	"go-api-ddd/aggregate"
	"testing"

	"github.com/google/uuid"
)

func init_products(t *testing.T) []aggregate.Product {
	beer, err := aggregate.NewProduct(
		"Beer",
		"Healty beverage",
		20,
	)

	if err != nil {
		t.Fatal(err)
	}

	peenuts, err := aggregate.NewProduct(
		"Peanuts",
		"Snack",
		0.99,
	)

	if err != nil {
		t.Fatal(err)
	}

	wine, err := aggregate.NewProduct(
		"Wine",
		"Nasty Drink",
		0.99,
	)

	if err != nil {
		t.Fatal(err)
	}

	return []aggregate.Product{
		beer, peenuts, wine,
	}
}

func TestOrder_NewOrder(t *testing.T) {
	products := init_products(t)
	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)

	if err != nil {
		t.Fatal(err)
	}

	cust, err := aggregate.NewCustomer("Mawwl")
	if err != nil {
		t.Error(err)
	}

	err = os.customers.Add(cust)
	if err != nil {
		t.Error(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	price, err := os.CreateOrder(cust.GetID(), order)
	if err != nil {
		t.Error(err)
	}

	t.Logf("Price: %0.0f\n", price)
}