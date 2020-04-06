package array

import "testing"
import "reflect"

func TestSum(t *testing.T)  {

	t.Run("collection of 5 numbers",func (t *testing.T){
		numbers := []int{1, 2, 3, 4, 5}
		
		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d, given %v", got, want, numbers)
		}
	})

	// t.Run("collection of 3 numbers",func(t *testing.T){
	// 	numbers := []int{1, 2, 3}

	// 	got := Sum(numbers)
	// 	want := 6

	// 	if got != want {
	// 		t.Errorf("got %d want %d, given %v", got, want, numbers)
	// 	}
	// })
}

func TestSumAll(t *testing.T){
	t.Run("sum 2 slices of numbers",func(t *testing.T){
		got := SumAll([]int{1,2},[]int{0,9})
		want := []int{3,9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestSumAllTails(t * testing.T){
	checkSums := func(t *testing.T, want, got []int){
		if !reflect.DeepEqual(want, got){
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func (t *testing.T){
		got := SumAllTails([]int{1,2}, []int{0,9})
		want := []int{2,9}

		checkSums(t, want, got)
	})

	t.Run("safely sum empty slices", func (t *testing.T){
		got := SumAllTails([]int{},[]int{1,3,5})
		want := []int{0,8}

		checkSums(t, want, got)
	})
}