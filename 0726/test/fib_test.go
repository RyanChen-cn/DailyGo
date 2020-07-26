package try

import (
	"testing"
)

func TestFibList(t *testing.T) {
	var (
		a = 1
		b = 1
	)

	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(b)
		a, b = b, a+b
	}
}
