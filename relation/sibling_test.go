package relation

import "testing"

func TestGetRelationshipForSibling(t *testing.T) {
	tree := testData()
	s := Sibling{}
	people, err := s.GetRelationship(tree, "Child1")
	if err != nil {
		t.Fatalf("failed to assert the relationship, %s", err.Error())
	}
	if len(people) != 3 {
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
}
