package main

import "fmt"

func main() {
	opts := generate(items)

	if len(opts) != 70 {
		fmt.Printf("Found %d, expected %d\n", len(opts), 70)
	}

	for _, c := range opts {
		fmt.Println(c)
	}
}
