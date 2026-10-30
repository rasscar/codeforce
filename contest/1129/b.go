package main
 
import (
	"bufio"
	"fmt"
	"os"
)
func main() {
reader := bufio.NewReader(os.Stdin)
writer := bufio.NewWriter(os.Stdout)
defer writer.Flush()
var a int
fmt.Scan(&a)
for i := 0; i < a; i++ {
	var n int
	fmt.Fscan(reader, &n)
	min := 1 << 30
		for j := 0; j < n; j++ {
		var x int
		fmt.Fscan(reader, &x)
		if x < min {
		min = x
			}
		}
		if min < n {
			fmt.Println(n)
		} else {
			fmt.Println(0)
		}
}
}
