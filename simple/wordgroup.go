package main

import word "./word"
//import "fmt"

type WordGroup struct {
	words []*word.Word
}

func New() *WordGroup {
	return &WordGroup{make([]*word.Word, 0)}
}

func (wordGroup *WordGroup) AddWord(theword *word.Word) {
	wordGroup.words = append(wordGroup.words, theword)
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
