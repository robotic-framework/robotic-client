package main

import "fmt"

func main() {
	var number uint16 = 5
	fmt.Printf("%016b\n", number)

	fmt.Printf("%016b\n", number&0x01)

	fmt.Printf("%016b\n", number&0x04)

	fmt.Printf("%X\n", 16)
}
