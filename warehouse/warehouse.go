package warehouse

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrCreateProduct = errors.New("Не удалось создать товар")
var ErrDelete = errors.New("Не удалось удалить товар")
var ErrEditStock = errors.New("Не удалось продать товар")
var ErrEditPrice = errors.New("Не удалось поменять цену у товара")

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

	return w.products
}

func CreateProduct(scanner *bufio.Scanner) (Product, error) {

	fmt.Print("ВВедите название: \n>> ")
	if !scanner.Scan() {
		return Product{}, ErrCreateProduct
	}
	name := strings.TrimSpace(scanner.Text())
	if len(name) == 0 {
		return Product{}, ErrCreateProduct
	}
	fmt.Print("ВВедите цену: \n>> ")
	if !scanner.Scan() {
		return Product{}, ErrCreateProduct
	}
	price, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil || price <= 0 {
		fmt.Println("Ошибка в вводе!")
		return Product{}, ErrCreateProduct
	}

	fmt.Print("ВВедите количество: \n>> ")
	if !scanner.Scan() {
		return Product{}, ErrCreateProduct
	}
	stock, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil || stock <= 0 {
		fmt.Println("Ошибка в вводе!")
		return Product{}, ErrCreateProduct
	}
	return Product{Name: name, Price: price, Stock: stock}, nil
}

func (w *Warehouse) DeleteProduct(id int) error {
	for i, p := range w.products {
		if p.ID == id {
			w.products = append(w.products[:i], w.products[i+1:]...)
			return nil
		}
	}
	return ErrDelete
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

	return ErrEditStock
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
	return ErrEditPrice
}
