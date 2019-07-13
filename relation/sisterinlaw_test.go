package relation

import (
	"testing"
)

func TestGetRelationshipForSisterInLaw(t *testing.T) {
	tree := testData()
	s := SisterInLaw{}
	t.Run("test when the relationship exists", func(t *testing.T) {
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
	})
	t.Run("test when the person does not exist", func(t *testing.T) {
		if _, err := s.GetRelationship(tree, "i-do-not-exist"); err == nil {
			t.Fatal("failed to assert error when the person does not exist")
		}
	})
}
