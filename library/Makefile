all: library

clean:
	rm -f bin/library

test: test-domain

test-domain:
	src/domain

library:
	go build -o bin/library

dist:
	CGO_ENABLED=0 go build -a -installsuffix cgo -v -o bin/library
