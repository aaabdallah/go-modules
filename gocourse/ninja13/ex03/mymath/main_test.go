package mymath

import (
	"fmt"
	"testing"
)

func ExampleCenteredAvg() {

	cavg := CenteredAvg( []int{ 0, 10, 20, 90, 1000 } )
	fmt.Println( cavg )

	// Output: 40
}

func TestCenteredAvg(t *testing.T) {

	cavg := CenteredAvg( []int{ 0, 10, 20, 90, 1000 } )

	if cavg != (10+20+90)/3 {
		t.Error("Got", cavg, "Wanted 40 ((10+20+90)/3)")
	}
}