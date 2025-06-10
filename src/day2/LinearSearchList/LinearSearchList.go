package linearsearchlist

func LinearSearch(haystack []int, needle int) bool {
	for _, val := range haystack {
		if val == needle {
			return true
		}
	}
	return false
}
