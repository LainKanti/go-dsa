package main

// cmpFunc should return 0 for x == y, 1 for x > y, -1 for x < y
type cmpFunc [T any]func(x, y T) int

// Takes a slice, target value, and comparison function as arguments
// Returns the index the target value was found at
// Return value of -1 means the value was not in the slice
func binarySearch[T any](arr []T, target T, low int, high int, f cmpFunc) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		if f(arr[mid], target) == 0 {
			return mid
		} else if f(arr[mid], target) == -1 {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main(){
	a := []int{1, 2, 3, 4, 5}
    res := binarySearch(a, 2, func(x, y int) int {
        if x == y {
            return 0
        } else if x < y {
            return -1
        } else {
            return 1
        }
    })
}