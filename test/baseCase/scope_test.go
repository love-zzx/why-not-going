package baseCase

import (
	"fmt"
	"testing"
)

func changeValue(p int) {
	p = 10
}

func referenceValue(p *int) {
	*p = 10
}

func TestValuePassing(t *testing.T) {
	var a int = 100
	changeValue(a)
	fmt.Println(a)
}

func TestReference(t *testing.T) {
	var a int = 100
	referenceValue(&a)
	fmt.Println(a)
}

func TestAll(t *testing.T) {
	TestValuePassing(&testing.T{})
	TestReference(&testing.T{})
}
