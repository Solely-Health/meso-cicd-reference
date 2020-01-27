package addition

import (
	"fmt"
	"testing"
)

func TestAddTwoInts(t *testing.T) {
	fmt.Println("...Testing addTwoInts...")
	for _, c := range []struct{ xValue, yValue, want int }{
		{1, 1, 2},
		{1, 2, 3},
		{3, 5, 8},
	} {
		ans := addTwoInts(c.xValue, c.yValue)
		if ans != c.want {
			t.Errorf("add method failed with x value: %q, y value: %q, got: %q, expected: %q", c.xValue, c.yValue, ans, c.want)
		}

	}
}
