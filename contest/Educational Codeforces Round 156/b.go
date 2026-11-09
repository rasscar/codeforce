package main
 
import (
	"bufio"
	"fmt"
	"os"
)
 
func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
 
	var t, tc int
	fmt.Fscan(r, &t)
 
	for tc = 0; tc < t; tc++ {
		var n int
		fmt.Fscan(r, &n)
 
		var a []int
		a = make([]int, n)
 
		var i, j, p, cnt, best int
		for i = 0; i < n; i++ {
			fmt.Fscan(r, &a[i])
		}
 
		best = 0
		for i = 0; i < n; i++ {
			p = a[i]
			cnt = 0
			for j = 0; j < n; j++ {
				if a[j]%p == 0 {
					cnt++
				}
			}
			if cnt > best {
				best = cnt
			}
		}
 
		fmt.Fprintln(w, n-best)
	}
}
