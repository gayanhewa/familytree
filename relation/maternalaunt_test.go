package relation

import "testing"

func TestWithNoMaternalAunts(t *testing.T) {
	tree := testData()
	m := MaternalAunt{}
	people, err := m.GetRelationship(tree, "GrandChild1")
	if err != nil {
		t.Fatalf("failed to assert the relationship, %s", err.Error())
	}
	if len(people) != 0 {
		t.Fatal("failed asserting that child1 has 0 maternal aunt's")
	}
}

func TestRelationshipWithMaternalAunts(t *testing.T) {
	tree := testData()
	m := MaternalAunt{}
	people, err := m.GetRelationship(tree, "GrandChild2")
	if err != nil {
		t.Fatalf("failed to assert the relationship, %s", err.Error())
	}
	if len(people) != 1 {
		t.Fatal("failed asserting that child1 has 1 maternal aunt's")
	}
	if people[0].Name() != "Child3" {
		t.Fatal("failed asserting the maternal aunt's name is Child3")
	}
}
