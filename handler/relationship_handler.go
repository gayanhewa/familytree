package handler

import (
	"strings"

	"github.com/gayanhewa/family-tree/relation"
	familytree "github.com/gayanhewa/family-tree/tree"
)

// RelationshipHandler resolve relationships.
func RelationshipHandler(familyTree familytree.Tree, name string, relationship string) (output []string) {
	relation := relation.NewRelationFactory(relationship)
	relations, err := relation.GetRelationship(familyTree, name)
	if err != nil {
		output = append(output, err.Error())
		return output
	}
	if len(relations) == 0 {
		output = append(output, "NONE")
		return output
	}
	line := []string{}
	for _, r := range relations {
		line = append(line, r.Name())
	}
	output = append(output, strings.Join(line, " "))
	return output
}
