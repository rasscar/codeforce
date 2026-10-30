package main
 
import (
	"bufio"
	"fmt"
	"os"
	"sort"
)
 
func sc(a []int, k int) int {
	n := len(a)
	s := make([]bool, n+1)
	f := make([]int, n)
	m := 0
	for i := 0; i < n; i++ {
		s[a[i]] = true
		for m <= n && s[m] {
			m++
		}
		f[i] = m
	}
	sort.Sort(sort.Reverse(sort.IntSlice(f)))
	return f[k-1]
}
 
func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
 
	var t int
	fmt.Fscan(r, &t)
	for ; t > 0; t-- {
		var n, k int
		fmt.Fscan(r, &n, &k)
 
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(r, &a[i])
		}
 
		os := sc(a, k)
		bs := os
		bi, bj := -1, -1
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				a[i], a[j] = a[j], a[i]
				cs := sc(a, k)
				if cs > bs {
					bs = cs
					bi, bj = i, j
				}
				a[i], a[j] = a[j], a[i]
			}
		}
 
		if bs == os {
			fmt.Fprintln(w, -1)
		} else {
			fmt.Fprintf(w, "%d %d\n", bi+1, bj+1)
		}
	}
}
