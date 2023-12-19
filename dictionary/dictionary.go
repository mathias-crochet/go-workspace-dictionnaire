package dictionary

import (
    "fmt"
)

type dictionary map[string]string

func New() dictionary {
	return make(dictionary)
}

func (d dictionary) Get(reference string) string {
	return d[reference]
}

func (d dictionary) Add(reference string, value string) {
	d[reference] = value
}

func (d dictionary) List() {
	fmt.Println("Dictionnaire :")
	for key, value := range d {
		fmt.Printf("%s: %s\n", key, value)
	}
}

func (d dictionary) Remove(reference string) {
	delete(d, reference)
}
