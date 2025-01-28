package main

import (
	"fmt"
	"maps"
)

func main() {
	// creating map
	m := make(map[string]string)

	// set a value
	m["language"] = "go lang"
	m["designation"] = "backend"
	// get a value
	fmt.Println(m["language"])

	// IMP: if key does not exists in map it returns the zerod value.
	// e.g.
	fmt.Println(m["course"])

	// length of map
	fmt.Println(len(m))

	// delete item from map
	delete(m, "language")
	fmt.Println(m)
	// empty map
	clear(m)
	fmt.Println(m)

	// creating map without make function
	m1 := map[string]string{"book": "go lang professional"}
	fmt.Println(m1)
	val, ok := m1["books"]
	if ok {
		fmt.Println("all ok", val)
	} else {
		fmt.Println("not ok")
	}

	// check map is equal
	m2 := map[string]string{"book": "go lang professional"}
	fmt.Println(maps.Equal(m1, m2))
}
