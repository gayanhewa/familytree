package relation

import "testing"

func TestGetRelationshipForFather(t *testing.T) {
	tree := testData()
	f := Father{}
	people, err := f.GetRelationship(tree, "Child1")
	if err != nil {
		t.Fatal("failed to assert the relationship")
	}

	if len(people) > 1 {
		t.Fatal("failed to assert that there is exactly one father")
	}

	for _, person := range people {
		if person.Name() != "Father" {
			t.Fatalf("failed to assert the relationship Father is the father of Child1")
		}
	}
}
