package main

import (
	"fmt"
	"go-tests/addresses"
)

func main() {
	typeAddress := addresses.AddressType("Street dos bobos")
	// typeAddress := addresses.AddressType("abc dos bobos") // output Invalid Type

	fmt.Println(typeAddress)
}
