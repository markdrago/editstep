package main

import word "./word"
import wordGroup "./wordgroup"
import "sync"
//import "fmt"

type WordIndex struct {
	wordGroups map[string]*wordGroup.WordGroup
	mapMutex   sync.Mutex
}

func New() *WordIndex {
	index := new(WordIndex)
	index.wordGroups = make(map[string]*wordGroup.WordGroup)
	return index
}

//return length of longest path away from this word
func (wordIndex *WordIndex) AddWord(theword *word.Word) {
	wordIndex.updateEditSteps(theword, theword.SearchKeys())

	//add this word to all req'd groups so later words can find it
	for _, key := range theword.ClassifyingKeys() {
		wordIndex.addWordToIndexWithKey(theword, key)
	}
}

func (wordIndex *WordIndex) addWordToIndexWithKey(theword *word.Word, key string) {
	if !wordIndex.hasGroupForKey(key) {
		wordIndex.mapMutex.Lock()
		wordIndex.wordGroups[key] = wordGroup.New()
		wordIndex.mapMutex.Unlock()
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

	//if we did not find anything to editstep to, record a 0
	if theword.PathLen() < 0 {
		theword.SetPathLen(0)
	}

	//fmt.Printf("    PathLen: %d\n\n", theword.PathLen())
}

func (wordIndex *WordIndex) hasGroupForKey(key string) bool {
	wordIndex.mapMutex.Lock()
	defer wordIndex.mapMutex.Unlock()
	_, present := wordIndex.wordGroups[key]
	return present
}
