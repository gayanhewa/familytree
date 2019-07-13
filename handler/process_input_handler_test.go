package handler

import (
	"errors"
	"strings"
	"testing"
)

type stubReader struct{}

func (s *stubReader) Read(p []byte) (int, error) {
	return 0, errors.New("error")
}
func TestHandleFileRead(t *testing.T) {
	t.Run("reading the input buffer successfully", func(t *testing.T) {
		output, _ := ProcessInput(strings.NewReader("GET_RELATIONSHIP James Brother-in-law\nGET_RELATIONSHIP Flora Siblings"))
		if output[0][0] != "GET_RELATIONSHIP" && output[0][1] != "James" && output[0][2] != "Brother-in-law" {
			t.Fatal("failed to assert the data from the file matched GET_RELATIONSHIP James Brother-in-law")
		}
		if output[1][0] != "GET_RELATIONSHIP" && output[1][1] != "Flora" && output[1][2] != "Siblings" {
			t.Fatal("failed to assert the data from the file matched GET_RELATIONSHIP James Brother-in-law")
		}
	})

	t.Run("simulate a scanner error", func(t *testing.T) {
		_, err := ProcessInput(&stubReader{})
		if err != nil && err.Error() != "unable to read the input" {
			t.Fatal("failed to assert that unable to read the input")
		}
	})

}
