package main
 
import "fmt"
 
func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var s string
		fmt.Scan(&s)
		n := len(s)
		run, pos, pjg := 1, 0, 1
		ge3, ins := false, false
		for x := 1; x < n; x++ {
			if s[x] != s[x-1] {
				run++
			}
		}
		for x := 1; x <= n; x++ {
			if x < n && s[x] == s[x-1] {
				pjg++
			} else {
				if pjg >= 3 {
					ge3 = true
				}
				if pjg == 1 && pos > 0 && pos < run-1 {
					ins = true
				}
				pjg = 1
				pos++
			}
		}
		h := 1
		if ge3 {
			h++
		}
		if ins {
			h++
		}
		fmt.Println(h)
	}
}
