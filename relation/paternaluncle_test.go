package relation

import "testing"

func TestRelationshipWithPaternalUncles(t *testing.T) {
	tree := testData()
	p := PaternalUncle{}
	t.Run("test when the relationship exists", func(t *testing.T) {
		people, err := p.GetRelationship(tree, "GrandChild1")
		if err != nil {
			t.Fatalf("failed to assert the relationship, %s", err.Error())
		}
		if len(people) != 2 {
			t.Fatal("failed asserting that child1 has 2 paternal uncle's")
		}
		if people[0].Name() != "Child4" && people[1].Name() != "Child4" {
			t.Fatal("failed asserting the paternal uncles name is Child4")
		}
		if people[0].Name() != "Child5" && people[1].Name() != "Child5" {
			t.Fatal("failed asserting the paternal uncles name is Child5")
		}
	})
	t.Run("test when the person does not exist", func(t *testing.T) {
		if _, err := p.GetRelationship(tree, "i-do-not-exist"); err == nil {
			t.Fatal("failed to assert error when the person does not exist")
		}
	})
}
