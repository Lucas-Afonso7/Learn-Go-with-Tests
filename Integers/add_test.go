package inteiros

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d', result '%d'", expected, sum)
	}
}

func ExampleAdiciona() {
    soma := Add(1, 5)
    fmt.Println(soma)
    // Output: 6
}