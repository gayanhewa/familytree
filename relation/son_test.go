package relation

import "testing"

func TestGetRelationshipForSon(t *testing.T) {
	tree := testData()
	s := Son{}
	people, err := s.GetRelationship(tree, "Mother")
	if err != nil {
		t.Fatalf("failed to assert the relationship, %s", err.Error())
	}
	if len(people) != 2 {
		t.Fatal("failed asserting that Mother has 2 sons.")
	}
	if people[0].Name() != "Child1" && people[1].Name() != "Child1" {
		t.Fatal("failed asserting the son's name is Child1")
	}
	if people[0].Name() != "Child4" && people[1].Name() != "Child4" {
		t.Fatal("failed asserting the son's name is Child4")
	}
}
