package relation

import (
	"testing"
)

func TestGetRelationshipForMother(t *testing.T) {
	tree := testData()
	m := Mother{}
	people, err := m.GetRelationship(tree, "Child1")
	if err != nil {
		t.Fatal("failed to assert the relationship")
	}

	if len(people) > 1 {
		t.Fatal("failed to assert that there is exactly one mother")
	}

	for _, person := range people {
		if person.Name() != "Mother" {
			t.Fatalf("failed to assert the relationship Mother is mother of Child1")
		}
	}
}
