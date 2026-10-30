package main
import "fmt"
 
func main() {
	var a int
	fmt.Scan(&a)
	for i := 0; i < a; i++ {
		var b int
		fmt.Scan(&b)
 
		var min, max int
		fmt.Scan(&min)
		max = min
 
		for j := 1; j < b; j++ {
			var x int
			fmt.Scan(&x)
			if x < min {
				min = x
			}
			if x > max {
				max = x
			}
		}
 
		hasil := min
		for hasil < max {
			hasil |= hasil + 1
		}
		fmt.Println(hasil)
	}
}
