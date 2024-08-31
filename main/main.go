package main

import (
	"fmt"
	"notee/notee"
)

func main() {
	// test

	fmt.Println("C Scale")
	for _, k := range notee.CScale.Keys {
		fmt.Println(notee.CScale.KeyToTranslation(k))
	}
}
