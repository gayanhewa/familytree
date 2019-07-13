package familytree

import (
	"errors"
	"fmt"

	"github.com/gayanhewa/family-tree/person"
)

// Tree interface defines the contract for a family tree implementation.
type Tree interface {
	Add(person.Person) error
	GetPerson(string) (person.Person, error)
	GetPeople() map[string]person.Person
	AddChild(string, string, string) error
	Exists(string) bool
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
	if _, err := t.GetPerson(name); err == nil {
		fmt.Println(err)
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
	if t.Exists(childsName) {
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
