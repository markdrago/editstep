package main

import word "./word"
import queueworker "./queueworker"
import wordindex "./wordindex"
import "fmt"

type WordQueue struct {
	maxPathLen    int
	active        map[*word.Word][]*word.Word
	workerChan    chan *word.Word
	inbound       chan *word.Word
	limiterChan   chan int
	completed     bool
	completedChan chan int
}

func New(inbound chan *word.Word, limiterChan chan int, completedChan chan int, workQueueBufferLen int) *WordQueue {
	index := wordindex.New()

	//spawn workers, give them channel handles
	numWorkers := 16
	workerChan := make(chan *word.Word, workQueueBufferLen)
	for i := 0; i < numWorkers; i++ {
		worker := queueworker.New(workerChan, inbound, index)
		go worker.Work()
	}

	queue := new(WordQueue)
	queue.maxPathLen = 0
	queue.active = make(map[*word.Word][]*word.Word)
	queue.workerChan = workerChan
	queue.inbound = inbound
	queue.limiterChan = limiterChan
	queue.completed = false
	queue.completedChan = completedChan
	return queue
}

func (queue *WordQueue) WaitForWords() {
	for true {
		theword := <-queue.inbound
		if theword == nil {
			//fmt.Printf("got nil in queue\n")
			queue.completed = true
			if len(queue.active) == 0 {
				queue.completeWork()
				return
			}
			continue
		}

		//words with pathLen >= 0 are coming back from a worker
		if theword.PathLen() >= 0 {
			//fmt.Printf("q got worker response: %s\n", theword.Name())
			<-queue.limiterChan
			if theword.PathLen() > queue.maxPathLen {
				queue.maxPathLen = theword.PathLen()
			}

			waiters := queue.active[theword]
			queue.active[theword] = nil, false //remove item from map
			queue.wakeUpWaitingWords(waiters)

			//fmt.Printf("activeLen: %d\n", len(queue.active))
			//for active, _ := range queue.active {
			//fmt.Printf("active: %s\n", active)
			//}
			if queue.completed == true && len(queue.active) == 0 {
				queue.completeWork()
				return
			}
		} else {
			//words with pathLen < 0 have not been processed yet
			//fmt.Printf("q found word: %s\n", theword.Name())
			queue.AddWord(theword)
		}
	}
}

func (queue *WordQueue) completeWork() {
	fmt.Printf("%d\n", queue.maxPathLen+1)
	queue.completedChan <- 1
}

func (queue *WordQueue) AddWord(theword *word.Word) {
	//will get woken up when another word completes
	mustWait := queue.wordMustWaitForActiveWord(theword)

	//add this word to the list of active words
	_, present := queue.active[theword]
	if !present {
		queue.active[theword] = make([]*word.Word, 0)
	}

	//wait for later if this word must wait
	if mustWait {
		return
	}

	queue.processWord(theword)
}

func (queue *WordQueue) processWord(theword *word.Word) {
	queue.workerChan <- theword
}

func (queue *WordQueue) wordMustWaitForActiveWord(theword *word.Word) bool {
	mustWait := false
	for activeWord, _ := range queue.active {
		if theword.CanEditStepTo(activeWord) {
			mustWait = true
			queue.addWaiterIfNotAlreadyWaiting(theword, activeWord)
		}
	}
	return mustWait
}

func (queue *WordQueue) addWaiterIfNotAlreadyWaiting(waiter *word.Word, activeWord *word.Word) {
	//fmt.Printf("%s must wait for %s to complete\n", waiter.Name(), activeWord.Name())
	for _, alreadyWaiting := range queue.active[activeWord] {
		if waiter == alreadyWaiting {
			return
		}
	}
	queue.active[activeWord] = append(queue.active[activeWord], waiter)
}

func (queue *WordQueue) wakeUpWaitingWords(waiters []*word.Word) {
	for _, waiter := range waiters {
		//fmt.Printf("Waking up %s\n", waiter.Name())
		queue.AddWord(waiter)
	}
}
