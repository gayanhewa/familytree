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
