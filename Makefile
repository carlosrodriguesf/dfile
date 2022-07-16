path-add: build
	bin/dfile -v -d dfile.db -l dfile.log path add ./test/

path-remove: build
	bin/dfile -v -d dfile.db -l dfile.log path remove "./test/"

path-sync: build
	bin/dfile -v -d dfile.db -l dfile.log path sync

path-list: build
	bin/dfile -v -d dfile.db -l dfile.log path list

sum-generate: build
	bin/dfile -v -d dfile.db -l dfile.log sum generate

sum-duplicated: build
	bin/dfile -v -d dfile.db -l dfile.log sum duplicated

sum-duplicated-json: build
	bin/dfile -v -d dfile.db -l dfile.log sum duplicated -o json

sum-duplicated-json-i: build
	bin/dfile -v -d dfile.db -l dfile.log sum duplicated -o json-i

watch: build
	bin/dfile -v -d dfile.db -l dfile.log watch

build:
	go build -o bin/dfile main.go
