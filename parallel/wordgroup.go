package main

import word "./word"
import "sync"
//import "fmt"

type WordGroup struct {
	words      []*word.Word
	sliceMutex sync.Mutex
}

func New() *WordGroup {
	wordGroup := new(WordGroup)
	wordGroup.words = make([]*word.Word, 0)
	return wordGroup
}

func (wordGroup *WordGroup) AddWord(theword *word.Word) {
	wordGroup.sliceMutex.Lock()
	wordGroup.words = append(wordGroup.words, theword)
	wordGroup.sliceMutex.Unlock()
}

func (wordGroup *WordGroup) Words() []*word.Word {
	return wordGroup.words
}

func (wordGroup *WordGroup) UpdateEditSteps(src *word.Word) {
	for _, target := range wordGroup.words {
		if src.Equals(target) {
			//do not add word to its own edit step list
			continue
		}

		if src.CanEditStepTo(target) {
			//fmt.Printf("  Found: '%s'\n", target.Name())
			src.AddStep(target)
		}
	}
}
