package relation

import (
	"github.com/gayanhewa/family-tree/person"
	familytree "github.com/gayanhewa/family-tree/tree"
)

// PaternalUncle relationship struct.
type PaternalUncle struct{}

// GetRelationship will return the relation Person or if not available an error.
func (m *PaternalUncle) GetRelationship(tree familytree.Tree, name string) (people []*person.Person, err error) {
	person, err := tree.GetPerson(name)
	if err != nil {
		return nil, err
	}
	for _, auntOrUncle := range person.Father().Siblings() {
		if auntOrUncle.Gender() == "male" {
			people = append(people, auntOrUncle)
		}
	}
	return
}
