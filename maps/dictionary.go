package maps

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	result := d[word]
	if result == "" {
		return "", errors.New("could not find the word you were looking for")
		
	}
	return result, nil
}