scan: build
	bin/dfile -d dfile.local.db -l dfile.local.log path scan ./test/

remove: build
	bin/dfile -d dfile.local.db -l dfile.local.log path remove "./test/*"

generate-sum: build
	bin/dfile -d dfile.local.db -l dfile.local.log sum generate

build:
	go build -o bin/dfile ./main.go