package main

import "fmt"

func main() {
	n := 234
	fmt.Println(subtractProductAndSum(n))
}

func subtractProductAndSum(n int) int {
	sum := 0
	product := 1
	substract := 0

	for n > 0 {
		a := n % 10
		sum = sum + a
		product = product * a
		n = n / 10
		substract = product - sum
	}

	return substract

}
