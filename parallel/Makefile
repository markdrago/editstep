all: editstep

editstep: word.6 wordgroup.6 wordindex.6 queueworker.6 wordqueue.6 editstep.6
	6l -o editstep editstep.6

editstep.6: editstep.go
	gofmt -w editstep.go
	6g -o editstep.6 editstep.go

word.6: word.go
	gofmt -w word.go
	6g -o word.6 word.go

wordindex.6: wordindex.go
	gofmt -w wordindex.go
	6g -o wordindex.6 wordindex.go

wordgroup.6: wordgroup.go
	gofmt -w wordgroup.go
	6g -o wordgroup.6 wordgroup.go

queueworker.6: queueworker.go
	gofmt -w queueworker.go
	6g -o queueworker.6 queueworker.go

wordqueue.6: wordqueue.go
	gofmt -w wordqueue.go
	6g -o wordqueue.6 wordqueue.go

clean:
	-rm -f *.6
	-rm -f editstep

test: editstep
	./editstep < wordlists/small

bigtest: editstep
	./editstep < wordlists/big
