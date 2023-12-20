package dictionary

import (
	"encoding/json"
	"fmt"
	"os"
	"log"
)

type dictionary map[string]string

type Definition struct {
	Mot        string `json:"mot"`
	Definition string `json:"definition"`
}

func NewDictionary() dictionary {
	return make(dictionary)
}

func NewDefinition(mot string, definition string) Definition {
	return Definition{mot, definition}
}

func NewFile(fileName string) (*os.File) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal("erreur lors de la création du fichier :", err) 
	}
	return file
}

func GetFileData(file *os.File)  (dictionary) {

	var dictionnaire dictionary

	// set cursor at line 0
	file.Seek(0, 0)

	decodeur := json.NewDecoder(file)
	err := decodeur.Decode(&dictionnaire)
	if err != nil {
		// check if end of file
		if err.Error() == "EOF" {
			return NewDictionary()
		}
		log.Fatal("erreur :", err) 
	}

	return dictionnaire
}

func Add(definition Definition, file *os.File) {

	dictionnaire := GetFileData(file)

	dictionnaire[definition.Mot] = definition.Definition

	file.Seek(0, 0)
	file.Truncate(0) 

	json.NewEncoder(file).Encode(dictionnaire)

}

func Get(mot string, file *os.File) {

	dictionnaire := GetFileData(file)

	definition, ok := dictionnaire[mot]
	if !ok {
		log.Fatal("Aucune définition trouvée pour le mot ", mot)
	}

	fmt.Printf("Définition pour le mot %s: %s\n", mot, definition)
}

func List(file *os.File) {

	dictionnaire := GetFileData(file)

	fmt.Println("Dictionnaire :")
	for mot, definition := range dictionnaire {
		fmt.Printf("%s: %s\n", mot, definition)
	}

}

func Remove(mot string, file *os.File) {

	dictionnaire := GetFileData(file)

	delete(dictionnaire, mot)

	// rewrite of file
	file.Seek(0, 0)
	file.Truncate(0)

	json.NewEncoder(file).Encode(dictionnaire)

}
