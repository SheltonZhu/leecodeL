package _08_implement_trie

import "testing"

func Test(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
	if !trie.Search("apple") {
		t.Errorf("not find %s", "apple")
	} //true
	if trie.Search("app") {
		t.Errorf("should not find %s", "app")
	} //false
	if !trie.StartsWith("app") {
		t.Errorf("should not start with %s", "app")
	} //true
	trie.Insert("app")
	if !trie.Search("app") {
		t.Errorf("not find %s", "app")
	} //true
}
