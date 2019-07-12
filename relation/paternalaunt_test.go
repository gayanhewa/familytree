package relation

import "testing"

func TestRelationshipWithPaternalAunt(t *testing.T) {
	tree := testData()
	p := PaternalAunt{}
	people, err := p.GetRelationship(tree, "GrandChild1")
	if err != nil {
		t.Fatalf("failed to assert the relationship, %s", err.Error())
	}
	if len(people) != 2 {
		t.Fatal("failed asserting that child1 has 2 paternal aunt's")
	}
	if people[0].Name() != "Child2" && people[1].Name() != "Child2" {
		t.Fatal("failed asserting the paternal aunt name is Child2")
	}
	if people[0].Name() != "Child3" && people[1].Name() != "Child3" {
		t.Fatal("failed asserting the paternal aunt name is Child3")
	}
}
