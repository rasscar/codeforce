package main
 
import (
	"bufio"
	"fmt"
	"os"
)
 
// helper f
func abs(x int) int { mask := x >> 63; return (x ^ mask) - mask }
 
// ef
 
func cek(X int, n int, a []int) bool {
	var b []int
	b = make([]int, 0)
	var i int
 
	for i = 0; i < n-1; i++ {
		if abs(a[i]-a[i+1]) > X {
			b = append(b, i)
		}
	}
 
	if len(b) == 0 {
		return true
	}
	if len(b) > 2 {
		return false
	}
 
	if len(b) == 2 {
		var l, r int
		l = b[0] + 1
		r = b[1]
		if l > r {
			return false
		}
		for i = l; i < r; i++ {
			if 2*abs(a[i]-a[i+1]) > X {
				return false
			}
		}
		if l > 0 && abs(a[l-1]-2*a[r]) > X {
			return false
		}
		if r < n-1 && abs(2*a[l]-a[r+1]) > X {
			return false
		}
		return true
	}
 
	var idx int
	idx = b[0]
 
	var l1, r1 int
	l1 = idx + 1
	for r1 = l1; r1 < n; r1++ {
		if r1 > l1 && 2*abs(a[r1-1]-a[r1]) > X {
			break
		}
		var ok bool
		ok = true
		if l1 > 0 && abs(a[l1-1]-2*a[r1]) > X {
			ok = false
		}
		if r1 < n-1 && abs(2*a[l1]-a[r1+1]) > X {
			ok = false
		}
		if ok {
			return true
		}
	}
 
	var r2, l2 int
	r2 = idx
	for l2 = r2; l2 >= 0; l2-- {
		if l2 < r2 && 2*abs(a[l2]-a[l2+1]) > X {
			break
		}
		var ok bool
		ok = true
		if l2 > 0 && abs(a[l2-1]-2*a[r2]) > X {
			ok = false
		}
		if r2 < n-1 && abs(2*a[l2]-a[r2+1]) > X {
			ok = false
		}
		if ok {
			return true
		}
	}
 
	return false
}
 
func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
 
	var t, tc int
	fmt.Fscan(r, &t)
 
	for tc = 0; tc < t; tc++ {
		var n, i int
		fmt.Fscan(r, &n)
 
		var a []int
		a = make([]int, n)
		var mks int
		mks = 0
		for i = 0; i < n; i++ {
			fmt.Fscan(r, &a[i])
			if a[i] > mks {
				mks = a[i]
			}
		}
 
		var l, r, m, ans int
		l = 0
		r = 2 * mks
		ans = r
		for l <= r {
			m = l + (r-l)/2
			if cek(m, n, a) {
				ans = m
				r = m - 1
			} else {
				l = m + 1
			}
		}
 
		fmt.Fprintln(w, ans)
	}
}
