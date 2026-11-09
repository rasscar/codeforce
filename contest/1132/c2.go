package main
 
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
 
func gcd(x int64, y int64) int64 {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
 
func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
 
	var t, tc int
	fmt.Fscan(r, &t)
 
	for tc = 0; tc < t; tc++ {
		var a, b int64
		fmt.Fscan(r, &a, &b)
 
		if a == b {
			w.WriteString("0\n")
			continue
		}
 
		if gcd(a, b) == 1 {
			w.WriteString("1\n")
			w.WriteString(strconv.FormatInt(b-a, 10))
			w.WriteString("\n")
			continue
		}
 
		var ket2 bool
		var lim int64
		ket2 = false
		lim = b
		if a+2000 < b {
			lim = a + 2000
		}
 
		var c int64
		for c = a + 1; c < lim; c++ {
			if gcd(a, c) == 1 && gcd(c, b) == 1 {
				w.WriteString("2\n")
				w.WriteString(strconv.FormatInt(c-a, 10))
				w.WriteString(" ")
				w.WriteString(strconv.FormatInt(b-c, 10))
				w.WriteString("\n")
				ket2 = true
				break
			}
		}
 
		if ket2 {
			continue
		}
 
		var ket3 bool
		var c1, c2 int64
		ket3 = false
		for c1 = a + 1; c1 < b; c1++ {
			if gcd(a, c1) == 1 {
				for c2 = c1 + 1; c2 < b; c2++ {
					if gcd(c1, c2) == 1 && gcd(c2, b) == 1 {
						w.WriteString("3\n")
						w.WriteString(strconv.FormatInt(c1-a, 10))
						w.WriteString(" ")
						w.WriteString(strconv.FormatInt(c2-c1, 10))
						w.WriteString(" ")
						w.WriteString(strconv.FormatInt(b-c2, 10))
						w.WriteString("\n")
						ket3 = true
						break
					}
				}
			}
			if ket3 {
				break
			}
		}
	}
}
