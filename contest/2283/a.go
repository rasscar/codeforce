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
		var kiri, kanan int
		fmt.Fscan(r, &kiri, &kanan)

		var n int
		n = kanan - kiri + 1

		if kiri > 0 || kanan < 0 {
			fmt.Fprintln(w, n-1)
			continue
		}

		var a, b, tot, target, mx int64
		a = int64(-kiri)
		b = int64(kanan)
		tot = b*(b+1)/2 - a*(a+1)/2
		if tot == 0 {
			fmt.Fprintln(w, 0)
			continue
		}
		target = tot
		if target < 0 {
			target = -target
		}
		if tot > 0 {
			mx = b
		} else {
			mx = a
		}

		var sum, k, i int64
		sum = 0
		k = 0
		for i = 1; i <= mx; i++ {
			sum += mx - i + 1
			if sum >= target {
				k = i
				break
			}
		}

		var hasil int64
		hasil = 0
		if k-1 > 0 {
			hasil = k - 1
		}
		fmt.Fprintln(w, hasil)
	}
}
