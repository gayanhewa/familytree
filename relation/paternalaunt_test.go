package relation

import "testing"

func TestRelationshipWithPaternalAunt(t *testing.T) {
	tree := testData()
	p := PaternalAunt{}
	t.Run("test when the relationship exists", func(t *testing.T) {
		people, err := p.GetRelationship(tree, "GrandChild1")
		if err != nil {
			t.Fatalf("failed to assert the relationship, %s", err.Error())
		}
		if len(people) != 3 {
			t.Fatal("failed asserting that child1 has 3 paternal aunt's")
		}
		if people[0].Name() != "Child2" && people[1].Name() != "Child2" && people[2].Name() != "Child2" {
			t.Fatal("failed asserting the paternal aunt name is Child2")
		}
		if people[0].Name() != "Child3" && people[1].Name() != "Child3" && people[2].Name() != "Child3" {
			t.Fatal("failed asserting the paternal aunt name is Child3")
		}
	})
	t.Run("test when the person does not exist", func(t *testing.T) {
		if _, err := p.GetRelationship(tree, "i-do-not-exist"); err == nil {
			t.Fatal("failed to assert error when the person does not exist")
		}
	})
}
