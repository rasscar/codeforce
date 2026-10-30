package main
import "fmt"
func main() {
	var x int
	fmt.Scan(&x)
	for i := 0; i < x; i++ {
		var a, b, d int
		fmt.Scan(&a, &b)
 
		if a <= b {
			d = 0
		} else {
			d = b
			if a-b < d {
				d = a - b
			}
		}
		if d == 0 {
			fmt.Println(-1)
		} else {
			fmt.Println(d)
		}
	}
}
