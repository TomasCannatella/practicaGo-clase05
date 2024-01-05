/*
Ejercicio 1
Crear un programa que cumpla los siguiente puntos:

Tener una estructura llamada Product con los campos ID, Name, Price, Description y Category.
Tener un slice global de Product llamado Products instanciado con valores.

2 métodos asociados a la estructura Product: Save(), GetAll(). El método Save()
deberá tomar el slice de Products y añadir el producto desde el cual se llama al método.
El método GetAll() deberá imprimir todos los productos guardados en el slice Products.

Una función getById() al cual se le deberá pasar un INT como parámetro y retorna el producto correspondiente al parámetro pasado.
Ejecutar al menos una vez cada método y función definido desde main().
*/
package main

import (
	product "ejercicio1/internal/Product"
	"fmt"
)

func main() {
	fmt.Println("Productos")
	p := product.Product{}

	p.GetAll()

	var newProduct product.Product = product.Product{
		Id:          3,
		Name:        "Mouse Logitech",
		Price:       25.6,
		Description: "Logitech g203",
		Category:    "Electronic",
	}

	p.Save(newProduct)

	p.GetAll()

	fmt.Println("Obteniendo el nuevo producto")
	fmt.Println(product.GetById(3))

}
