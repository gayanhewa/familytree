package relation

import (
	"github.com/gayanhewa/family-tree/person"
	familytree "github.com/gayanhewa/family-tree/tree"
)

// SisterInLaw relationship struct.
type SisterInLaw struct{}

// GetRelationship will return the relation Person or if not available an error.
func (m *SisterInLaw) GetRelationship(tree familytree.Tree, name string) (people []person.Person, err error) {
	person, err := tree.GetPerson(name)
	if err != nil {
		return nil, err
	}
	// iterate trough the siblings
	for _, sibling := range person.Siblings() {
		// we are only interested in female spouses.
		if sibling.HasSpouse() && sibling.Spouse().Gender() == "female" {
			people = append(people, sibling.Spouse())
		}
	}
	return
}
