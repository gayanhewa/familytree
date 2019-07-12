package relation

import (
	"github.com/gayanhewa/family-tree/person"
	familytree "github.com/gayanhewa/family-tree/tree"
)

// Father relationship struct.
type Father struct{}

// GetRelationship will return the relation Person or if not available an error.
func (m *Father) GetRelationship(tree familytree.Tree, name string) (people []*person.Person, err error) {
	person, err := tree.GetPerson(name)
	if err != nil {
		return nil, err
	}
	if person.Father() != nil {
		people = append(people, person.Father())
		return people, nil
	}
	return
}
