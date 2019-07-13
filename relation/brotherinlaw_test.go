package relation

import (
	"testing"
)

func TestGetRelationshipForBrotherInLaw(t *testing.T) {
	tree := testData()
	b := BrotherInLaw{}
	t.Run("test when the relationships exist", func(t *testing.T) {
		people, err := b.GetRelationship(tree, "Child1")
		if err != nil {
			t.Fatalf("failed to assert the relationship, %s", err.Error())
		}
		if len(people) != 2 {
			t.Fatal("failed asserting that child1 has 2 Brother in Law's")
		}
		if people[0].Name() != "Child2Spouse" && people[1].Name() != "Child2Spouse" {
			t.Fatal("failed asserting the spouse name is Child2Spouse")
		}
		if people[0].Name() != "Child3Spouse" && people[1].Name() != "Child3Spouse" {
			t.Fatal("failed asserting the spouse name is Child3Spouse")
		}
	})
	t.Run("test when the person does not exist", func(t *testing.T) {
		if _, err := b.GetRelationship(tree, "i-do-not-exist"); err == nil {
			t.Fatal("failed to assert error when the person does not exist")
		}
	})
}
