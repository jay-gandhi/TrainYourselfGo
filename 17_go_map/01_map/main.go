package main

import "fmt"

func main() {
	// A map maps keys to values. unordered
	// Maps built on hashtable
	// value of unintialized map is nil
	// Map is reference type
	m := make(map[string]int)
	// Golang SPec
	// A map is an unordered group of elements of one type,
	// called the element type, indexed by a set of
	// unique keys of another type, called the key type.

	m["k1"] = 21
	m["k2"] = 22
	fmt.Println("map", m)
	fmt.Println("length", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	// The optional second return value when getting a
	// value from a map indicates if the key was present
	// in the map. This can be used to disambiguate
	// between missing keys and keys with zero values
	// like `0` or `""`. Here we didn't need the value
	// itself, so we ignored it with the _blank identifier_
	// `_`.
	v, prs := m["k2"]
	v1, prs1 := m["k1"]
	fmt.Println("k2:", prs, v)   // It will give zero value
	fmt.Println("k2:", prs1, v1) // It will give current value

	// you can declared and initialize a new map in
	// the same line with syntax
	n := map[string]int{"foo": 123, "bar": 2}
	fmt.Println("map: ", n)
}
