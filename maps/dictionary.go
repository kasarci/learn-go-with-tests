package maps

import "errors"

type Dictionary map[string]string

var (
	ErrNotFound = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("cannot add word because it already exists")
	ErrWordNotExists = errors.New("cannot update the word because it does not exist")
)

func (d Dictionary) Search(word string) (string, error) {
	result, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return result, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Edit(word, newDefinition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordNotExists
	case nil:
		d[word] = newDefinition
	default:
		return err
	}
	return nil
}
