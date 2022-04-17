package divide

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func Test_quickSort(t *testing.T) {

	for i := 0; i < 5; i++ {
		slice := generateIntSlice()
		target := make([]int, len(slice))
		copy(target, slice)
		sort.Ints(target)
		quickSort(slice)
		if !reflect.DeepEqual(target, slice) {
			t.Errorf("expect %v, got %v", target, slice)
		}
	}

}

func generateIntSlice() []int {
	length := (rand.Int() % 31) + 1
	slice := make([]int, length)
	for i := range slice {
		slice[i] = rand.Int()
	}
	return slice
}
