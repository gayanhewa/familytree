package relation

import (
	"github.com/gayanhewa/family-tree/person"
	familytree "github.com/gayanhewa/family-tree/tree"
)

// Son relationship struct.
type Son struct{}

// GetRelationship will return the relation Person or if not available an error.
func (m *Son) GetRelationship(tree familytree.Tree, name string) (people []person.Person, err error) {
	person, err := tree.GetPerson(name)
	if err != nil {
		return nil, err
	}
	for _, child := range person.Children() {
		if child.Gender() == "male" {
			people = append(people, child)
		}
	}
	return
}
