package main

import (
	"os"
	"bufio"
	"runtime"
	word "./word"
	wordqueue "./wordqueue"
)

func main() {
	runtime.GOMAXPROCS(3)

	//create work queue
	queueChan := make(chan *word.Word)
	workQueueBufferLen := 16
	limiterChan := make(chan int, workQueueBufferLen)
	completedChan := make(chan int)
	queue := wordqueue.New(queueChan, limiterChan, completedChan, workQueueBufferLen)
	go queue.WaitForWords()

	//loop over list of words in reverse
	lines := getLines(getReader())
	for i := len(lines) - 1; i >= 0; i-- {
		theword := word.New(string(lines[i]))
		limiterChan <- 1
		queueChan <- theword
	}
	queueChan <- nil

	<-completedChan
}

func getReader() *bufio.Reader {
	return bufio.NewReader(os.Stdin)
}

func getLines(reader *bufio.Reader) []string {
	lines := make([]string, 0)

	var line []byte
	var err os.Error
	for err == nil {
		line, _, err = reader.ReadLine()
		if len(line) != 0 {
			lines = append(lines, string(line))
		}
	}

	return lines
}
