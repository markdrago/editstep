package main

//import "fmt"

type Word struct {
	name    string
	keys    []string
	pathLen int
}

func New(name string) *Word {
	keys := make([]string, 3)
	keys[0] = name[0:2]
	keys[1] = name[0:1] + name[2:3]
	keys[2] = name[1:3]

	return &Word{name, keys, -1}
}

func (word *Word) Name() string {
	return word.name
}

func (word *Word) PathLen() int {
	return word.pathLen
}

func (word *Word) SetPathLen(pathLen int) {
	word.pathLen = pathLen
}

func (word *Word) Equals(other *Word) bool {
	return word.Name() == other.Name()
}

func (word *Word) Length() int {
	return len(word.Name())
}

//keys are all 3 perms of 2 of first 3 chars
func (word *Word) ClassifyingKeys() []string {
	length := word.Length()
	results := make([]string, 3)
	for i := 0; i < 3; i++ {
		results[i] = string(length) + word.keys[i]
	}
	return results
}

func (word *Word) SearchKeys() []string {
	length := word.Length()
	results := make([]string, 9)
	for i := length - 1; i <= length+1; i++ {
		for j := 0; j < 3; j++ {
			results = append(results, string(i)+word.keys[j])
		}
	}
	return results
}

func (word *Word) AddStep(word2 *Word) {
	//fmt.Printf("wordPathLen: %s:%d, word2PathLen: %s:%d\n", word.name, word.pathLen, word2.name, word2.pathLen)
	if word2.pathLen+1 > word.pathLen {
		word.pathLen = word2.pathLen + 1
	}
}

func (word1 *Word) CanEditStepTo(word2 *Word) bool {
	if word1.Name() > word2.Name() {
		return false
	}

	if word1.Length() == word2.Length() {
		return word1.isEditStepWithSameLength(word2)
	}

	lengthDifference := word1.Length() - word2.Length()
	if lengthDifference > 1 || lengthDifference < -1 {
		return false
	}

	if word1.Length() < word2.Length() {
		return word1.isEditStepShortFirst(word2)
	} else {
		return word2.isEditStepShortFirst(word1)
	}

	return false
}

func (word1 *Word) isEditStepWithSameLength(word2 *Word) bool {
	missCount := 0
	for i := 0; i < word1.Length(); i++ {
		if word1.getNthChar(i) != word2.getNthChar(i) {
			missCount++
			if missCount >= 2 {
				return false
			}
		}
	}

	//if words are identical, that is not an edit step
	if missCount == 0 {
		return false
	}

	return true
}

func (word1 *Word) isEditStepShortFirst(word2 *Word) bool {
	missCount := 0
	for i := 0; i < word1.Length(); i++ {
		if word1.getNthChar(i) != word2.getNthChar(i+missCount) {
			missCount++
			i--
			if missCount >= 2 {
				return false
			}
		}
	}
	return true
}

func (word *Word) getNthChar(n int) string {
	return string(word.Name()[n : n+1])
}
