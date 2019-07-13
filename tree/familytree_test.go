package familytree

import (
	"testing"

	"github.com/gayanhewa/family-tree/person"
)

func TestInitialize(t *testing.T) {
	tree := NewFamilyTree()
	if len(tree.GetPeople()) != 0 {
		t.Fatal("failed asserting that the initialization was successful.")
	}
	// test if the seeding works
	tree, _ = SeedFamilyTree(tree)
	if len(tree.GetPeople()) != 31 {
		t.Fatal("failed asserting that the seeding was successful.")
	}
}

func TestAddPersonToTree(t *testing.T) {
	tree, _ := SeedFamilyTree(NewFamilyTree())
	father, _ := tree.GetPerson("King Arthur")
	mother, _ := tree.GetPerson("Queen Margret")
	randomPerson := person.NewPerson("Test Person", "Female")
	randomPerson.SetFather(father)
	randomPerson.SetMother(mother)
	if err := tree.Add(randomPerson); err != nil {
		t.Fatal("failed adding a person with a unique name")
	}
	if err := tree.Add(randomPerson); err == nil {
		t.Fatal("failed asserting that an error is returned when attempting to add a person with an already existing name")
	}
}

func TestGetPersonFromTree(t *testing.T) {
	tree, _ := SeedFamilyTree(NewFamilyTree())
	randomName := "Random non existing person name"
	if _, err := tree.GetPerson(randomName); err == nil {
		t.Fatalf("failed asserting that the person with %s does not exists", randomName)
	}

	tree.Add(person.NewPerson("Test Person", "Female"))
	person, err := tree.GetPerson("Test Person")
	if err != nil {
		t.Fatal("failed asserting that there is no error when fetching a valid person by the name")
	}
	if person.Name() != "Test Person" {
		t.Fatalf("failed asserting that the name is %s equal to Bill", person.Name())
	}
}

func TestCheckForExistingPerson(t *testing.T) {
	tree := NewFamilyTree()
	if tree.Exists("some-random-name") == true {
		t.Fatal("failed to assert that some-random-name does not exist")
	}
	tree.Add(person.NewPerson("some-random-name", "male"))
	if tree.Exists("some-random-name") == false {
		t.Fatal("failed to assert that some-random-name does not exist")
	}
}

func TestAddChild(t *testing.T) {
	tree := NewFamilyTree()
	father := person.NewPerson("Father", "male")
	mother := person.NewPerson("Mother", "female")
	father.SetSpouse(mother)
	mother.SetSpouse(father)
	tree.Add(father)
	tree.Add(mother)
	t.Run("check for errors adding a child", func(t *testing.T) {
		if err := tree.AddChild("Mother", "Child1", "female"); err != nil {
			t.Fatalf("failed asserting that we can add a child to mother %v", err.Error())
		}
	})
	t.Run("test case where the mother doesn't exist", func(t *testing.T) {
		if err := tree.AddChild("Non-existing-mother", "Child1", "female"); err == nil {
			t.Fatal("failed asserting that we can add a child to mother")
		}
	})
	t.Run("test case where the mother isn't a female", func(t *testing.T) {
		if err := tree.AddChild("Father", "Child1", "female"); err == nil {
			t.Fatal("failed asserting that we can add a child to father")
		}
	})
	t.Run("test case where the child is already in the tree", func(t *testing.T) {
		if err := tree.AddChild("Mother", "Child1", "female"); err == nil {
			t.Fatal("failed asserting that we cannot add a child that is already in the tree")
		}
	})
}
