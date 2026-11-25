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
		var x int64
		fmt.Fscan(r, &x)
 
		for m := int64(1); ; m++ {
			var maxDmg int64
			maxDmg = 0
 
			for a := int64(1); a <= m; a++ {
				c := m - a
				if a <= c+1 {
					dmg := a*(c+1) - a*(a-1)/2
					if dmg > maxDmg {
						maxDmg = dmg
					}
				}
			}
 
			if maxDmg >= x {
				fmt.Fprintln(w, m)
				break
			}
		}
	}
}
