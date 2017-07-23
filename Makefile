all: library

clean:
	rm -f bin/library

library:
	go build -o bin/library
