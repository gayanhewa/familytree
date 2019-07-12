package person

import "testing"

type person struct {
	name   string
	gender string
}

var tests = []struct {
	description string
	given       person
}{
	{
		"random person male",
		person{"random", "male"},
	},
	{
		"random person female",
		person{"random female", "female"},
	},
}

func TestNewPersonFactory(t *testing.T) {
	for _, test := range tests {
		person := NewPerson(test.given.name, test.given.gender)
		if person.Name() != test.given.name {
			t.Fatalf("failed to assert that %s is %s", person.Name(), test.given.name)
		}
		if person.Gender() != test.given.gender {
			t.Fatalf("failed to assert that %s is %s", person.Gender(), test.given.gender)
		}
	}
}
