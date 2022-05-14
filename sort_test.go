package sorts

import "testing"

func TestSort(t *testing.T) {
	var s Sort
	s = new(HeapSort)
	arr := []int{4,3,1,2,5,2,5,7,4,33,32}
	s.Sort(arr)
	t.Log(s)
}
