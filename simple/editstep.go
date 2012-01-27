package main

import (
	"fmt"
	"os"
	"bufio"
	wordindex "./wordindex"
)

func main() {
	lines := getLines(getReader())
	index := wordindex.New()

	//loop over list of words in reverse
	maxPathLen := 0
	for i := len(lines) - 1; i >= 0; i-- {
		pathLen := index.AddWord(lines[i])
		if pathLen > maxPathLen {
			maxPathLen = pathLen
		}
	}

	fmt.Printf("%d\n", maxPathLen+1)
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
