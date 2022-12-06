build:
	go build -o bin/dfile main.go

test-path-add: build
	bin/dfile path add ./test/

test-path-sync: build
	bin/dfile path sync

test-generate-sum: build
	bin/dfile file generate-sum