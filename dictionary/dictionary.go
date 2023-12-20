package dictionary

import (
	"encoding/json"
	"fmt"
	"os"
	"log"
	"net/http"
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

func NewFile(fileName string) {
	_, err := os.Create(fileName)
	if err != nil {
		log.Fatal("erreur lors de la création du fichier :", err) 
	}
}

func GetFile(fileName string) *os.File {
	file, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal("erreur lors de l'ouverture du fichier :", err)
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

// url: http://localhost:8080/add?mot=go&definition=langage%20de%20programmation%20compil%C3%A9%20cr%C3%A9%C3%A9%20par%20Google
func Add(w http.ResponseWriter, req *http.Request) {

	file := GetFile("dictionary.json")

	dictionnaire := GetFileData(file)

	mot := req.URL.Query().Get("mot")
	definition := req.URL.Query().Get("definition")

	if mot != "" && definition != "" {
		dictionnaire[mot] = definition
	}

	file.Seek(0, 0)
	file.Truncate(0) 

	json.NewEncoder(file).Encode(dictionnaire)

}

// url: http://localhost:8080/get?mot=go
func Get(w http.ResponseWriter, req *http.Request) {

	file := GetFile("dictionary.json")

	dictionnaire := GetFileData(file)

	mot := req.URL.Query().Get("mot")

	definition, ok := dictionnaire[mot]
	if !ok {
		log.Fatal("Aucune définition trouvée pour le mot ", mot)
	}

	fmt.Printf("Définition pour le mot %s: %s\n", mot, definition)
}

// url: http://localhost:8080/list
func List(w http.ResponseWriter, req *http.Request) {

	file := GetFile("dictionary.json")

	dictionnaire := GetFileData(file)

	fmt.Println("Dictionnaire :")
	for mot, definition := range dictionnaire {
		fmt.Printf("%s: %s\n", mot, definition)
	}

}

// url: http://localhost:8080/remove?mot=go
func Remove(w http.ResponseWriter, req *http.Request) {

	file := GetFile("dictionary.json")

	dictionnaire := GetFileData(file)

	mot := req.URL.Query().Get("mot")

	delete(dictionnaire, mot)

	// rewrite of file
	file.Seek(0, 0)
	file.Truncate(0)

	json.NewEncoder(file).Encode(dictionnaire)
}
