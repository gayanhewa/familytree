package relation

import "testing"

func TestWithNoMaternalUncles(t *testing.T) {
	tree := testData()
	m := MaternalUncle{}
	people, err := m.GetRelationship(tree, "GrandChild1")
	if err != nil {
		t.Fatalf("failed to assert the relationship, %s", err.Error())
	}
	if len(people) != 0 {
		t.Fatal("failed asserting that child1 has 0 maternal uncles's")
	}
}

func TestRelationshipWithMaternalUncles(t *testing.T) {
	tree := testData()
	m := MaternalUncle{}
	people, err := m.GetRelationship(tree, "GrandChild2")
	if err != nil {
		t.Fatalf("failed to assert the relationship, %s", err.Error())
	}
	if len(people) != 2 {
		t.Fatal("failed asserting that child1 has 2 maternal uncle's")
	}
	if people[0].Name() != "Child1" && people[1].Name() != "Child1" {
		t.Fatal("failed asserting the maternal uncles name is Child1")
	}
	if people[0].Name() != "Child4" && people[1].Name() != "Child4" {
		t.Fatal("failed asserting the maternal uncles name is Child4")
	}
}
