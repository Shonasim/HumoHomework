package main

import (
	"fmt"
	"math"
)

func deposit(balance *int, amount int) {
	*balance += amount
}
func checkBalance(balance *int) bool {
	fmt.Printf("Текущий баланс: %d рублей\n", *balance)
	if *balance == 500000 {
		fmt.Println("Достигнут лимит накоплений")
		return false
	}
	return true
}

// Задача2
func monthlyPayment(initialAmount *float64, year int) float64 {
	r := 10.0 / (12 * 100)
	n := year * 12
	compoundFactor := math.Pow(1+r, float64(n))
	M := (*initialAmount * r * compoundFactor) / (compoundFactor - 1)
	return M
}
func checkAmount(initialAmount *float64) bool {
	if *initialAmount < 500000 {
		fmt.Println("Кредит почти погашен")
		return false
	}
	return true
}
func main() {
	/*balance := 10000
	depositAmount := 50000
	for {
		deposit(&balance, depositAmount)
		if !checkBalance(&balance) {
			break
		}
	}*/
	//Задача2
	var initial_amount float64 = 3000000.0
	var year int
	fmt.Scan(&year)
	x := monthlyPayment(&initial_amount, year)
	//fmt.Println(monthlyPayment(&initial_amount, year))
	//fmt.Println(initial_amount)
	for i := 0; i < year; i++ {
		for j := 0; j < 12; j++ {
			initial_amount -= x
			fmt.Println("Остаток по кредиту после каждой выплаты", initial_amount)
			checkAmount(&initial_amount)
		}
	}

}
