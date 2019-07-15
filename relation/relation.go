package relation

import (
	"errors"
	"strings"

	"github.com/gayanhewa/family-tree/person"
	familytree "github.com/gayanhewa/family-tree/tree"
)

// Relationship interface defines the contract for getting a relationship for a given person.
type Relationship interface {
	GetRelationship(tree familytree.Tree, name string) ([]person.Person, error)
}

// NewRelationFactory get the corect relation to resolve.
func NewRelationFactory(name string) (Relationship, error) {
	name = strings.ToLower(strings.TrimSpace(name))
	switch name {
	case "mother":
		return &Mother{}, nil
	case "father":
		return &Father{}, nil
	case "paternal-uncle":
		return &PaternalUncle{}, nil
	case "paternal-aunt":
		return &PaternalAunt{}, nil
	case "maternal-uncle":
		return &MaternalUncle{}, nil
	case "maternal-aunt":
		return &MaternalAunt{}, nil
	case "sister-in-law":
		return &SisterInLaw{}, nil
	case "brother-in-law":
		return &BrotherInLaw{}, nil
	case "son":
		return &Son{}, nil
	case "daughter":
		return &Daughter{}, nil
	case "siblings":
		return &Sibling{}, nil
	}
	return nil, errors.New("NONE")
}
