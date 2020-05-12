package main

import "fmt"

func main() {
	test := [][]int{
		{},
		{1},
		{1, 1, 1, 1, 4},
		{25000, 24999, 24998, 24997, 24996, 24995, 24994, 24993, 24992, 24991, 24990, 24989, 24988, 24987, 24986, 24985, 24984, 24983, 24982, 24981, 24980, 24979, 24978, 24977, 24976, 24975, 24974, 24973, 24972, 24971, 24970, 24969, 24968, 24967, 24966, 24965, 24964, 24963, 24962, 24961, 24960, 24959, 24958, 24957, 24956, 24955, 24954, 24953, 24952, 24951, 24950, 24949, 24948, 24947, 24946, 24945, 24944, 24943, 24942, 24941, 24940, 24939, 24938, 24937, 24936, 24935, 24934, 24933, 24932, 24931, 24930, 24929, 24928, 24927, 24926, 24925, 24924, 24923, 24922, 24921, 24920, 24919, 24918, 24917, 24916, 24915, 24914, 24913, 24912, 24911, 24910, 24909, 24908, 24907, 24906, 24905, 24904, 24903, 24902, 24901, 24900, 24899, 24898, 24897, 24896, 24895, 24894, 24893, 24892, 24891, 24890, 24889, 24888, 24887, 24886, 24885, 24884, 24883, 24882, 24881, 24880, 24879, 24878, 24877, 24876, 24875, 24874, 24873, 24872, 24871, 24870, 24869, 24868, 24867, 24866, 24865, 24864, 24863, 24862, 24861, 24860, 24859, 24858, 24857, 24856, 24855, 24854, 24853, 24852, 24851, 24850, 24849, 24848, 24847, 24846, 24845, 24844, 24843, 24842, 24841, 24840, 24839, 24838, 24837, 24836, 24835, 248},
	}

	for _, t := range test {
		fmt.Println(canJump(t))
	}
}

//运行超时（查看官方，贪心算法.查看second）
func canJump(nums []int) bool {
	l := len(nums) - 1
	if l < 1 {
		return true
	}

	queue, m := []int{0}, make(map[int]bool)

	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]
		if q == l {
			return true
		}
		for i := 1; i <= nums[q]; i++ {
			j := i + q
			if j > l {
				break
			}
			if j == l {
				return true
			}
			if m[j] == false {
				queue = append(queue, j)
				m[j] = true
			}
		}
	}

	return false
}