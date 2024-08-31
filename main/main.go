package main

import (
	"fmt"
	"notee/notee"
)

func main() {
	cMajor := notee.CMajor
	cMinor := notee.CMinor

	fmt.Println("CMajor:")
	fmt.Println(cMajor.ToString())

	fmt.Println("CMinor:")
	fmt.Println(cMinor.ToString())
}
