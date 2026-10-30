package main
 
import (
	"bufio"
	"fmt"
	"os"
)
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<16)
	writer := bufio.NewWriterSize(os.Stdout, 1<<16)
	defer writer.Flush()
	var t int
	fmt.Fscan(reader, &t)
	i := 0
	for {
		if i >= t {
			break
		}
		var n, k, P int
		var rbs string
		fmt.Fscan(reader, &n, &k)
		fmt.Fscan(reader, &rbs)
		P = n / 2
		dep := make([]int, n+1)
		cc := make([]int, n+1)
		iStr := 0
		for {
			if iStr >= len(rbs) {
				break
			}
			dep[iStr+1] = dep[iStr]
			cc[iStr+1] = cc[iStr]
			if rbs[iStr] == '(' {
				dep[iStr+1]++
			} else {
				dep[iStr+1]--
				cc[iStr+1]++
			}
			iStr++
		}
 
		ans := 0
		for p := 1; p <= n; p++ {
			c := cc[p]
			if c > k {
				break
			}
			op := p - c
			extra := k - c
			if P-op < extra {
				extra = P - op
			}
			if op+extra > ans {
				ans = op + extra
			}
		}
		dq := make([]int, n+1)
		head, tail := 0, 0
		x := 0
		for R := 0; R <= n; R++ {
			for tail > head && dep[dq[tail-1]] >= dep[R] {
				tail--
			}
			dq[tail] = R
			tail++
			for cc[R]-cc[x] > k {
				x++
			}
			for dq[head] < x {
				head++
			}
			if dep[R]-dep[dq[head]] > ans {
				ans = dep[R] - dep[dq[head]]
			}
		}
 
		if ans > P {
			ans = P
		}
		fmt.Fprintln(writer, ans)
		i++
	}
}
