package relation

import (
	"strings"

	"github.com/gayanhewa/family-tree/person"
	familytree "github.com/gayanhewa/family-tree/tree"
)

// Relationship interface defines the contract for getting a relationship for a given person.
type Relationship interface {
	GetRelationship(tree familytree.Tree, name string) ([]*person.Person, error)
}

// NewRelationFactory get the corect relation to resolve.
func NewRelationFactory(name string) Relationship {
	name = strings.ToLower(strings.TrimSpace(name))
	switch name {
	case "mother":
		return &Mother{}
	case "father":
		return &Father{}
	case "paternal-uncle":
		return &PaternalUncle{}
	case "paternal-aunt":
		return &PaternalAunt{}
	case "maternal-uncle":
		return &MaternalUncle{}
	case "maternal-aunt":
		return &MaternalAunt{}
	case "sister-in-law":
		return &SisterInLaw{}
	case "brother-in-law":
		return &BrotherInLaw{}
	case "son":
		return &Son{}
	case "daughter":
		return &Daughter{}
	case "siblings":
		return &Sibling{}
	}
	return nil
}
