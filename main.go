package main

import (
	"fmt"
)

/*
   For each fruit in the list find two other which summed price is equal to this fruit price.
*/

type product struct {
	name  string
	price int
}

var products = []product{
	{name: "kiwi", price: 2},
	{name: "grapefruit", price: 4},
	{name: "pomegranate", price: 5},
	{name: "mango", price: 6},
	{name: "pumpkin", price: 7},
	{name: "pineapple", price: 11},
	{name: "watermelon", price: 15},
}

// example result
// mango:[[kiwi grapefruit]]
// pumpkin:[[kiwi pomegranate]]
// pineapple:[[grapefruit pumpkin] [pomegranate mango]]
// watermelon:[[grapefruit pineapple]]]

func main() {
	out := map[string][][]string{}
	for i := range products {
		ourName := products[i].name
		ourPrice := products[i].price

		//var firstProductIndex = -1
		for j := range products {
			if i == j {
				continue
			}

			if products[j].price > ourPrice {
				continue
			}

			for k := j + 1; k < len(products); k++ {
				if i == k || j == k {
					continue
				}

				if products[j].price+products[k].price == ourPrice {
					pairs := out[ourName]
					out[ourName] = append(pairs, []string{
						products[j].name,
						products[k].name,
					})
					//firstProductIndex = j
				}
			}
		}
	}

	for _, p := range products {
		combinations, found := out[p.name]
		if !found {
			continue
		}
		fmt.Printf("%s: %v\n", p.name, combinations)
	}
}
