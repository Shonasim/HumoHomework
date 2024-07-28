package main

import "fmt"

// Задача 16
type Product struct {
	name  string
	price int
}

func (p Product) show() {
	fmt.Println("name", p.name, "price", p.price)
}

// Задача 17
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) show1() {
	fmt.Println(p.FirstName, p.LastName, p.Age)
}

// Задача 18
func Expensive(p1, p2 Product) Product {
	if p1.price > p2.price {
		return p1
	}
	return p2
}

func main() {
	//Задача 16
	p := Product{"kakao", 22}
	p.show()

	//Задача 17
	p1 := Person{"Shonasim", "Raimdodov", 21}
	p1.show1()

	//Задача 18
	p2 := Product{"smetana", 16}
	p3 := Product{"Yogurt", 8}
	fmt.Println(Expensive(p2, p3))
}
