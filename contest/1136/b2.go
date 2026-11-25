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
 
	var t int
	fmt.Fscan(r, &t)
 
	for tc := 0; tc < t; tc++ {
		var n, k, y int
		fmt.Fscan(r, &n)
		if n == -1 {
			return
		}
 
		w.WriteString("+")
		for i := 0; i < n-1; i++ {
			w.WriteString("--+")
		}
		w.WriteByte('\n')
		w.Flush()
 
		fmt.Fscan(r, &k)
		if k == -1 {
			return
		}
 
		if k == 0 {
			y = 0
		} else {
			y = (k + 2) / 3
		}
 
		fmt.Fprintln(w, n+y)
		w.Flush()
	}
}
