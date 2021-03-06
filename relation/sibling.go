package relation

import (
	"github.com/gayanhewa/family-tree/person"
	familytree "github.com/gayanhewa/family-tree/tree"
)

// Sibling relationship struct.
type Sibling struct{}

// GetRelationship will return the relation Person or if not available an error.
func (m *Sibling) GetRelationship(tree familytree.Tree, name string) (people []person.Person, err error) {
	person, err := tree.GetPerson(name)
	if err != nil {
		return nil, err
	}
	return person.Siblings(), nil
}
