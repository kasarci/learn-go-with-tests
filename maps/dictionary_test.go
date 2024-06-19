package maps

import "testing"
const( 
	testWord = "test"
	testDefinition = "this is just a test"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {		
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {

	t.Run("Adding new word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add(testWord, testDefinition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, testWord, testDefinition)
	})

	t.Run("Adding existing word", func(t *testing.T) {
		dictionary := Dictionary{testWord:testDefinition}
		err := dictionary.Add(testWord, "new definition")
		assertError(t, err, ErrWordExists)
	})
}

func TestUpdate(t *testing.T) {
	newDefinition := "new definition"

	t.Run("Existing word", func(t *testing.T) {
		dictionary := Dictionary{testWord:testDefinition}
		err := dictionary.Edit(testWord, newDefinition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, testWord, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Edit(testWord, newDefinition)
		assertError(t, err, ErrWordNotExists)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Existing word", func(t *testing.T) {
		dictionary := Dictionary{testWord:testDefinition}
		dictionary.Delete(testWord)
		_, searchErr := dictionary.Search(testWord)
		assertError(t, searchErr, ErrNotFound)
	})
}

func assertStrings(t testing.TB, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given %q", got, want, "test")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word: ", err)
	}
	assertStrings(t, got, definition)
}