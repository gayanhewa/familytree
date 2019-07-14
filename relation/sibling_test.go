package relation

import "testing"

func TestGetRelationshipForSibling(t *testing.T) {
	tree := testData()
	s := Sibling{}
	t.Run("test when the relationship exists", func(*testing.T) {
		people, err := s.GetRelationship(tree, "Child1")
		if err != nil {
			t.Fatalf("failed to assert the relationship, %s", err.Error())
		}
		if len(people) != 5 {
			t.Fatal("failed asserting that Child1 has 3 siblings.")
		}
		if people[0].Name() != "Child2" && people[1].Name() != "Child2" && people[2].Name() != "Child2" {
			t.Fatal("failed asserting the sibling's name is Child2")
		}
		if people[0].Name() != "Child3" && people[1].Name() != "Child3" && people[2].Name() != "Child3" {
			t.Fatal("failed asserting the sibling's name is Child3")
		}
		if people[0].Name() != "Child4" && people[1].Name() != "Child4" && people[2].Name() != "Child4" {
			t.Fatal("failed asserting the sibling's name is Child4")
		}
		if people[0].Name() != "Child4" && people[1].Name() != "Child4" && people[2].Name() != "Child4" {
			t.Fatal("failed asserting the sibling's name is Child4")
		}
		if people[0].Name() != "Child4" && people[1].Name() != "Child4" && people[2].Name() != "Child4" {
			t.Fatal("failed asserting the sibling's name is Child4")
		}
	})
	t.Run("test when the person does not exist", func(t *testing.T) {
		if _, err := s.GetRelationship(tree, "i-do-not-exist"); err == nil {
			t.Fatal("failed to assert error when the person does not exist")
		}
	})
}
