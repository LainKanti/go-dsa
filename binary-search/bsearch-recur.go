package main
import "fmt"

// cmpFunc should return 0 for x == y, 1 for x > y, -1 for x < y
type cmpFunc [T any]func(x, y T) int

// Expects low to be 0, high to be len(arr) on first call, unfortunately go doesn't have default arguments
func binarySearch[T any](arr []T, target T, low int, high int, f cmpFunc) int {
// TIME O(logN), multiply by cmpFunc time
// SPACE O(logN) due to recursive callstack, multiply by cmpFunc space
    if low > high {
        return -1
    }
    mid := (low + high) / 2
    if f(arr[mid], target) == 0 {
        return mid
    } else if f(arr[mid], target) == -1 {
        return binarySearch(arr, target, mid+1, high, f)
    } else {
        return binarySearch(arr, target, low, mid-1, f)
    }
}

func main() {
    a := []int{1, 2, 3, 4, 5}
    res := binarySearch(a, 2, 0, len(a)-1, func(x, y int) int {
        if x == y {
            return 0
        } else if x < y {
            return -1
        } else {
            return 1
        }
    })
}