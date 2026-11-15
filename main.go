package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"warehouse/warehouse"
)

func main() {
	store := warehouse.NewWarehouse()
	store.AddProduct(warehouse.Product{

		Name:  "krem",
		Price: 1012,
		Stock: 10,
	})
	fmt.Println(store)
	scanner := bufio.NewScanner(os.Stdin)
	for {

		fmt.Println("1 — список товаров")
		fmt.Println("2 — добавить товар")
		fmt.Println("3 — удалить товар по ID")
		fmt.Println("4 — продать товар (уменьшить остаток)")
		fmt.Println("5 — поменять цену у товара")
		fmt.Println("0 — выход")
		fmt.Print(">> ")
		if !scanner.Scan() {
			return
		}
		text := strings.TrimSpace(scanner.Text())
		query, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("Вы ввели не число")
			continue
		}

		switch query {
		case 1:

			products := store.ListProducts()
			if len(products) == 0 {
				fmt.Println("Товаров нет!")
			}
			for _, p := range products {
				fmt.Printf("id: %d,  %s  %d RUB, in stock: %d\n", p.ID, p.Name, p.Price, p.Stock)
			}
		case 2:

			newP, err := warehouse.CreateProduct(scanner)
			if err != nil {
				fmt.Println(err)
				continue
			}
			store.AddProduct(warehouse.Product{
				Name:  newP.Name,
				Price: newP.Price,
				Stock: newP.Stock,
			})
			fmt.Println("Товар добавлен успешно")
		case 3:
			fmt.Print("Введите айди товара для удаления \n>> ")
			if !scanner.Scan() {
				return
			}
			id, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				fmt.Println("Введите ЧИСЛО")
				continue
			}
			err = store.DeleteProduct(id)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("Товар удален успешно")
		case 4:
			fmt.Print("Введите айди товара для продажи \n>> ")
			if !scanner.Scan() {
				return
			}
			id, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				fmt.Println("Введите ЧИСЛО")
				continue
			}

			fmt.Print("Введите количетсво штук продукта для продажи \n>> ")
			if !scanner.Scan() {
				return
			}
			amount, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				fmt.Println("Введите ЧИСЛО")
				continue
			}
			err = store.SellProduct(id, amount)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("Товар продан успешно")
		case 5:
			fmt.Print("Введите айди у которого хотите поменять цену \n>> ")
			if !scanner.Scan() {
				return
			}
			id, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				fmt.Println("Введите ЧИСЛО")
				continue
			}
			fmt.Print("Введите новую цену \n>> ")
			if !scanner.Scan() {
				return
			}
			amount, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				fmt.Println("Введите ЧИСЛО")
				continue
			}
			err = store.EditProductPrice(id, amount)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("Цена изменена успешно!")
		case 0:
			return
		default:
			fmt.Println("Такого пункта нет!")

		}
	}
}
