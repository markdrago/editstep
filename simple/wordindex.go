package main

import word "./word"
import wordGroup "./wordgroup"
//import "fmt"

type WordIndex struct {
	wordGroups map[string]*wordGroup.WordGroup
}

func New() *WordIndex {
	return &WordIndex{make(map[string]*wordGroup.WordGroup)}
}

//return length of longest path away from this word
func (wordIndex *WordIndex) AddWord(name string) int {
	theword := word.New(name)

	wordIndex.updateEditSteps(theword, theword.SearchKeys())

	//add this word to all req'd groups so later words can find it
	for _, key := range theword.ClassifyingKeys() {
		wordIndex.addWordToIndexWithKey(theword, key)
	}

	return theword.PathLen()
}

func (wordIndex *WordIndex) addWordToIndexWithKey(theword *word.Word, key string) {
	if !wordIndex.hasGroupForKey(key) {
		wordIndex.wordGroups[key] = wordGroup.New()
	}
	wordIndex.wordGroups[key].AddWord(theword)
}

func (wordIndex *WordIndex) updateEditSteps(theword *word.Word, keys []string) {
	//fmt.Printf("Checking for: '%s'\n", theword.Name())

	//check groups with this words keys
	for _, key := range keys {
		if !wordIndex.hasGroupForKey(key) {
			continue
		}
		wordIndex.wordGroups[key].UpdateEditSteps(theword)
	}

	//fmt.Printf("    PathLen: %d\n\n", theword.PathLen())
}

func (wordIndex *WordIndex) hasGroupForKey(key string) bool {
	_, present := wordIndex.wordGroups[key]
	return present
}
