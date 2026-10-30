package main
 
import (
	"bufio"
	"fmt"
	"os"
)
 
func main() {
	baca := bufio.NewReader(os.Stdin)
	tulis := bufio.NewWriter(os.Stdout)
	defer tulis.Flush()
	var t int
	fmt.Fscan(baca, &t)
 
	for tc := 0; tc < t; tc++ {
		var n, m int
		fmt.Fscan(baca, &n, &m)
 
		adj := make([][]int, n+1)
		for i := 0; i < n-1; i++ {
			var a, b int
			fmt.Fscan(baca, &a, &b)
			adj[a] = append(adj[a], b)
			adj[b] = append(adj[b], a)
		}
 
		vs := make([]int, m)
		for i := 0; i < m; i++ {
			fmt.Fscan(baca, &vs[i])
		}
 
		dist := make([]int, n+1)
		for i := 1; i <= n; i++ {
			dist[i] = -1
		}
 
		q := make([]int, 0, n+1)
		q = append(q, n)
		dist[n] = 0
		for head := 0; head < len(q); head++ {
			u := q[head]
			for _, v := range adj[u] {
				if v == 1 {
					dist[1] = dist[u] + 1
					continue
				}
				if dist[v] == -1 {
					dist[v] = dist[u] + 1
					q = append(q, v)
				}
			}
		}
 
		mks := -1
		for _, v := range vs {
			if dist[v] != -1 && (v != 1 || n == 1) {
				if dist[v] > mks {
					mks = dist[v]
				}
			}
		}
		hasil := dist[1]
		if mks != -1 && 1+mks > hasil {
			hasil = 1 + mks
		}
		fmt.Fprintln(tulis, hasil)
	}
}
