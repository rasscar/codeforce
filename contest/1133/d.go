package main
 
import (
	"bufio"
	"fmt"
	"os"
)
 
const MOD = 998244353
 
func main() {
	var r *bufio.Reader = bufio.NewReader(os.Stdin)
	var w *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer w.Flush()
 
	var t, tc int
	fmt.Fscan(r, &t)
 
	for tc = 0; tc < t; tc++ {
		var n int
		fmt.Fscan(r, &n)
 
		var a []int
		a = make([]int, n)
		var mks, i int
		mks = 0
		for i = 0; i < n; i++ {
			fmt.Fscan(r, &a[i])
			if a[i] > mks {
				mks = a[i]
			}
		}
 
		var pos []int
		pos = make([]int, mks+3)
		for i = 0; i < len(pos); i++ {
			pos[i] = -1
		}
		for i = 0; i < n; i++ {
			pos[a[i]] = i
		}
 
		var dp [][3]int64
		dp = make([][3]int64, n)
		var tot int64
		tot = 0
 
		var v, nx int
		var ways int64
 
		for v = 1; v <= mks; v++ {
			i = pos[v]
			if i == -1 {
				continue
			}
 
			if v-1 >= 1 && pos[v-1] != -1 {
				dp[i][0] = (dp[i][0] + 1) % MOD
			}
			if v-2 >= 1 && pos[v-2] != -1 {
				dp[i][2] = (dp[i][2] + 1) % MOD
			}
 
			if v+1 <= mks && pos[v+1] != -1 {
				nx = pos[v+1]
				dp[nx][0] = (dp[nx][0] + dp[i][2]) % MOD
				tot = (tot + dp[i][2]) % MOD
 
				dp[nx][1] = (dp[nx][1] + dp[i][0]) % MOD
				tot = (tot + dp[i][0]) % MOD
			}
 
			if v+2 <= mks && pos[v+2] != -1 {
				nx = pos[v+2]
				ways = (dp[i][0] + dp[i][1]) % MOD
				dp[nx][2] = (dp[nx][2] + ways) % MOD
				tot = (tot + ways) % MOD
			}
		}
 
		var ans, sz2 int64
		ans = int64(n) % MOD
		sz2 = (int64(n) * int64(n-1) / 2) % MOD
		ans = (ans + sz2 + tot) % MOD
 
		fmt.Fprintln(w, ans)
	}
}
