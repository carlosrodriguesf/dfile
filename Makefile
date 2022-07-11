scan:
	go run main.go path scan ./test/

remove:
	go run main.go path remove "./test/*"

generate-sum:
	go run main.go sum generate

build:
	go build -o bin/dfile ./main.go