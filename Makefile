build:
	go build -o bin/ft
run:
	./bin/ft input.txt
build-run:
	go build -o bin/ft
	./bin/ft input.txt
test:
	go test -v ./tree
	go test -v ./relation
	go test -v ./person
	go test -v ./handler
test-coverage:
	go test -cover ./...