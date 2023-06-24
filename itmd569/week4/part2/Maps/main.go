package main

import "fmt"

func main() {
	m0 := make(map[string]int)
	fmt.Println("Empty map m0: ", m0)

	// adding values to the map
	m0["key1"] = 1
	m0["key2"] = 2
	m0["key3"] = 3
	m0["key4"] = 469
	m0["key5"] = 569
	m0["password"] = 12345

	fmt.Println("m0: ", m0)
	fmt.Println("Length of m0: ", len(m0))

	// deleting values from map
	delete(m0, "key4")
	fmt.Println("After deleting key4: ", m0)

	// retrieving values from the lab
	password := m0["password"]
	fmt.Println("Password grabbed from map: ", password)

	// second return value from map, if item exists

	key := "key5"
	item, exists := m0[key]

	if exists {
		fmt.Println(exists)
		fmt.Println("item exists: ", key, item)
	} else {
		fmt.Println(exists)
		fmt.Println("item does not exist")
	}

}
