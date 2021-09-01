package binarysearch



// func BinarySearch(inputArr []int, target int) bool {

// 	sort.Ints(inputArr)

// 	begin := 0
// 	end := len(inputArr) - 1

// 	for begin < end {
// 		median := (begin + end) / 2

// 		if inputArr[median] < target {
// 			//往右找
// 			// 因為 inputArr[median]不等於 target
// 			//所以 右移一個
// 			begin = median + 1
// 		} else {
// 			//往左找
// 			end = median - 1
// 		}
// 	}

// 	//如果陣列長度為0或 沒有找到 回傳false
// 	if begin == len(inputArr) || inputArr[begin] != target {
// 		return false
// 	}

// 	return true

// }

func BinarySearch(arr []int, val int) bool {
	first := 0
	last := len(arr) - 1

	for first <= last {
		mid := (first + last) / 2
		switch {
		case arr[mid] == val:
			return true
		case arr[mid] < val:
			first = mid + 1
		case arr[mid] > val:
			last = mid - 1
		}
	}
	return false
}

// func BinarySearch(inputArr []int, target int) bool {

// 	sort.Ints(inputArr)

// 	begin := 0
// 	end := len(inputArr) - 1

// 	for begin <= end {
// 		middle := (begin + end) / 2

// 		if inputArr[middle] < target {
// 			begin = middle + 1
// 		} else {
// 			end = middle - 1
// 		}
// 	}

// 	if begin == len(inputArr) || inputArr[begin] != target {
// 		return false
// 	}
// 	return true
// }
