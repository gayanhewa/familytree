package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gayanhewa/family-tree/handler"
	familytree "github.com/gayanhewa/family-tree/tree"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 || args[0] == "" {
		fmt.Println("please specify input file name")
		return
	}
	// read the input file
	file, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// initialize the family tree
	tree, err := familytree.SeedFamilyTree(familytree.NewFamilyTree())
	if err != nil {
		fmt.Println("unable to initialize the family tree")
		return
	}
	var output []string
	processedInputs, err := handler.ProcessInput(file)
	if err != nil {
		fmt.Println("unable to process the input file")
		return
	}
	for _, input := range processedInputs {
		switch input[0] {
		case "ADD_CHILD":
			output = append(output, handler.AddChildHandler(tree, input[1], input[2], input[3])...)
		case "GET_RELATIONSHIP":
			output = append(output, handler.RelationshipHandler(tree, input[1], input[2])...)
		}
	}
	for _, o := range output {
		fmt.Println(o)
	}
}
