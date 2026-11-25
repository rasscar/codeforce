package main
 
import (
	"bufio"
	"fmt"
	"os"
)
 
func up(nx []bool, a int, b int, lim int, h *int) {
	nx[a*lim+b] = true
	if a+b > *h {
		*h = a + b
	}
}
 
func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
 
	var t int
	fmt.Fscan(r, &t)
 
	for tc := 0; tc < t; tc++ {
		var x, y, m int
		fmt.Fscan(r, &x, &y, &m)
 
		var s string
		fmt.Fscan(r, &s)
 
		lim := 205
		
		dp := make([]bool, lim*lim)
		nx := make([]bool, lim*lim)
		
		dp[x*lim+y] = true
 
		var ans int
		ans = x + y
 
		for i := 0; i < m; i++ {
			for j := 0; j < lim*lim; j++ {
				nx[j] = false
			}
 
			for a := 0; a < lim; a++ {
				for b := 0; b < lim; b++ {
					if !dp[a*lim+b] {
						continue
					}
 
					if s[i] == '+' {
						if b > 0 && b+1 < lim {
							up(nx, a, b+1, lim, &ans)
						}
						if a > 0 && b+1 < lim {
							up(nx, a-1, b+1, lim, &ans)
						}
					} else {
						if b > 0 && a+1 < lim {
							up(nx, a+1, b-1, lim, &ans)
						}
						if a > 0 {
							up(nx, a-1, b, lim, &ans)
						}
					}
				}
			}
			dp, nx = nx, dp
		}
 
		fmt.Fprintln(w, ans)
	}
}
