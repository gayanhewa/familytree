package relation

import "testing"

func TestGetRelationshipForDaughter(t *testing.T) {
	tree := testData()
	d := Daughter{}
	people, err := d.GetRelationship(tree, "Mother")
	if err != nil {
		t.Fatalf("failed to assert the relationship, %s", err.Error())
	}
	if len(people) != 2 {
		t.Fatal("failed asserting that Mother has 2 daughters.")
	}
	if people[0].Name() != "Child2" && people[1].Name() != "Child2" {
		t.Fatal("failed asserting the daugher's name is Child2")
	}
	if people[0].Name() != "Child3" && people[1].Name() != "Child3" {
		t.Fatal("failed asserting the daugher's name is Child3")
	}
}
