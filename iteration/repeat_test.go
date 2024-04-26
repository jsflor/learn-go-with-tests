package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("default times", func(t *testing.T) {
		got := Repeat("a", 5)
		wanted := "aaaaa"

		if got != wanted {
			t.Errorf("wanted %q but got %q", wanted, got)
		}
	})
	t.Run("10 times", func(t *testing.T) {
		got := Repeat("a", 10)
		wanted := "aaaaaaaaaa"

		if got != wanted {
			t.Errorf("wanted %q but got %q", wanted, got)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	res := Repeat("a", 5)
	fmt.Println(res)
	// Output: aaaaa
}
