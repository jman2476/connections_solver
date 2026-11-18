package main

import "fmt"

var (
	items = []string{
		"honey",
		"yellow",
		"venom",
		"juicy",
		"comb",
		"buzz",
		"sour",
		"ovoid",
	}
)

func generate(list []string) Options {
	var options Options
	fmt.Printf("Length of list: %d\n", len(list))
	for i := 0; i < len(list)-3; i++ {
		for j := i + 1; j < len(list)-2; j++ {
			for k := j + 1; k < len(list)-1; k++ {
				for l := k + 1; l < len(list); l++ {
					options = append(options, Category{
						list[i], list[j], list[k], list[l],
					})
				}
			}
		}
	}

	return options
}
