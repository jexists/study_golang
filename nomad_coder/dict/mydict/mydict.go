package mydict

import (
	"errors"
)

// Dictionary type
type Dictionary map[string]string

type Money int

// var errNotFound = errors.New("없음")
// var errCantUpdate = errors.New("없어서 업데이트불가")
// var errWordExists = errors.New("존재")
// 깔끔해지게 만들기
var (
	errNotFound   = errors.New("없음")
	errCantUpdate = errors.New("없어서 업데이트불가")
	errWordExists = errors.New("존재")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	// if err == errNotFound {
	// 	d[word] = def
	// } else if err == nil {
	// 	return errWordExists
	// }
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}

// Update a word
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = def
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
