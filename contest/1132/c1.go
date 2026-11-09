package main
 
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
 
var isPrime []bool
 
func hitungSieve(maxV int) {
	isPrime = make([]bool, maxV+1)
	var i, p int
	for i = 2; i <= maxV; i++ {
		isPrime[i] = true
	}
	for p = 2; p*p <= maxV; p++ {
		if isPrime[p] {
			for i = p * p; i <= maxV; i += p {
				isPrime[i] = false
			}
		}
	}
}
 
func main() {
	siv := 200000
	hitungSieve(siv)
 
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
 
	var t, tc, n, i int
	fmt.Fscan(r, &t)
 
	for tc = 0; tc < t; tc++ {
		fmt.Fscan(r, &n)
 
		var a []int
		a = make([]int, n)
		for i = 0; i < n; i++ {
			fmt.Fscan(r, &a[i])
		}
 
		var res []int
		res = make([]int, n)
		for i = 0; i < n; i++ {
			var val int
			val = a[i]
 
			if isPrime[val] {
				res[i] = val
			} else {
				var p int
				p = val + 1
				for !isPrime[p] {
					p++
				}
				res[i] = p
			}
		}
 
		for i = 0; i < n; i++ {
			w.WriteString(strconv.Itoa(res[i]))
			if i < n-1 {
				w.WriteByte(' ')
			}
		}
		w.WriteByte('\n')
	}
}
