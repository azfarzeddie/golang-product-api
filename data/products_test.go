package data

import "testing"

func TestCheckValidation(t *testing.T) {
    p := &Product{
        Name: "Azfar",
        Price: 1.0,
        SKU: "abc-dec-kha",
    }
    err := p.Validate()

    if err != nil {
        t.Fatal(err)
    }
}
