package main
 
import "fmt"
 
func main() {
    var t int
    fmt.Scan(&t)
    for i := 0; i < t; i++ {
        var n, k int
        fmt.Scan(&n, &k)
        p := make([]int, n)
        freq := make([]int, n+2)
        for j := 0; j < n; j++ {
            fmt.Scan(&p[j])
        }
		for j := 0; j < n; j++ {
			freq[p[j]]++
		}
        hasil := 0
        pl := 0
        for {
            if freq[hasil] > 0 {
                freq[hasil]--
                pl += freq[hasil]
                if pl >= k-1 {
                    pl -= (k-1)
                    hasil++
                } else {
                    break
                }
            } else {
                break
            }
        }
        fmt.Println(hasil)
    }
}
