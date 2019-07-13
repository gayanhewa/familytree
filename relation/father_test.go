package relation

import "testing"

func TestGetRelationshipForFather(t *testing.T) {
	tree := testData()
	f := Father{}
	t.Run("test when the relationship exists", func(t *testing.T) {
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
	})

	t.Run("test when the person does not exist", func(t *testing.T) {
		if _, err := f.GetRelationship(tree, "i-do-not-exist"); err == nil {
			t.Fatal("failed to assert error when the person does not exist")
		}
	})
}
