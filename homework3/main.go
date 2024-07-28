package main

import (
	"fmt"
	"math"
)

// Функции без аргументов и с возвращаемым значением

func GetWelcomeMessage() string {
	return "Welcome!"
}

func GetPiValue() float64 {
	return math.Pi
}

func IsServerActive() bool {
	return true
}

// Функции с аргументами, но без возвращаемого значения

func PrintGreeting() {
	fmt.Println("Hello, World!")
}

func DisplaySeparator() {
	for i := 0; i < 20; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}

func NotifyStart() {
	fmt.Println("Program started")
}

func PrintNumber(number int) {
	fmt.Println(number)
}

func GreetUser(name string) {
	fmt.Println("Hello,", name)
}

func ToggleLight(isOn bool) {
	if isOn {
		fmt.Println("Light on")
	} else {
		fmt.Println("Light off")
	}
}

// Функции с аргументами и возвращаемым значением

func Add(a int, b int) int {
	return a + b
}

func Concat(str1 string, str2 string) string {
	return str1 + str2
}

func IsEven(number int) bool {
	return number%2 == 0
}

// Функции как переменные

var adder func(int, int) int = Add
var concatenator func(string, string) string = Concat
var isPositive func(int) bool = func(number int) bool { return number > 0 }

// Функции как аргументы функций

func Calculate(a int, b int, operation func(int, int) int) int {
	return operation(a, b)
}

func Execute(condition bool, action func(bool)) {
	action(condition)
}

func Apply(number int, operation func(int) int) int {
	return operation(number)
}

// Функции как значения функций (callback)

func Multiplier(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

func StringRepeater(n int) func(string) string {
	return func(str string) string {
		result := ""
		for i := 0; i < n; i++ {
			result += str
		}
		return result
	}
}

func ConditionalPrinter(condition bool) func(string) {
	return func(str string) {
		if condition {
			fmt.Println(str)
		}
	}
}

func main() {
	// Вызов функций без аргументов и с возвращаемым значением
	fmt.Println(GetWelcomeMessage())
	fmt.Println(GetPiValue())
	fmt.Println(IsServerActive())

	// Вызов функций с аргументами, но без возвращаемого значения
	PrintGreeting()
	DisplaySeparator()
	NotifyStart()
	PrintNumber(10)
	GreetUser("John")
	ToggleLight(true)

	// Вызов функций с аргументами и возвращаемым значением
	fmt.Println(Add(5, 7))
	fmt.Println(Concat("Hello", " World!"))
	fmt.Println(IsEven(4))

	// Использование функций как переменных
	fmt.Println(adder(5, 7))
	fmt.Println(concatenator("Hello", " World!"))
	fmt.Println(isPositive(5))

	// Использование функций как аргументов функций
	fmt.Println(Calculate(5, 7, Add))
	Execute(true, func(condition bool) {
		fmt.Println("Action executed:", condition)
	})
	fmt.Println(Apply(5, func(number int) int {
		return number * 2
	}))

	// Использование функций как значений функций (callback)
	multiplyBy3 := Multiplier(3)
	fmt.Println(multiplyBy3(5))

	repeat3Times := StringRepeater(3)
	fmt.Println(repeat3Times("Hello"))

	printIfTrue := ConditionalPrinter(true)
	printIfTrue("This will be printed")

	printIfFalse := ConditionalPrinter(false)
	printIfFalse("This will not be printed")
}
