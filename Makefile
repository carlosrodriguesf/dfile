path-add: build
	bin/dfile -v -d dfile.local.db -l dfile.local.log path add ./test/

path-remove: build
	bin/dfile -v -d dfile.local.db -l dfile.local.log path remove "./test/*"

path-sync: build
	bin/dfile -v -d dfile.local.db -l dfile.local.log path sync

path-list: build
	bin/dfile -v -d dfile.local.db -l dfile.local.log path list

sum-generate: build
	bin/dfile -v -d dfile.local.db -l dfile.local.log sum generate

build:
	go build -o bin/dfile main.go
