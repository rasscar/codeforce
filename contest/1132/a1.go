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
		var n, b, c int64
		fmt.Fscan(r, &n)
 
		b = n / 2
		c = n - b
 
		fmt.Fprintln(w, b*c)
	}
}
