package main

func main() {

}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := [26]int{}

	for _, c := range s {
		m[c-'a']++
	}

	for _, c := range t {
		m[c-'a']--
	}
}
