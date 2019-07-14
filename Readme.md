# Familytree Problem

The code challenge was taken up with the following assumptions.

## Assumptions
    - The initial input data for the tree ( diagram - https://www.geektrust.in/api/pdf/open/PS1Evaluate)
    - Every person is expected to have a unique name.
    - The input will only accept a single input file and the output is printed to sdtout

## Makefile
    - The make file has a few helper commands that can be used instead of the command in the build/run and tests step.
```
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
```
## Building and running the code
    - The make file included in the file allows you to build the code for the current env.
    - Make sure to be in the GOPATH

```
    go build -o bin/ft
    ./bin/ft input.txt
```

## Testing the code
    - Running the tests can be done with the following command

```
    go test -v ./...
```

## Design decisions
    - The program was split into 4 packages. Handler package, Person package, Relation package and Familytree package
    - The Person package holds the interface and the structure that holds what it means to be a Person. In the family tree, each node is considered to be a struct that implements the Person interface.
    - The Familytree package holds the data for everyone in the tree. The data is stored as hash map using the name of the Person as a key.
    - The Relations package is where we implement each relationship. With this method we can add new relationships that implement the Relationship interface and use them to assert certain relationships for the validity.
    - The seeder files hold the data we use to populate the tree for the first use. This is reused in some tests as well for simplicity.
    - The handlers are the functions what we use to process the input buffers, and generate the instructions for the application to use to map into which calls to make to other packages.
    - We have bubled up the file read operation out of the packages to make the packages independent of where the input source is, although for now we have an explicit file read in the main method we are free to change that to anything that implements the io.Reader interface.

    ## Tests
    - I have added 100% test coverage on all packages. Except for the main method.