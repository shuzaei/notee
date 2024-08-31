package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	p := basicChordPattern
	s := p.Extend(RoyalRoadC)

	b, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}

	os.Stdout.Write(b)
}
