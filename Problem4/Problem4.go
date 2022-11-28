package main

import "fmt"

func SegiTiga(l float64, ting float64) float64 {
	res := 0.5 * l * ting

	//eksekusi
	return res
}

func main() {
	fmt.Println(SegiTiga(20, 25))
	fmt.Println(SegiTiga(21, 24))
	fmt.Println(SegiTiga(22, 20))
	fmt.Println(SegiTiga(23, 22))
	fmt.Println(SegiTiga(24, 21))

}
