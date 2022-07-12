path-add: build
	bin/dfile -v -d dfile.local.db -l dfile.local.log path add ./test/

path-remove: build
	bin/dfile -v -d dfile.local.db -l dfile.local.log path remove "./test/*"

path-sync: build
	bin/dfile -v -d dfile.local.db -l dfile.local.log path sync

path-list: build
	bin/dfile -v -d dfile.local.db -l dfile.local.log path list

path-watch: build
	bin/dfile -v -d dfile.local.db -l dfile.local.log path watch

sum-generate: build
	bin/dfile -v -d dfile.local.db -l dfile.local.log sum generate

sum-duplicated: build
	bin/dfile -v -d dfile.local.db -l dfile.local.log sum duplicated

sum-duplicated-json: build
	bin/dfile -v -d dfile.local.db -l dfile.local.log sum duplicated -o json

sum-duplicated-json-i: build
	bin/dfile -v -d dfile.local.db -l dfile.local.log sum duplicated -o json-i

db-rewrite: build
	bin/dfile -v -d dfile.local.db -l dfile.local.log db rewrite

db-rewrite-indented: build
	bin/dfile -v -d dfile.local.db -l dfile.local.log db rewrite -i

build:
	go build -o bin/dfile main.go
