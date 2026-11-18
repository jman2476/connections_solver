package main

import "slices"

func doesShare(items [2]string, cat Category) bool {
	for _, i := range items {
		if !slices.Contains(cat, i) {
			return false
		}
	}
	return true
}

func hasThree(items [4]string, cat Category) bool {
	count := 0

	for _, i := range items {
		if slices.Contains(cat, i) {
			count++
		}
	}

	return count >= 3
}

func filter(opts Options) Options {
	var filtered Options
	cantShare := [2]string{"yellow", "honey"}
	badCategory := [4]string{"honey", "juicy", "sour", "buzz"}

	for _, o := range opts {
		if hasThree(badCategory, o) {
			continue
		}
		if doesShare(cantShare, o) {
			continue
		}
		filtered = append(filtered, o)
	}

	return filtered
}

func pair(opts Options) []Pair {
	pairs := []Pair{}
	for i, o := range opts {
		for j := len(opts) - 1; j > i; j-- {
			if shareNone(o, opts[j]) {
				pairs = append(pairs, Pair{o, opts[j]})
			}
		}
	}
	return pairs
}

func shareNone(alpha, beta Category) bool {
	for _, a := range alpha {
		if slices.Contains(beta, a) {
			return false
		}
	}
	return true
}
