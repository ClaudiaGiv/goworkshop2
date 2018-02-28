package main

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
	"goworkshop/importer"
)

type Animal struct{
	NoLegS int `json:"-"`
	NoLegs int `json:"noLegs"`
	Name string `json:"name"`
}

type Talker interface{
	CanTalk() bool
}

func (a Animal) CanTalk() bool{
	return false
}

func (a Animal) String() string{
	return fmt.Sprintf("Animal{NoLegs=%d, Name=%s}",a.NoLegs, a.Name)
}

func main() {
	var creature Talker
	creature = Animal{NoLegs:3, Name:"Pig",}
	fmt.Println(creature.CanTalk())

	fileContent, err := ioutil.ReadFile("json/animals.json")
	if err != nil{
		fmt.Println("Unable to open file")
		panic(err)
	}
	fmt.Println(string(fileContent))

	var animals []Animal
	err = json.Unmarshal(fileContent,&animals)
	if err!= nil {
		fmt.Println("Unable to de-serialize the aimals")
		panic(err)
	}

	//check the de-serialized animals
	fmt.Println("The animals are:")
	fmt.Println(animals)

	if serializedAnimals, err := json.Marshal(animals); err != nil {
		fmt.Println("Unable to serialize the aimals")
		panic(err)
	} else {
		fmt.Println(string(serializedAnimals))
	}


	authors := importer.ImportAuthors()
	fmt.Println(authors)
	books := importer.ImportBooks()
	fmt.Println(books)
}
