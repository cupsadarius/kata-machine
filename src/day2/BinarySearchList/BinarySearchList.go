package binarysearchlist

func search(arr []int, needle, lo, hi int) bool {
	if lo >= hi {
		return false
	}
	mid := lo + (hi-lo)/2
	val := arr[mid]
	if val == needle {
		return true
	}
	if val > needle {
		return search(arr, needle, lo, mid-1)
	}
	return search(arr, needle, mid+1, hi)
}

func BinarySearch(haystack []int, needle int) bool {
	return search(haystack, needle, 0, len(haystack))
}
