package relation

import (
	"testing"
)

func TestGetRelationshipForSisterInLaw(t *testing.T) {
	tree := testData()
	s := SisterInLaw{}
	people, err := s.GetRelationship(tree, "Child1")
	if err != nil {
		t.Fatalf("failed to assert the relationship, %s", err.Error())
	}
	if len(people) != 1 {
		t.Fatal("failed asserting that child1 has 1 Sister in Law")
	}
	if people[0].Name() != "Child4Spouse" {
		t.Fatal("failed asserting the spouse name is Child4Spouse")
	}
}
