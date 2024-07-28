package main

import "fmt"

func PrintGreeting() {
	var dayType string
	fmt.Print("Введите время суток (утро, день, вечер): ")
	fmt.Scanln(&dayType)

	switch dayType {
	case "утро":
		fmt.Println("Доброе утро!")
	case "день":
		fmt.Println("Добрый день!")
	case "вечер":
		fmt.Println("Добрый вечер!")
	default:
		fmt.Println("Некорректный ввод.")
	}
}

func PrintWeather() {
	var weatherType string
	fmt.Print("Введите погоду (солнечно, облачно, дождливо): ")
	fmt.Scanln(&weatherType)

	switch weatherType {
	case "солнечно":
		fmt.Println("Солнечно")
	case "облачно":
		fmt.Println("Облачно")
	case "дождливо":
		fmt.Println("Дождливо")
	default:
		fmt.Println("Некорректный ввод.")
	}
}

func PrintTrafficLight() {
	var lightColor string
	fmt.Print("Введите цвет светофора (красный, желтый, зеленый): ")
	fmt.Scanln(&lightColor)

	switch lightColor {
	case "красный":
		fmt.Println("Стоп")
	case "желтый":
		fmt.Println("Внимание")
	case "зеленый":
		fmt.Println("Идите")
	default:
		fmt.Println("Некорректный ввод.")
	}
}

func GetGrade() string {
	var score int
	fmt.Print("Введите балл: ")
	fmt.Scanln(&score)

	if score >= 90 {
		return "A"
	} else if score >= 80 {
		return "B"
	} else if score >= 70 {
		return "C"
	} else {
		return "Неудовлетворительно"
	}
}

func GetDiscount() string {
	var amount int
	fmt.Print("Введите сумму покупки: ")
	fmt.Scanln(&amount)

	if amount > 1000 {
		return "10%"
	} else if amount > 500 {
		return "5%"
	} else {
		return "0%"
	}
}

func GetTemperatureDescription() string {
	var temperature int
	fmt.Print("Введите температуру: ")
	fmt.Scanln(&temperature)

	if temperature < 10 {
		return "Холодно"
	} else if temperature <= 25 {
		return "Тепло"
	} else {
		return "Жарко"
	}
}

func CheckNumber(number int) {
	if number > 0 {
		fmt.Println("Положительное")
	} else if number < 0 {
		fmt.Println("Отрицательное")
	} else {
		fmt.Println("Ноль")
	}
}

func CheckAge(age int) {
	if age >= 18 {
		fmt.Println("Совершеннолетний")
	} else {
		fmt.Println("Несовершеннолетний")
	}
}

func CheckPassword(password string) {
	if password == "1234" {
		fmt.Println("Пароль верный")
	} else {
		fmt.Println("Пароль неверный")
	}
}

func Add(a int, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	return a + b
}

func CompareStrings(str1 string, str2 string) string {
	if str1 == str2 {
		return "равны"
	} else {
		return "не равны"
	}
}

func main() {
	// Вызов функций
	PrintGreeting()
	PrintWeather()
	PrintTrafficLight()
	fmt.Println(GetGrade())
	fmt.Println(GetDiscount())
	fmt.Println(GetTemperatureDescription())
	CheckNumber(5)
	CheckAge(17)
	CheckPassword("1234")
	fmt.Println(Add(5, -3))
	fmt.Println(CompareStrings("Hello", "Hello"))
}
