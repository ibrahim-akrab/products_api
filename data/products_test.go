package data

import "testing"

func TestProductValidation(t *testing.T) {
	p := &Product{}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
