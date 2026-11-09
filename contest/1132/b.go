package main
 
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
 
func sortUnik(arr []int64) {
	if len(arr) <= 1 {
		return
	}
	mid := len(arr) / 2
	left := make([]int64, mid)
	right := make([]int64, len(arr)-mid)
	copy(left, arr[:mid])
	copy(right, arr[mid:])
 
	sortUnik(left)
	sortUnik(right)
 
	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}
	for i < len(left) {
		arr[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		arr[k] = right[j]
		j++
		k++
	}
}
 
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
 
	var t int
	fmt.Fscan(reader, &t)
 
	for i := 0; i < t; i++ {
		var n, k int
		fmt.Fscan(reader, &n, &k)
 
		a := make([]int64, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &a[j])
		}
 
		sortUnik(a)
 
		maxVal := a[n-1]
		c := 0
		for j := 0; j < n; j++ {
			if a[j] == maxVal {
				c++
			}
		}
 
		if k < c {
			writer.WriteString("NO\n")
		} else {
			writer.WriteString("YES\n")
 
			res := make([]int64, 0, n)
 
			for j := n - k; j < n; j++ {
				res = append(res, a[j])
			}
 
			for j := n - k - 1; j >= 0; j-- {
				res = append(res, a[j])
			}
 
			for idx, v := range res {
				writer.WriteString(strconv.FormatInt(v, 10))
				if idx < n-1 {
					writer.WriteString(" ")
				}
			}
			writer.WriteString("\n")
		}
	}
