package familytree

import "github.com/gayanhewa/family-tree/person"

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
