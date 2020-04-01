package iteration

import "testing"
import "fmt"

func TestRepeat(t *testing.T){
	repeat := Repeat("a",10)
	expected := "aaaaaaaaaa"

	if repeat != expected {
		t.Errorf("expected %q but got %q", expected, repeat)
	}
}

func BenchmarkRepeat(b *testing.B){
	for i := 0; i < b.N; i++ {
		Repeat("a",10)
	}
}

func ExampleRepeat(){
	repeated := Repeat("a",10)
	fmt.Println(repeated)
	// Output: aaaaaaaaaa
}