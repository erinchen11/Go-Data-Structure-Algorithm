package linearsearch

func LinearSearch(inputArr []int, target int) bool {
	for _, item := range inputArr {
		if item == target {
			return true
		}

	}
	return false
}
