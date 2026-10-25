package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var names []string
	names = make([]string, n)
	var scores []int
	scores = make([]int, n)
	var total map[string]int
	total = make(map[string]int)

	var i int
	for i = 0; i < n; i++ {
		fmt.Scan(&names[i], &scores[i])
		total[names[i]] += scores[i]
	}

	var terbaik int
	terbaik = -1
	for i = 0; i < n; i++ {
		if total[names[i]] > terbaik {
			terbaik = total[names[i]]
		}
	}

	var skr map[string]int
	skr = make(map[string]int)
	for i = 0; i < n; i++ {
		skr[names[i]] += scores[i]
		if skr[names[i]] >= terbaik && total[names[i]] == terbaik {
			fmt.Println(names[i])
			return
		}
	}
}
