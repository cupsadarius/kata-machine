package binarysearchlist

func search(arr []int, needle int, lo int, hi int) bool {
	if lo >= hi {
		return false
	}
	middle := lo + (hi-lo)/2
	val := arr[middle]

	if val == needle {
		return true
	}
	if val > needle {
		return search(arr, needle, lo, middle-1)
	}

	return search(arr, needle, middle+1, hi)
}

func BinarySearch(haystack []int, needle int) bool {
	return search(haystack, needle, 0, len(haystack))
}
