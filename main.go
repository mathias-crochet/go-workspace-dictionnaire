package main

import (
	"dictionnaire/dictionary"
    "fmt"
)

func main() {

	d := dictionary.New()

	d.Add("etudiant", "mathias")
	d.Add("ecole", "estiam")
	d.Add("date", "19/12/2023")

	etu := d.Get("etudiant")
	fmt.Println(etu)

	d.List()

	d.Remove("ecole")
	fmt.Println("dictionnaire sans 'ecole':", d)

}