package main

import (
	"fmt"
)

func NamaOrang(n string) string {
	var name string
	name += "Hello " + n + "! Saya Golang bahasa yang sangat Menyenangkan"

	return name
}

func main() {
	fmt.Println(NamaOrang("Kobar"))
	fmt.Println(NamaOrang("lukman"))
	fmt.Println(NamaOrang("alif"))
	fmt.Println(NamaOrang("al"))
	fmt.Println(NamaOrang("moh"))
}
