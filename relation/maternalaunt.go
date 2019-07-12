package relation

import (
	"github.com/gayanhewa/family-tree/person"
	familytree "github.com/gayanhewa/family-tree/tree"
)

// MaternalAunt relationship struct.
type MaternalAunt struct{}

// GetRelationship will return the relation Person or if not available an error.
func (m *MaternalAunt) GetRelationship(tree familytree.Tree, name string) (people []person.Person, err error) {
	person, err := tree.GetPerson(name)
	if err != nil {
		return nil, err
	}
	for _, auntOrUncle := range person.Mother().Siblings() {
		if auntOrUncle.Gender() == "female" {
			people = append(people, auntOrUncle)
		}
	}
	return
}
