CFLAGS=-std=c11 -g -static

9cc: main.go

test: 9cc
	go build main.go
	./test.sh
	rm -f 9cc *.o *~ tmp* main

debug: 9cc
	go build main.go
	bash -x test.sh
	rm -f 9cc *.o *~ tmp* main

clean:
	rm -f 9cc *.o *~ tmp* main

.PHONY: test clean