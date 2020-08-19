package word

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {

	cnt := Count("   ABC DEF   GHI TOO FIVE        ")
	if cnt != 5 {
		t.Error("Got", cnt, ", Want 5")
	}

}

func ExampleCount() {

	cnt := Count("   ABC DEF   GHI TOO FIVE        ")
	fmt.Println( cnt )
	// Output: 5
}