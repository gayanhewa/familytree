package handler

import (
	"strings"
	"testing"
)

func TestRelationshipHandler(t *testing.T) {
	tree := testData()
	output := RelationshipHandler(tree, "Child1", "Siblings")
	expected := []string{"Child2", "Child3", "Child4"}
	if output[0] != strings.Join(expected, " ") {
		t.Fatalf("failed to assert the siblings for Child1 %v", expected)
	}

	output = RelationshipHandler(tree, "Child4", "Son")
	if output[0] != "NONE" {
		t.Fatalf("failed to assert the son's for Child4 is NONE")
	}

	output = RelationshipHandler(tree, "Throw and error", "Son")
	if output[0] != "PERSON_NOT_FOUND" {
		t.Fatalf("failed to assert PERSON_NOT_FOUND")
	}

	output = RelationshipHandler(tree, "Child4", "bogus-relation")
	if output[0] != "NONE" {
		t.Fatalf("failed to assert NONE")
	}
}
