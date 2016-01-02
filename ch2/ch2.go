package ch2

// Basic insertion sort.
func InsertSort(a []int) {
	for j := 1; j < len(a); j++ {
		key := a[j]
		i := j - 1
		for i > -1 && a[i] > key {
			a[i + 1] = a[i]
			i -= 1
		}
		a[i + 1] = key
	}
}
