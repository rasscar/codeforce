package main
 
import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)
 
func urutDesc(ar []int, kiri int, kanan int) {
	if kiri >= kanan {
		return
	}
	var mid int
	mid = (kiri + kanan) / 2
	urutDesc(ar, kiri, mid)
	urutDesc(ar, mid+1, kanan)
 
	var temp []int
	temp = make([]int, kanan-kiri+1)
	var i, j, k int
	i = kiri
	j = mid + 1
	k = 0
 
	for i <= mid && j <= kanan {
		if ar[i] > ar[j] {
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
 
	var t, tc, n int
	fmt.Fscan(r, &t)
 
	for tc = 0; tc < t; tc++ {
		fmt.Fscan(r, &n)
 
		var a []int
		a = make([]int, n)
		var i int
		for i = 0; i < n; i++ {
			fmt.Fscan(r, &a[i])
		}
 
		urutDesc(a, 0, n-1)
 
		var x, y, v *big.Int
		x = big.NewInt(0)
		y = big.NewInt(0)
 
		for i = 0; i < n; i++ {
			v = new(big.Int).Lsh(big.NewInt(1), uint(a[i]))
			
			if x.Cmp(y) <= 0 {
				x.Add(x, v)
			} else {
				y.Add(y, v)
			}
		}
 
		var ans *big.Int
		ans = new(big.Int).Mul(x, y)
		
		fmt.Fprintln(w, ans.String())
	}
