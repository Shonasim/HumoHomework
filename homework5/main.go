package main

import "fmt"

func main() {
	// 1. Вывод чисел от 1 до 10
	fmt.Println("Задача 1: Вывод чисел от 1 до 10")
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 2. Вывод квадратов чисел от 1 до 5
	fmt.Println("Задача 2: Вывод квадратов чисел от 1 до 5")
	for i := 1; i <= 5; i++ {
		fmt.Println(i * i)
	}

	// 3. Таблица умножения на 3 от 1 до 10
	fmt.Println("Задача 3: Таблица умножения на 3 от 1 до 10")
	for i := 1; i <= 10; i++ {
		fmt.Println("3 *", i, "=", 3*i)
	}

	// 4. Первые 10 чисел Фибоначчи
	fmt.Println("Задача 4: Первые 10 чисел Фибоначчи")
	a, b := 0, 1
	for i := 0; i < 10; i++ {
		fmt.Println(a)
		a, b = b, a+b
	}

	// 5. НОД двух чисел (алгоритм Евклида)
	fmt.Println("Задача 5: НОД двух чисел")
	var a1, b1 int
	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)
	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)
	for b != 0 {
		a1, b1 = b1, a1%b1
	}
	fmt.Println("НОД:", a)

	// 6. FizzBuzz
	fmt.Println("Задача 6: FizzBuzz")
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

	// 7. Проверка простого числа
	fmt.Println("Задача 7: Проверка простого числа")
	var number int
	fmt.Print("Введите число: ")
	fmt.Scanln(&number)
	isPrime := true
	if number <= 1 {
		isPrime = false
	} else {
		for i := 2; i*i <= number; i++ {
			if number%i == 0 {
				isPrime = false
				break
			}
		}
	}
	if isPrime {
		fmt.Println("Число простое")
	} else {
		fmt.Println("Число не простое")
	}

	// 8. Вывод делителей числа
	fmt.Println("Задача 8: Вывод делителей числа")
	fmt.Print("Введите число: ")
	fmt.Scanln(&number)
	fmt.Println("Делители:", 1)
	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			fmt.Println(i)
			if i*i != number {
				fmt.Println(number / i)
			}
		}
	}

	// 9. Сумма цифр числа
	fmt.Println("Задача 9: Сумма цифр числа")
	fmt.Print("Введите число: ")
	fmt.Scanln(&number)
	sum := 0
	for number > 0 {
		sum += number % 10
		number /= 10
	}
	fmt.Println("Сумма цифр:", sum)

	// 10. Ввод положительного числа
	fmt.Println("Задача 10: Ввод положительного числа")
	for {
		fmt.Print("Введите положительное число: ")
		fmt.Scanln(&number)
		if number > 0 {
			break
		}
		fmt.Println("Число должно быть положительным.")
	}

	// 11. Произведение чисел (прерывание при превышении 1000)
	fmt.Println("Задача 11: Произведение чисел")
	fmt.Print("Введите число: ")
	fmt.Scanln(&number)
	product := 1
	for i := 1; i <= number; i++ {
		product *= i
		if product > 1000 {
			fmt.Println("Произведение превысило 1000")
			break
		}
	}
	fmt.Println("Произведение:", product)

	// 12. Палиндром
	fmt.Println("Задача 12: Палиндром")
	var numberStr string
	fmt.Print("Введите число: ")
	fmt.Scanln(&numberStr)
	isPalindrome := true
	for i := 0; i < len(numberStr)/2; i++ {
		if numberStr[i] != numberStr[len(numberStr)-i-1] {
			isPalindrome = false
			break
		}
	}
	if isPalindrome {
		fmt.Println("Число - палиндром")
	} else {
		fmt.Println("Число - не палиндром")
	}

	// 13. Пирамида из '*'
	fmt.Println("Задача 13: Пирамида из '*'")
	fmt.Print("Введите высоту пирамиды: ")
	fmt.Scanln(&number)
	for i := 1; i <= number; i++ {
		for j := 1; j <= number-i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// 14. Таблица умножения
	fmt.Println("Задача 14: Таблица умножения")
	fmt.Print("Введите число: ")
	fmt.Scanln(&number)
	for i := 1; i <= number; i++ {
		for j := 1; j <= number; j++ {
			fmt.Printf("%3d ", i*j)
		}
		fmt.Println()
	}

	// 15. Треугольник Паскаля
	fmt.Println("Задача 15: Треугольник Паскаля")
	fmt.Print("Введите высоту треугольника: ")
	fmt.Scanln(&number)
	for i := 0; i < number; i++ {
		for j := 0; j < number-i; j++ {
			fmt.Print("  ")
		}
		for j := 0; j <= i; j++ {
			var value int
			if j == 0 || j == i {
				value = 1
			} else {
				value = 1
				for k := 1; k < j; k++ {
					value *= (i - k + 1)
					value /= k
				}
			}
			fmt.Printf("%3d ", value)
		}
		fmt.Println()
	}

	// 16. Факториал (бесконечный цикл до отрицательного числа)
	fmt.Println("Задача 16: Факториал (бесконечный цикл)")
	for {
		fmt.Print("Введите число: ")
		fmt.Scanln(&number)
		if number < 0 {
			break
		}
		factorial := 1
		for i := 1; i <= number; i++ {
			factorial *= i
		}
		fmt.Println("Факториал:", factorial)
	}

	// 17. Сумма двух чисел (бесконечный цикл)
	fmt.Println("Задача 17: Сумма двух чисел (бесконечный цикл)")
	for {
		var a, b int
		fmt.Print("Введите первое число: ")
		fmt.Scanln(&a)
		fmt.Print("Введите второе число: ")
		fmt.Scanln(&b)
		fmt.Println("Сумма:", a+b)
	}

	// 18. Вывод чисел, не являющихся квадратами
	fmt.Println("Задача 18: Вывод чисел, не являющихся квадратами")
	for i := 1; i <= 100; i++ {
		isSquare := false
		for j := 1; j*j <= i; j++ {
			if j*j == i {
				isSquare = true
				break
			}
		}
		if !isSquare {
			fmt.Println(i)
		}
	}

	// 19. Вывод простых чисел от 1 до 50
	fmt.Println("Задача 19: Вывод простых чисел от 1 до 50")
	for i := 2; i <= 50; i++ {
		isPrime := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println(i)
		}
	}

	// 20. Вывод чисел, не делящихся на 2 или 3
	fmt.Println("Задача 20: Вывод чисел, не делящихся на 2 или 3")
	for i := 1; i <= 30; i++ {
		if i%2 != 0 && i%3 != 0 {
			fmt.Println(i)
		}
	}

	// 21. Вывод чисел до куба
	fmt.Println("Задача 21: Вывод чисел до куба")
	for i := 1; i <= 50; i++ {
		if i*i*i > 50 {
			break
		}
		fmt.Println(i)
	}

	// 22. Вывод чисел с ограничением суммы
	fmt.Println("Задача 22: Вывод чисел с ограничением суммы")
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
		fmt.Println(i)
		if sum > 200 {
			break
		}
	}

	// 23. Бесконечный цикл с прерыванием при делении на 5
	fmt.Println("Задача 23: Бесконечный цикл с прерыванием")
	for {
		fmt.Print("Введите число: ")
		fmt.Scanln(&number)
		if number%5 == 0 {
			break
		}
	}
}
