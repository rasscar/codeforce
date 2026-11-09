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
		var n, k int
		fmt.Fscan(r, &n, &k)
 
		var kPar int
		kPar = k % 2
		var cA, cB int
		cA = 0
		cB = 0
 
		var i, val int
		for i = 0; i < n; i++ {
			fmt.Fscan(r, &val)
			if (val % 2) != kPar {
				cA++
			}
		}
 
		for i = 0; i < n; i++ {
			fmt.Fscan(r, &val)
			if (val % 2) != kPar {
				cB++
			}
		}
 
		if cA >= cB {
			fmt.Fprintln(w, "YES")
		} else {
			fmt.Fprintln(w, "NO")
		}
	}
}
