package main

import (
	"dictionnaire/dictionary"
)

func main() {

	file := dictionary.NewFile("dictionary.json")

	langageGo := dictionary.NewDefinition("go", "langage de programmation compilé créé par Google")
	estiam := dictionary.NewDefinition("estiam", "École supérieure des technologies de l’information appliquées aux métiers")

	dictionary.Add(langageGo, file)
	dictionary.Add(estiam, file)
	
	dictionary.Get("go", file)

	dictionary.List(file)

	dictionary.Remove("estiam", file)
	dictionary.List(file)
	

}
