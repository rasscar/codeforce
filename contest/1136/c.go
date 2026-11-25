package main
 
import (
	"bufio"
	"fmt"
	"os"
)
 
func urut(ar []int64, kiri int, kanan int) {
	if kiri >= kanan {
		return
	}
	mid := (kiri + kanan) / 2
	urut(ar, kiri, mid)
	urut(ar, mid+1, kanan)
 
	temp := make([]int64, kanan-kiri+1)
	i := kiri
	j := mid + 1
	k := 0
 
	for i <= mid && j <= kanan {
		if ar[i] <= ar[j] {
			temp[k] = ar[i]
			i++
		} else {
			temp[k] = ar[j]
			j++
		}
		k++
	}
	for i <= mid {
		temp[k] = ar[i]
		i++
		k++
	}
	for j <= kanan {
		temp[k] = ar[j]
		j++
		k++
	}
	for i = 0; i < len(temp); i++ {
		ar[kiri+i] = temp[i]
	}
}
 
func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
 
	var t int
	fmt.Fscan(r, &t)
 
	for tc := 0; tc < t; tc++ {
		var n int
		var x int64
		fmt.Fscan(r, &n, &x)
 
		var h []int64
		h = make([]int64, n)
 
		for i := 0; i < n; i++ {
			var a int64
			fmt.Fscan(r, &a)
			h[i] = (a + x - 1) / x
		}
 
		urut(h, 0, n-1)
 
		var ans int64
		ans = 0
 
		for i := 0; i < n-1; i++ {
			var val int64
			val = h[i] - int64(i+1)
			if val > ans {
				ans = val
			}
		}
 
		fmt.Fprintln(w, ans)
	}
