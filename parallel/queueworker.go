package main

import word "./word"
import wordindex "./wordindex"
//import "fmt"

type QueueWorker struct {
	wordIndex *wordindex.WordIndex
	inbound   chan *word.Word
	outbound  chan *word.Word
}

func New(inbound chan *word.Word, outbound chan *word.Word, index *wordindex.WordIndex) *QueueWorker {
	worker := new(QueueWorker)
	worker.wordIndex = index
	worker.inbound = inbound
	worker.outbound = outbound
	return worker
}

func (worker *QueueWorker) Work() {
	for true {
		theword := <-worker.inbound
		if theword == nil {
			return
		}

		//fmt.Printf("worker found word: %s\n", theword.Name())

		worker.wordIndex.AddWord(theword)

		//fmt.Printf("worker responding: %s -> %d\n", theword.Name(), theword.PathLen())
		worker.outbound <- theword
	}
}
