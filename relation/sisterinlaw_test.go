package relation

import (
	"testing"
)

func TestGetRelationshipForSisterInLaw(t *testing.T) {
	tree := testData()
	s := SisterInLaw{}
	t.Run("test in-laws from spouse", func(t *testing.T) {
		people, err := s.GetRelationship(tree, "Child2Spouse")
		if err != nil {
			t.Fatalf("failed to assert the relationship, %s", err.Error())
		}
		if len(people) != 2 {
			t.Fatal("failed asserting that child1 has 2 Sister in Law")
		}
		if people[0].Name() != "Child3" && people[1].Name() != "Child3" {
			t.Fatal("failed asserting the sister in law's name is Child3")
		}
		if people[0].Name() != "Child6" && people[1].Name() != "Child6" {
			t.Fatal("failed asserting the sister in law's name is Child6")
		}
	})
	t.Run("test person without a spouse", func(t *testing.T) {
		people, err := s.GetRelationship(tree, "Child6")
		if err != nil {
			t.Fatalf("failed to assert the relationship, %s", err.Error())
		}
		if len(people) != 2 {
			t.Fatal("failed asserting that child1 has 2 Sister in Law's")
		}
		if people[0].Name() != "Child1Spouse" && people[1].Name() != "Child1Spouse" {
			t.Fatal("failed asserting the sister in law's name is Child1Spouse")
		}
		if people[0].Name() != "Child4Spouse" && people[1].Name() != "Child4Spouse" {
			t.Fatal("failed asserting the sister in law's name is Child4Spouse")
		}
	})
	t.Run("test in-laws from siblings", func(t *testing.T) {
		people, err := s.GetRelationship(tree, "Child1")
		if err != nil {
			t.Fatalf("failed to assert the relationship, %s", err.Error())
		}
		if len(people) != 1 {
			t.Fatal("failed asserting that child1 has 1 Sister in Law")
		}
		if people[0].Name() != "Child4Spouse" {
			t.Fatal("failed asserting the sister in law's name is Child4Spouse")
		}
	})
	t.Run("test when the person does not exist", func(t *testing.T) {
		if _, err := s.GetRelationship(tree, "i-do-not-exist"); err == nil {
			t.Fatal("failed to assert error when the person does not exist")
		}
	})
}
