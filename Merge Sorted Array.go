package main

func MergeSortedArray() {
	var nums1 []int = []int{1, 2, 3, 4, 5}
	var nums2 []int = []int{6, 7, 8, 9, 10}
	var m int = 4
	var n int = 5
	for i := n + m - 1; i >= n; i-- {
		nums1[i] = nums1[i-n]
	}
	var (
		i = n
		j = 0
		k = 0
	)
	for k < n+m {
		if i >= n+m {
			nums1[k] = nums2[j]
			k++
			j++
		} else if j >= n {
			break

		} else if nums1[i] < nums2[j] {
			nums1[k] = nums1[i]
			i++
			k++
		} else {
			nums1[k] = nums2[j]
			k++
			j++
		}
	}
}

func main() {
	MergeSortedArray()

}
