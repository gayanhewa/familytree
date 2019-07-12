package person

import "testing"

type personStub struct {
	name   string
	gender string
}

var tests = []struct {
	description string
	given       personStub
}{
	{
		"random person male",
		personStub{"random", "male"},
	},
	{
		"random person female",
		personStub{"random female", "female"},
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
