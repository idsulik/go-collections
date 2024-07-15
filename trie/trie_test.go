package trie

import "testing"

func TestTrie_InsertAndSearch(t *testing.T) {
	tr := New()

	// Test inserting and searching for a word
	tr.Insert("hello")
	if !tr.Search("hello") {
		t.Errorf("Expected 'hello' to be found in the Trie")
	}

	// Test searching for a word that does not exist
	if tr.Search("hell") {
		t.Errorf("Expected 'hell' not to be found in the Trie")
	}

	// Test searching for a word that partially matches an existing word
	if tr.Search("helloo") {
		t.Errorf("Expected 'helloo' not to be found in the Trie")
	}
}

func TestTrie_StartsWith(t *testing.T) {
	tr := New()

	// Test prefix search
	tr.Insert("hello")
	if !tr.StartsWith("hel") {
		t.Errorf("Expected Trie to have words starting with 'hel'")
	}

	if !tr.StartsWith("hello") {
		t.Errorf("Expected Trie to have words starting with 'hello'")
	}

	if tr.StartsWith("helloo") {
		t.Errorf("Expected Trie not to have words starting with 'helloo'")
	}

	if tr.StartsWith("hez") {
		t.Errorf("Expected Trie not to have words starting with 'hez'")
	}
}

func TestTrie_InsertMultipleWords(t *testing.T) {
	tr := New()

	// Insert multiple words
	tr.Insert("hello")
	tr.Insert("helium")

	// Test searching for multiple words
	if !tr.Search("hello") {
		t.Errorf("Expected 'hello' to be found in the Trie")
	}

	if !tr.Search("helium") {
		t.Errorf("Expected 'helium' to be found in the Trie")
	}

	if tr.Search("helix") {
		t.Errorf("Expected 'helix' not to be found in the Trie")
	}
}

func TestTrie_EmptyString(t *testing.T) {
	tr := New()

	// Test inserting and searching for an empty string
	tr.Insert("")
	if !tr.Search("") {
		t.Errorf("Expected empty string to be found in the Trie")
	}

	// Test prefix search with empty string
	if !tr.StartsWith("") {
		t.Errorf("Expected Trie to have words starting with empty string")
	}
}
