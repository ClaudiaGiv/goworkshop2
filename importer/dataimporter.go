package importer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"goworkshop/model"
)

func ImportAuthors() []model.AuthorDto {

	var authors []model.AuthorDto
	if authorsFile, err := ioutil.ReadFile("json/authors.json");  err != nil{
		fmt.Println("Unable to open file")
		panic(err)
	} else {
		if err = json.Unmarshal(authorsFile,&authors); err!= nil {
			fmt.Println("Unable to de-serialize the authors")
			panic(err)
		}
	}
	return authors
}



func ImportBooks() []model.BookDto {

	var books []model.BookDto
	if booksFile, err := ioutil.ReadFile("json/books.json"); err != nil{
		fmt.Println("Unable to open file")
		panic(err)
	} else {
		if err = json.Unmarshal(booksFile,&books); err!= nil {
			fmt.Println("Unable to de-serialize the books")
			panic(err)
		}
	}
	return books

}
