package data

import "testing"

func TestProduct_Validate(t *testing.T) {
	p := &Product{}
	p.Name = "Test"
	p.Price = 10.1
	p.SKU = "ads-asd-asd"

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
