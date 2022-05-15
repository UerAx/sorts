package sorts

type QuickSort struct {
	arr  []int
	size int
}

func (p *QuickSort) Sort(arr []int) {
	p.add(arr)
	p.quickSort(0, p.size - 1)
}

func (p *QuickSort) add(arr []int) {
	p.arr = arr
	p.size = len(arr)
}

func (p *QuickSort) quickSort(left, right int) {
	if right <= left {
		return
	}
	
	key := p.getKey(left, right)

	l,r := left,right

	for(l < r) {
		for l < r && p.arr[r] >= key {r--}
		p.arr[l] = p.arr[r]
		for l < r && p.arr[l] <= key {l++}
		p.arr[r] = p.arr[l]
	}

	p.arr[l] = key
	p.quickSort(left, l-1)
	p.quickSort(l+1, right)
	

}

func (p *QuickSort) getKey(left, right int) int {
	mid := (left+right)>>1
	p.arr[left], p.arr[mid] = p.arr[mid], p.arr[left]
	return p.arr[left]
}

