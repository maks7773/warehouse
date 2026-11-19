package warehouse_test

import (
	"errors"
	"testing"
	wh "warehouse/warehouse"
)

func TestAdd(t *testing.T) {
	warehouse := wh.NewWarehouse()
	warehouse.AddProduct(wh.Product{
		Name:  "Product 1",
		Price: 123,
		Stock: 2,
	})
	warehouse.AddProduct(wh.Product{
		Name:  "Product 2",
		Price: 321,
		Stock: 3,
	})
	cart := warehouse.ListProducts()
	if len(cart) != 2 {
		t.Fatalf("len(cart)=%d", len(cart))
	}
	if cart[0].Name != "Product 1" || cart[1].Name != "Product 2" {
		t.Errorf("Ошибка в имени продуктов")
	}
	if cart[0].Price != 123 || cart[1].Price != 321 {
		t.Errorf("Ошибка в ценах товаров")
	}
	if cart[0].Stock != 2 || cart[1].Stock != 3 {
		t.Errorf("Ошибка в количестве товаров")
	}
	if cart[0].ID != 1 || cart[1].ID != 2 {
		t.Errorf("Ошибка в айди товаров")
	}

}

func TestSellProduct(t *testing.T) {
	warehouse := wh.NewWarehouse()
	warehouse.AddProduct(wh.Product{
		Name:  "Product 1",
		Price: 123,
		Stock: 10,
	})
	err := warehouse.SellProduct(1, 3)
	if err != nil {
		t.Fatalf("не ожидали ошибку, но получили: %v", err)
	}
	cart := warehouse.ListProducts()
	if cart[0].Stock != 7 {
		t.Errorf("остаток стал %d а ожидали 7", cart[0].Stock)
	}

	err = warehouse.SellProduct(1, 100)
	if !errors.Is(err, wh.ErrEditStock) {
		t.Errorf("ожидали ошибку ErrEditStock, но получили: %v", err)
	}

	err = warehouse.SellProduct(999, 2)
	if !errors.Is(err, wh.ErrIdNotFound) {
		t.Errorf("ожидали ошибку ErrIdNotFound, но получили: %v", err)
	}
}
