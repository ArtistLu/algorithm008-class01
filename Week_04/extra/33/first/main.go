package main

func main() {

}

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l )>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < nums[l] && (target >= nums[l] || target < nums[mid]) {
			r = mid - 1
        } else if nums[mid] >= nums[l] && (target < nums[l] || target > nums[mid]) {
			l = mid + 1
		} else if nums[mid] < nums[l] {
			l = mid + 1
	
		} else if nums[mid] >= nums[l] {
			r = mid - 1
		}
	}

	return -1
}
