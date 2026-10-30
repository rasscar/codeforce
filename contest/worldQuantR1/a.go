package main
 
import "fmt"
 
func main() {
	var t int
	fmt.Scan(&t)
 
	for i := 0; i < t; i++ {
		var n int
		var s string
		fmt.Scan(&n, &s)
 
		awal := -1
		akhir := -1
		jumlah := 0
 
		for j := 0; j < n; j++ {
			if s[j] == '1' {
				if jumlah == 0 {
					awal = j
				}
				akhir = j
				jumlah++
			}
		}
 
		if jumlah == 1 {
			fmt.Println(0)
		} else if jumlah == 2 {
			fmt.Println(1)
		} else {
			fmt.Println(akhir - awal)
		}
	}
}
