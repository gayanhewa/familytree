package familytree

import (
	"testing"

	"github.com/gayanhewa/family-tree/person"
)

func TestSeeder(t *testing.T) {
	t.Run("test the seeder returns 31 records when the seeder runs without errors", func(t *testing.T) {
		tree := NewFamilyTree()
		tree, err := SeedFamilyTree(tree)
		if err != nil {
			t.Fatal("failed to assert that the seeding completed")
		}
		if len(tree.GetPeople()) != 31 {
			t.Fatal("failed asserting the seeder seeded 31 people")
		}
	})
	t.Run("test when the user already exists", func(t *testing.T) {
		tree := NewFamilyTree()
		tree.Add(person.NewPerson("King Arthur", "Male"))
		if _, err := SeedFamilyTree(tree); err == nil {
			t.Fatal("failed to assert that the seeding failed")
		}
	})
}
