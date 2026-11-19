package warehouse

import (
	"encoding/json"
	"errors"

	"os"
)

var ErrIdNotFound = errors.New("Товара с таким id нет")
var ErrCreateProduct = errors.New("Не удалось создать товар")
var ErrDelete = errors.New("Не удалось удалить товар")
var ErrEditStock = errors.New("Не удалось продать товар")
var ErrEditPrice = errors.New("Не удалось поменять цену у товара")
var ErrSaveToFile = errors.New("Невалидное название файла")

type Product struct {
	ID    int
	Name  string
	Price int
	Stock int
}

type Warehouse struct {
	products []Product
	nextId   int
}

func NewWarehouse() *Warehouse {
	return &Warehouse{
		products: make([]Product, 0),
		nextId:   1,
	}
}

func (w *Warehouse) AddProduct(pr Product) {
	pr.ID = w.nextId
	w.products = append(w.products, pr)
	w.nextId++
}

func (w *Warehouse) ListProducts() []Product {
	dst := make([]Product, len(w.products))
	copy(dst, w.products)
	return dst
}

func (w *Warehouse) DeleteProduct(id int) error {
	for i, p := range w.products {
		if p.ID == id {
			w.products = append(w.products[:i], w.products[i+1:]...)
			return nil
		}
	}
	return ErrIdNotFound
}

func (w *Warehouse) SellProduct(id int, amount int) error {

	for i, p := range w.products {
		if p.ID == id {
			if amount <= 0 || amount > w.products[i].Stock {
				return ErrEditStock
			}
			w.products[i].Stock = w.products[i].Stock - amount
			return nil
		}
	}

	return ErrIdNotFound
}

func (w *Warehouse) EditProductPrice(id int, amount int) error {
	if amount <= 0 {
		return ErrEditPrice
	}
	for i, p := range w.products {
		if id == p.ID {
			w.products[i].Price = amount
			return nil
		}
	}
	return ErrIdNotFound
}

func (w *Warehouse) SaveToFile(filename string) error {
	if filename == "" {
		return ErrSaveToFile
	}
	data, err := json.MarshalIndent(w.products, "", " ")
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (w *Warehouse) LoadFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &w.products)
	if err != nil {
		return err
	}
	for _, pr := range w.products {
		if pr.ID >= w.nextId {
			w.nextId = pr.ID + 1
		}
	}
	return nil
}
