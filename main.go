package main

import (
    "fmt"
)

func main() {

    m := make(map[string]int)

    m["a"] = 2
    m["b"] = 4
	m["c"] = 6

    fmt.Println("map:", m)

	// get
    a := get(m, "a")
	fmt.Println("valeur de 'a' :", a)

	//list
	list(m)

	//remove
	remove(m, "c")
    fmt.Println("map sans 'c' :", m)

}

func get(dictionnaire map[string]int, reference string) int {
	return dictionnaire[reference]
}

func list(dictionnaire map[string]int) {
	fmt.Println("Dictionnaire :")
	for key, value := range dictionnaire {
		fmt.Printf("%s: %d\n", key, value)
	}
}

func remove(dictionnaire map[string]int, reference string) {
	delete(dictionnaire, reference)
}