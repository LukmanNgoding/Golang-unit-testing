package main

import "fmt"

func CariLuas(t float64, r float64) float64 {
	L := 2 * 3.14 * r * (r + t)

	return L
}

func main() {
	fmt.Println(CariLuas(20, 4))
	fmt.Println(CariLuas(19, 3))
	fmt.Println(CariLuas(18, 1))
	fmt.Println(CariLuas(17, 3))
	fmt.Println(CariLuas(20, 1))
}
