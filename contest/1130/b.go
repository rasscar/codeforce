package main
 
import (
	"bufio"
	"fmt"
	"os"
)
 
func main() {
	var baca *bufio.Reader = bufio.NewReader(os.Stdin)
	var tulis *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer tulis.Flush()
	var t, tc, n, i, beda int
	var r1, r2 string
	fmt.Fscan(baca, &t)
 
	for tc = 0; tc < t; tc++ {
		fmt.Fscan(baca, &n)
		fmt.Fscan(baca, &r1)
		fmt.Fscan(baca, &r2)
		beda = 0
		for i = 0; i < n; i++ {
			if r1[i] == 'f' && r2[i] == 'o' {
				beda++
			} else if r1[i] == 'o' && r2[i] == 'f' {
				beda--
			}
		}
 
		if beda < 0 {
			beda = -beda
		}
		hasil := beda % 2
		fmt.Fprintln(tulis, hasil)
	}
}
