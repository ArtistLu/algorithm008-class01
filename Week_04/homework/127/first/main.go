package main

import "fmt"

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "coo"}
	fmt.Println(ladderLength(beginWord, endWord, wordList))
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	queue1, queue2, re := []string{beginWord}, []string{}, 0

	cur, wordList1 := 1, []string{}
	for len(queue1) > 0 || len(queue2) > 0 {
		q := queue1[0]
		queue1 = queue1[1:]

		if q == endWord {
			re = cur
			break
		}

		for i := 0; i < len(wordList); i++ {
			if isMatch(q, wordList[i]) {
				queue2 = append(queue2, wordList[i])
				continue
			}
			wordList1 = append(wordList1, wordList[i])
		}

		wordList, wordList1 = wordList1, []string{}

		if len(queue1) == 0 {
			queue1, queue2 = queue2, []string{}
			cur++
		}
	}

	return re
}

func isMatch(s, n string) bool {
	var diff int
	for i := 0; i < len(s); i++ {
		if s[i] != n[i] {
			diff++
		}

		if diff > 1 {
			return false
		}
	}

	return diff == 1
}
