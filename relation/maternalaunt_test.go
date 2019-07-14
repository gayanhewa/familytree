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
	t.Run("test when the relationship exists", func(t *testing.T) {
		people, err := m.GetRelationship(tree, "GrandChild2")
		if err != nil {
			t.Fatalf("failed to assert the relationship, %s", err.Error())
		}
		if len(people) != 2 {
			t.Fatal("failed asserting that GrandChild2 has 2 maternal aunt's")
		}
		if people[0].Name() != "Child3" && people[1].Name() != "Child3" {
			t.Fatal("failed asserting the maternal aunt's name is Child3")
		}
		if people[0].Name() != "Child6" && people[1].Name() != "Child6" {
			t.Fatal("failed asserting the maternal aunt's name is Child6")
		}
	})
	t.Run("test when the person does not exist", func(t *testing.T) {
		if _, err := m.GetRelationship(tree, "i-do-not-exist"); err == nil {
			t.Fatal("failed to assert error when the person does not exist")
		}
	})
}
