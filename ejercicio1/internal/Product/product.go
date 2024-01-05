package product

import (
	"fmt"
)

type Product struct {
	Id          int     //`json: "-"`
	Name        string  //`json : "Name"`
	Price       float64 //`json : "Price"`
	Description string  //`json : "Description"`
	Category    string  //`json : "Price,omitempty"`
}

var monitor Product = Product{Id: 1,
	Name:        "Monitor Samsung",
	Price:       100.00,
	Description: "Monitor Samsung",
	Category:    "Electronic"}

var headset Product = Product{Id: 2,
	Name:        "Headset Sony",
	Price:       50.00,
	Description: "Headset Sony",
	Category:    "Electronic"}

var Products = []Product{monitor, headset}

func (p Product) Save(product Product) {
	Products = append(Products, product)
	fmt.Println("Se guardo el elemento")
}

func (p Product) GetAll() {
	for _, product := range Products {
		fmt.Println(product)
	}
}

func GetById(id int) (p Product) {
	for _, product := range Products {
		if product.Id == id {
			p = product
		}
	}
	return
}
