package main

import (
	"fmt"
	"slices"
)

func main() {
	cart := []string{"apple", "orange", "banana"}
	fmt.Println("len:", len(cart))
	fmt.Println("cart[1]:", cart[1])

	// Indices
	for i := range cart {
		fmt.Println(i)
	}
	// index + values
	for i, c := range cart {
		fmt.Println(i, c)
	}
	// values
	for _, c := range cart {
		fmt.Println(c)
	}
	cart = append(cart, "milk")
	fmt.Println(cart)

	// slicing operator
	fruit := cart[:3]
	fmt.Println("fruit:", fruit)
	fruit = append(fruit, "lemon")
	fmt.Println("fruit:", fruit)
	fmt.Println("cart:", cart)

	var s []int
	for i := range 10_000 {
		s = appendInt(s, i)
	}
	fmt.Println(s[:20])

	// Exercise: concat. without using a "for" loop
	out := concat([]string{"A", "B"}, []string{"C"})
	fmt.Println("concat:", out) // [A B C]

	// Exercise: find the median of the array
	values := []float64{3, 1, 2}
	fmt.Println(median(values)) // 2
	values = []float64{3, 1, 2, 4}
	fmt.Println(median(values)) // 2.5
	//fmt.Println(median([]float64{})) //discarded for this exercise

	players := []Player{
		{"Rick", 10_000},
		{"Morty", 11},
	}

	// Add a bonus

	// Value semantics "for" loop
	for _, p := range players {
		// Go is working by value
		p.Score += 100
	}

	fmt.Println(players)

	// "Pointer" semantics "for" loop
	for i := range players {
		players[i].Score += 100
	}
	fmt.Println(players)

}

type Player struct {
	Name  string
	Score int
}

/*
	func concat(s1 []string, s2 []string) {
		return nil // FIXME: Your code goes here
	}
*/
func concat(s1 []string, s2 []string) []string {
	out := make([]string, len(s1)+len(s2))
	copy(out, s1)
	copy(out[len(s1):], s2)

	return out
}

/*
	median:

- sort values
- if odd number of values: return middle
- return average of middles
*/

func median(values []float64) float64 {

	s := slices.Clone(values)
	slices.Sort(s)

	n := len(s)
	var result float64
	result = 0

	if n%2 == 0 {
		result = (s[n/2-1] + s[n/2]) / 2
	} else {
		result = s[n/2]
	}

	return result
}

func appendInt(s []int, v int) []int {
	i := len(s)
	if len(s) == cap(s) {
		// nore more space in underlying array
		// need to reallocate and copy
		size := 2 * (len(s) + 1)
		fmt.Println(cap(s), "->", size)
		ns := make([]int, size)
		copy(ns, s)
		s = ns[:len(s)]
	}
	s = s[:len(s)+1]
	s[i] = v
	return s
}
