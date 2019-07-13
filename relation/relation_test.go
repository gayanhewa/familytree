package relation

import (
	"reflect"
	"testing"

	"github.com/gayanhewa/family-tree/person"
	familytree "github.com/gayanhewa/family-tree/tree"
)

func testData() familytree.Tree {
	tree := familytree.NewFamilyTree()
	father := person.NewPerson("Father", "Male")
	mother := person.NewPerson("Mother", "Female")
	child1 := person.NewPerson("Child1", "Male")
	child1Spouse := person.NewPerson("Child1Spouse", "Female")
	grandchild1 := person.NewPerson("GrandChild1", "Male")
	child2 := person.NewPerson("Child2", "Female")
	child2Spouse := person.NewPerson("Child2Spouse", "Male")
	grandchild2 := person.NewPerson("GrandChild2", "Male")
	child3 := person.NewPerson("Child3", "Female")
	child3Spouse := person.NewPerson("Child3Spouse", "Male")
	child4 := person.NewPerson("Child4", "Male")
	child4Spouse := person.NewPerson("Child4Spouse", "Female")

	father.SetSpouse(mother)
	mother.SetSpouse(father)
	child1.SetMother(mother)
	child1.SetFather(father)
	child1.SetSpouse(child1Spouse)
	child1Spouse.SetSpouse(child1)
	grandchild1.SetFather(child1)
	grandchild1.SetMother(child1Spouse)
	child2.SetMother(mother)
	child2.SetFather(father)
	grandchild2.SetFather(child2Spouse)
	grandchild2.SetMother(child2)
	child2.SetSpouse(child2Spouse)
	child2Spouse.SetSpouse(child2)
	child3.SetMother(mother)
	child3.SetFather(father)
	child3.SetSpouse(child3Spouse)
	child3Spouse.SetSpouse(child3)
	child4.SetMother(mother)
	child4.SetFather(father)
	child4.SetSpouse(child4Spouse)
	child4Spouse.SetSpouse(child4)

	// Add everyone to the tree
	tree.Add(father)
	tree.Add(mother)
	tree.Add(child1)
	tree.Add(child2)
	tree.Add(child3)
	tree.Add(child4)
	tree.Add(child1Spouse)
	tree.Add(child2Spouse)
	tree.Add(child3Spouse)
	tree.Add(child4Spouse)
	tree.Add(grandchild1)
	tree.Add(grandchild2)
	return tree
}

func TestRelationFactory(t *testing.T) {
	tests := []struct {
		name       string
		structType Relationship
	}{
		{"mother", &Mother{}},
		{"father", &Father{}},
		{"paternal-uncle", &PaternalUncle{}},
		{"paternal-aunt", &PaternalAunt{}},
		{"maternal-uncle", &MaternalUncle{}},
		{"maternal-aunt", &MaternalAunt{}},
		{"sister-in-law", &SisterInLaw{}},
		{"brother-in-law", &BrotherInLaw{}},
		{"son", &Son{}},
		{"daughter", &Daughter{}},
		{"siblings", &Sibling{}},
	}
	for _, test := range tests {
		relation := NewRelationFactory(test.name)
		if reflect.TypeOf(relation) != reflect.TypeOf(test.structType) {
			t.Fatalf("failed asserting that %v is the same type of %v", reflect.TypeOf(relation), reflect.TypeOf(test.structType))
		}
	}
	if NewRelationFactory("non-existing name") != nil {
		t.Fatal("failed to assert that invalid type returns nil")
	}

}
