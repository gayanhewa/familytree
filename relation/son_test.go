package relation

import "testing"

func TestGetRelationshipForSon(t *testing.T) {
	tree := testData()
	s := Son{}
	t.Run("test when the relationship exists", func(t *testing.T) {
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
	})
	t.Run("test when the person does not exist", func(t *testing.T) {
		if _, err := s.GetRelationship(tree, "i-do-not-exist"); err == nil {
			t.Fatal("failed to assert error when the person does not exist")
		}
	})
}
