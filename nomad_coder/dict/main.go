package main

import (
	"dict/mydict"
	"fmt"
)

func main() {

	dictionary := mydict.Dictionary{"first": "First word"}

	dictionary["hello"] = "hello"

	fmt.Println(dictionary)
	fmt.Println(dictionary["first"])
	definition, err := dictionary.Search("first")
	// definition, err := dictionary.Search("second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

	word := "plus"
	difinition := "add"
	err = dictionary.Add(word, difinition)
	if err != nil {
		fmt.Println(err)
	}

	definition2, err := dictionary.Search(word)
	fmt.Println(definition2)

	err = dictionary.Add(word, difinition)
	if err != nil {
		fmt.Println(err)
	}
	definition3, err := dictionary.Search(word)
	fmt.Println(definition3)

	err = dictionary.Update(word, "더하기")
	if err != nil {
		fmt.Println(err)
	}
	finalDef, _ := dictionary.Search(word)
	fmt.Println(finalDef)

	fmt.Println("삭제===")
	dictionary.Delete(word)

	delDict, err := dictionary.Search(word)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(delDict)
	}
}
