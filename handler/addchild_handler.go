package handler

import (
	"strings"

	familytree "github.com/gayanhewa/family-tree/tree"
)

// AddChildHandler resolves adding a child.
func AddChildHandler(tree familytree.Tree, motherName string, childsName string, gender string) (output []string) {
	motherName = strings.TrimSpace(motherName)
	childsName = strings.TrimSpace(childsName)
	gender = strings.ToLower(strings.TrimSpace(gender))
	// if the gender is not specified correctly skip adding the child.
	if (gender != "female") && (gender != "male") {
		output = append(output, "CHILD_ADDITION_FAILED")
		return output
	}
	// everything is in order so, add the child to the tree.
	if err := tree.AddChild(motherName, childsName, gender); err != nil {
		output = append(output, err.Error())
		return output
	}
	output = append(output, "CHILD_ADDED")
	return output
}
