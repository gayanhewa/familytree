package person

import "strings"

// Person hold the data for a individual person.
type Person interface {
	AddChild(child Person)
	SetMother(mother Person)
	SetFather(father Person)
	SetSpouse(spouse Person)
	Name() string
	Gender() string
	Mother() Person
	Father() Person
	Spouse() Person
	HasSpouse() bool
	Children() []Person
	Siblings() (siblings []Person)
}

type person struct {
	name     string
	gender   string
	mother   Person
	father   Person
	spouse   Person
	children []Person
}

// NewPerson returns a Person struct initialized with the data.
func NewPerson(name string, gender string) Person {
	return &person{name, gender, nil, nil, nil, []Person{}}
}

// AddChild will update the references to the children.
func (p *person) AddChild(child Person) {
	for _, c := range p.children {
		// exit the adding if the child already exists
		if c == child {
			return
		}
	}
	p.children = append(p.children, child)
}

// SetMother will update the reference to the mother.
func (p *person) SetMother(mother Person) {
	p.mother = mother
	// update the parent with the child
	mother.AddChild(p)
}

// SetFather will update the reference to the father.
func (p *person) SetFather(father Person) {
	p.father = father
	// update the parent with the child
	father.AddChild(p)
}

// SetSpouse will update the reference to the spouse.
func (p *person) SetSpouse(spouse Person) {
	p.spouse = spouse
}

// Name returns the name of the person.
func (p *person) Name() string {
	return p.name
}

// Gender returns the gender of the person.
func (p *person) Gender() string {
	return strings.ToLower(p.gender)
}

// Mother returns the person's mothers name or empty string if not set.
func (p *person) Mother() Person {
	return p.mother
}

// Father returns the person's fathers name or empty string if not set.
func (p *person) Father() Person {
	return p.father
}

// Spouse returns the person's spouses name and empty string if not set.
func (p *person) Spouse() Person {
	return p.spouse
}

// HasSpouse checks if you have a spouse.
func (p *person) HasSpouse() bool {
	return p.spouse != nil
}

// Children returns the child names or empty slice.
func (p *person) Children() []Person {
	return p.children
}

// Siblings returns a Persons siblings if any.
func (p *person) Siblings() (siblings []Person) {
	if p.Mother() != nil {
		for _, child := range p.Mother().Children() {
			if child.Name() == p.Name() {
				continue
			}
			siblings = append(siblings, child)
		}
	}
	return
}
