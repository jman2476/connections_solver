package main

import "fmt"

func main() {
	opts := generate(items)

	filt := filter(opts)

	for _, f := range pair(filt) {
		fmt.Println(f)
	}
}
