package familytree

import (
	"errors"

	"github.com/gayanhewa/family-tree/person"
)

// Tree interface defines the contract for a family tree implementation.
type Tree interface {
	Add(person.Person) error
	GetPerson(string) (person.Person, error)
	GetPeople() map[string]person.Person
	AddChild(string, string, string) error
}

// FamilyTree holds the family tree structure for an in-memor implementation.
type FamilyTree struct {
	root   person.Person
	people map[string]person.Person
}

// Add handles adding new people to the family tree.
func (t FamilyTree) Add(p person.Person) error {
	// The Person already exists in the tree, name has to be unique so we can identify them easily.
	if _, pre := t.people[p.Name()]; pre == true {
		return errors.New("the name is already taken, use a unique name")
	}
	t.people[p.Name()] = p
	return nil
}

// GetPerson lets you fetch a pointer to a person struct by the name in the hashmap.
func (t FamilyTree) GetPerson(name string) (person.Person, error) {
	person, pre := t.people[name]
	if pre == true {
		return person, nil
	}
	return nil, errors.New("PERSON_NOT_FOUND")
}

// GetPeople returns a map of people in the tree.
func (t FamilyTree) GetPeople() map[string]person.Person {
	return t.people
}

// Exists checks if the user exists.
func (t FamilyTree) Exists(name string) bool {
	if _, err := t.GetPerson(name); err != nil {
		return true
	}
	return false
}

// AddChild to a given parent.
func (t FamilyTree) AddChild(mothersName string, childsName string, childsGender string) error {
	mother, err := t.GetPerson(mothersName)
	if err != nil {
		return err
	}
	if mother.Gender() != "female" {
		return errors.New("CHILD_ADDITION_FAILED")
	}
	// if we alrady have a child with the same name, we will not add them,
	if !t.Exists(childsName) {
		return errors.New("CHILD_ADDITION_FAILED")
	}
	newChild := person.NewPerson(childsName, childsGender)
	newChild.SetMother(mother)
	if mother.HasSpouse() {
		newChild.SetFather(mother.Spouse())
	}
	t.Add(newChild)
	return nil
}

// NewFamilyTree will initialize the family tree into structure. TODO : Refactor
func NewFamilyTree() Tree {
	return &FamilyTree{
		people: map[string]person.Person{},
	}
}

// SeedFamilyTree will seed the family tree with the initial data set.
func SeedFamilyTree(tree Tree) (Tree, error) {
	// This needs to be a param.
	people := []struct {
		name   string
		gender string
		mother string
		father string
		spouse string
	}{
		{"King Arthur", "Male", "", "", "Queen Margret"},
		{"Queen Margret", "Female", "", "", "King Arthur"},
		{"Charlie", "Male", "Queen Margret", "King Arthur", ""},
		{"Bill", "Male", "Queen Margret", "King Arthur", "Flora"},
		{"Flora", "Female", "", "", "Bill"},
		{"Dominique", "Female", "Flora", "Bill", ""},
		{"Louis", "Male", "Flora", "Bill", ""},
		{"Victoire", "Female", "Flora", "Bill", "Ted"},
		{"Ted", "Male", "", "", "Victoire"},
		{"Remus", "Male", "Victoire", "Ted", ""},
		{"Percy", "Male", "Queen Margret", "King Arthur", "Audrey"},
		{"Audrey", "Female", "", "", "Percy"},
		{"Molly", "Female", "Audrey", "Percy", ""},
		{"Lucy", "Female", "Audrey", "Percy", ""},
		{"Ronald", "Male", "Queen Margret", "King Arthur", "Helen"},
		{"Helen", "Female", "", "", "Ronald"},
		{"Rose", "Female", "Helen", "Ronald", "Malfoy"},
		{"Malfoy", "Male", "", "", "Rose"},
		{"Draco", "Ma qle", "Rose", "Malfoy", ""},
		{"Aster", "Female", "Rose", "Malfoy", ""},
		{"Hugo", "Male", "Helen", "Ronald", ""},
		{"Ginerva", "Female", "Queen Margret", "King Arthur", "Harry"},
		{"Harry", "Male", "", "", "Ginerva"},
		{"James", "Male", "Ginerva", "Harry", "Darcy"},
		{"Darcy", "Female", "", "", "James"},
		{"William", "Male", "Darcy", "James", ""},
		{"Albus", "Male", "Ginerva", "Harry", "Alice"},
		{"Alice", "Female", "", "", "Albus"},
		{"Ron", "Male", "Alice", "Albus", ""},
		{"Ginny", "Female", "Alice", "Albus", ""},
		{"Lily", "Female", "Ginerva", "Harry", ""},
	}

	// Load all the people, before adding the relationships
	for _, p := range people {
		if err := tree.Add(person.NewPerson(p.name, p.gender)); err != nil {
			return nil, err
		}
	}
	// Resolve relationships for the first data load
	for _, p := range people {
		person, _ := tree.GetPerson(p.name)
		if p.mother != "" {
			mother, _ := tree.GetPerson(p.mother)
			person.SetMother(mother)
		}
		if p.father != "" {
			father, _ := tree.GetPerson(p.father)
			person.SetFather(father)
		}
		if p.spouse != "" {
			spouse, _ := tree.GetPerson(p.spouse)
			person.SetSpouse(spouse)
		}
	}
	return tree, nil
}
