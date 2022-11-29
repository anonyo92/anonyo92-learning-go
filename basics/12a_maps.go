package basics

import "fmt"

func runMaps() {
	m := make(map[string]int)

	// Set/Get
	m["k1"] = 5
	m["k2"] = 7
	fmt.Println("map:", m)
	fmt.Println("value with key 'k1' :", m["k2"])

	fmt.Println("len:", len(m))

	// Remove specific element from map
	delete(m, "k2")
	fmt.Println("map:", m)

	/* The optional second return value when getting a value from a map
	 * indicates if the key was present in the map. This can be used to
	 * disambiguate between missing keys and keys with zero values like 0 or "". */
	_, prs := m["k2"] // 'prs' will be false here, as we have deleted key 'k2'
	/* Here we didn’t need the value itself, so we ignored it with the blank identifier _ */
	fmt.Println("prs:", prs)

	// Dynamic initialization
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
