package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but write %q", expected, repeated)
	}

	fmt.Println(repeated)

}

// func BenchmarkRepeat(b *testing.B) {
// 	for i := 0; i < 5; i++ {
// 		Repeat("a")
// 	}
// }
