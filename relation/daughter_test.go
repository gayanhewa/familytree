package relation

import "testing"

func TestGetRelationshipForDaughter(t *testing.T) {
	tree := testData()
	d := Daughter{}
	t.Run("test when relationship exist", func(t *testing.T) {
		people, err := d.GetRelationship(tree, "Mother")
		if err != nil {
			t.Fatalf("failed to assert the relationship, %s", err.Error())
		}
		if len(people) != 3 {
			t.Fatal("failed asserting that Mother has 3 daughters.")
		}
		if people[0].Name() != "Child2" && people[1].Name() != "Child2" && people[2].Name() != "Child2" {
			t.Fatal("failed asserting the daugher's name is Child2")
		}
		if people[0].Name() != "Child3" && people[1].Name() != "Child3" && people[2].Name() != "Child3" {
			t.Fatal("failed asserting the daugher's name is Child3")
		}
		if people[0].Name() != "Child6" && people[1].Name() != "Child6" && people[2].Name() != "Child6" {
			t.Fatal("failed asserting the daugher's name is Child6")
		}
	})
	t.Run("test when the person does not exist", func(t *testing.T) {
		if _, err := d.GetRelationship(tree, "i-do-not-exist"); err == nil {
			t.Fatal("failed to assert error when the person does not exist")
		}
	})
}
