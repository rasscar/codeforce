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
		var n int
		fmt.Fscan(r, &n)
 
		var adj [][]int
		adj = make([][]int, n+1)
		var i, u, v int
		for i = 0; i < n-1; i++ {
			fmt.Fscan(r, &u, &v)
			adj[u] = append(adj[u], v)
			adj[v] = append(adj[v], u)
		}
 
		if n%2 == 0 {
			fmt.Fprintln(w, n+1)
			continue
		}
 
		var m int
		m = n / 2
 
		var kanK, kanX [3]int
		kanK[0], kanX[0] = n, n
		kanK[1], kanX[1] = n+1, m+1
		kanK[2], kanX[2] = n+2, 1
 
		var ans int
		ans = n
 
		var idx int
		for idx = 0; idx < 3; idx++ {
			var k, x int
			k = kanK[idx]
			x = kanX[idx]
 
			var comp []int
			comp = make([]int, n+1)
			for i = 1; i <= n; i++ {
				comp[i] = -2
			}
			comp[x] = -1
 
			var cid int
			cid = 0
 
			var j, node, ttg int
			for j = 0; j < len(adj[x]); j++ {
				var nei int
				nei = adj[x][j]
				if comp[nei] == -2 {
					var st []int
					st = make([]int, 0)
					st = append(st, nei)
					comp[nei] = cid
 
					for len(st) > 0 {
						node = st[len(st)-1]
						st = st[:len(st)-1]
						for i = 0; i < len(adj[node]); i++ {
							ttg = adj[node][i]
							if comp[ttg] == -2 {
								comp[ttg] = cid
								st = append(st, ttg)
							}
						}
					}
					cid++
				}
			}
 
			var ada bool
			ada = false
 
			if k == n {
				for i = 1; i <= m; i++ {
					u = i
					v = n - i
					if comp[u] != comp[v] {
						ada = true
						break
					}
				}
			} else if k == n+1 {
				for i = 1; i <= m; i++ {
					u = i
					v = n + 1 - i
					if comp[u] != comp[v] {
						ada = true
						break
					}
				}
			} else {
				for i = 2; i <= m+1; i++ {
					u = i
					v = n + 2 - i
					if comp[u] != comp[v] {
						ada = true
						break
					}
				}
			}
 
			if ada {
				ans = k
				break
			}
		}
 
		fmt.Fprintln(w, ans)
	}
}
