package main

import (
	"testing"
)

func TestDictionary(t *testing.T) {
	dict := Dictionary{}
	dict.AddMessage("en", "English")

	if val, ok := dict["en"]; ok {
		if val != "English" {
			t.Fatal("Wrong msgstr")
		}
	} else {
		t.Fatal("msgid not found")
	}
	dict.RemoveMessage("en")
}
