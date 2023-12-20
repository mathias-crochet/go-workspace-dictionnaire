package main

import (
	"dictionnaire/dictionary"
	"fmt"
)

func main() {

	file := dictionary.NewFile("dictionary.json")

	langageGo := dictionary.NewDefinition("go", "langage de programmation compilé créé par Google")
	estiam := dictionary.NewDefinition("estiam", "École supérieure des technologies de l’information appliquées aux métiers")

	addChannel := make(chan string)
	removeChannel := make(chan string)

	go func()  {
		dictionary.Add(langageGo, file)
		addChannel <- "la définition a été ajoutée"
	}()

    fmt.Println(<-addChannel)

	go func()  {
		dictionary.Add(estiam, file)
		addChannel <- "la définition a été ajoutée"
	}()

    fmt.Println(<-addChannel)

	
	go func()  {
		dictionary.Remove("estiam", file)
		removeChannel <- "la définition a été supprimé"
	}()

    fmt.Println(<-removeChannel)

		
	dictionary.Get("go", file)
	dictionary.List(file)
	
	dictionary.List(file)
	

}
