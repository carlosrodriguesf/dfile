path-add: build
	bin/dfile -d dfile.local.db -l dfile.local.log path add ./test/

path-remove: build
	bin/dfile -d dfile.local.db -l dfile.local.log path remove "./test/*"

path-sync: build
	bin/dfile -d dfile.local.db -l dfile.local.log path sync

sum-generate: build
	bin/dfile -d dfile.local.db -l dfile.local.log sum generate

build:
	go build -o bin/dfile main.go
