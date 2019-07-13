package relation

import (
	"github.com/gayanhewa/family-tree/person"
	familytree "github.com/gayanhewa/family-tree/tree"
)

// Mother relationship struct.
type Mother struct{}

// GetRelationship will return the relation Person or if not available an error.
func (m *Mother) GetRelationship(tree familytree.Tree, name string) (people []person.Person, err error) {
	person, err := tree.GetPerson(name)
	if err != nil {
		return nil, err
	}
	if person.Mother() != nil {
		people = append(people, person.Mother())
	}
	return
}
