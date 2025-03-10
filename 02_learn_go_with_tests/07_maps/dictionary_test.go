package dictionary

import "testing"

func TestDictionary(t *testing.T) {
	dictionary := map[string]string{
		"verb":    "A verb is an action word",
		"pronoun": "A pronoun is a word that takes the place of a noun in a sentence.",
	}

	got := dictionary.Search("verb")
	want := "A verb is an action word"

	asserDictionary(t, got, want)
}

func asserDictionary(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got: %s want: %s", got, want)
	}
}
