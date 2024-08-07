package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		if got == nil {
			t.Fatal("expected to get an error.")
		}

		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("adding a word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		description := "this is just a test"

		err := dictionary.Add(word, description)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, description)
	})

	t.Run("word already exists", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"

		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, definition)
		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		newWord := "unknown"
		newDefinition := "new definition"

		err := dictionary.Update(newWord, newDefinition)

		assertError(t, err, ErrWordDoesNotExist)
	})
	t.Run("update existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete a word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		dictionary.Delete(word)
		_, err := dictionary.Search(word)

		assertError(t, err, ErrNotFound)
	})
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, description string) {
	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find added word", err)
	}

	assertString(t, got, description)
}

func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
