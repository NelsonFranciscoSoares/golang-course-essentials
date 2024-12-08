package main

import (
	"example/hello/printer"
	"fmt"

	"rsc.io/quote/v4"
)

func main() {
	fmt.Println("Hello, Go Essentials!!!")
	fmt.Printf("Valor 10 em hex, 0x%X\n", 10)
	fmt.Println(quote.Go())
	printer.Print()
}
