package relation

import "testing"

func TestRelationshipWithPaternalUncles(t *testing.T) {
	tree := testData()
	p := PaternalUncle{}
	people, err := p.GetRelationship(tree, "GrandChild1")
	if err != nil {
		t.Fatalf("failed to assert the relationship, %s", err.Error())
	}
	if len(people) != 1 {
		t.Fatal("failed asserting that child1 has 1 paternal uncle's")
	}
	if people[0].Name() != "Child4" {
		t.Fatal("failed asserting the paternal uncles name is Child4")
	}
}
