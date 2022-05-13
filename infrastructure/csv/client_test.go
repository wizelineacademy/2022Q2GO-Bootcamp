package csv

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	aa := [][]string{
		[]string{"a", "b", "c"},
		[]string{"1", "2", "3"},
		[]string{"4", "5", "6"},
	}

	// fmt.Println(aa)
	fmt.Println(aa[1:][0][1])
}
