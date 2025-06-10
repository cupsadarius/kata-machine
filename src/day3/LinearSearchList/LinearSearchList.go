package linearsearchlist

func LinearSearch(haystack []int, needle int) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}
	return false
}
