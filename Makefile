build:
	go build -o bin/dfile main.go

test-path-add: build
	bin/dfile path add ./test/
