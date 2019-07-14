package relation

import (
	"testing"
)

func TestGetRelationshipForBrotherInLaw(t *testing.T) {
	tree := testData()
	b := BrotherInLaw{}
	t.Run("test in-laws from spouse", func(t *testing.T) {
		people, err := b.GetRelationship(tree, "Child1Spouse")
		if err != nil {
			t.Fatalf("failed to assert the relationship, %s", err.Error())
		}
		if len(people) != 2 {
			t.Fatal("failed asserting that Child1Spouse has 2 Brother in Law's")
		}
		if people[0].Name() != "Child4" && people[1].Name() != "Child4" {
			t.Fatal("failed asserting the brother in law's name is Child4")
		}
		if people[0].Name() != "Child5" && people[1].Name() != "Child5" {
			t.Fatal("failed asserting the brother in law's name is Child5")
		}
	})
	t.Run("test person without a spouse", func(t *testing.T) {
		people, err := b.GetRelationship(tree, "Child5")
		if err != nil {
			t.Fatalf("failed to assert the relationship, %s", err.Error())
		}
		if len(people) != 2 {
			t.Fatal("failed asserting that child1 has 2 Brother in Law's")
		}
		if people[0].Name() != "Child2Spouse" && people[1].Name() != "Child2Spouse" {
			t.Fatal("failed asserting the brother in law's name is Child2Spouse")
		}
		if people[0].Name() != "Child3Spouse" && people[1].Name() != "Child3Spouse" {
			t.Fatal("failed asserting the brother in law's name is Child3Spouse")
		}
	})
	t.Run("test in-laws from siblings", func(t *testing.T) {
		people, err := b.GetRelationship(tree, "Child1")
		if err != nil {
			t.Fatalf("failed to assert the relationship, %s", err.Error())
		}
		if len(people) != 2 {
			t.Fatal("failed asserting that Child1Spouse has 0 Brother in Law's")
		}
		t.Log(people)
		if people[0].Name() != "Child3Spouse" && people[1].Name() != "Child3Spouse" {
			t.Fatal("failed asserting the brother in law's name is Child3Spouse")
		}
		if people[0].Name() != "Child2Spouse" && people[1].Name() != "Child2Spouse" {
			t.Fatal("failed asserting the brother in law's name is Child2Spouse")
		}
	})
	t.Run("test when the person does not exist", func(t *testing.T) {
		if _, err := b.GetRelationship(tree, "i-do-not-exist"); err == nil {
			t.Fatal("failed to assert error when the person does not exist")
		}
	})
}
