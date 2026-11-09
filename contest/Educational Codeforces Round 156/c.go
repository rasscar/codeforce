package main
 
import (
	"bufio"
	"fmt"
	"os"
)
 
func xor1toN(n int) int {
	if n%4 == 0 {
		return n
	}
	if n%4 == 1 {
		return 1
	}
	if n%4 == 2 {
		return n + 1
	}
	return 0
}
 
func hitungBit(x int) int {
	var c int
	c = 0
	for x > 0 {
		c += x & 1
		x >>= 1
	}
	return c
}
 
func hitungCtz(x int) int {
	var c int
	c = 0
	for (x & 1) == 0 {
		c++
		x >>= 1
	}
	return c
}
 
func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
 
	var noTsts, tc int
	fmt.Fscan(r, &noTsts)
 
	for tc = 0; tc < noTsts; tc++ {
		var n int
		fmt.Fscan(r, &n)
 
		if n == 1 {
			fmt.Fprintln(w, 1)
			continue
		}
 
		var pVal, x, d int
		pVal = 1
		for (pVal << 1) <= n {
			pVal <<= 1
		}
 
		x = xor1toN(n)
 
		d = x ^ (pVal - 2)
		if hitungBit(d)%2 != 0 {
			d = x ^ (pVal - 1)
		}
 
		var p, t, prev, i, bit, cur int
		p = hitungCtz(pVal)
		t = 0
		prev = 0
 
		for i = 0; i < p; i++ {
			bit = (d >> i) & 1
			cur = bit ^ prev
			if cur == 1 {
				t |= 1 << i
			}
			prev = cur
		}
 
		for i = 1; i <= n; i++ {
			if i > 1 {
				w.WriteByte(' ')
			}
			if i == pVal || (t >= 1 && i == t) {
				w.WriteByte('1')
			} else {
				w.WriteByte('0')
			}
		}
		w.WriteByte('\n')
	}
}
