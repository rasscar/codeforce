package main
 
import "fmt"
 
func main() {
	var n int
	fmt.Scan(&n)
 
	if n == 3 {
		fmt.Println(`111
111
001
 
011
111
100
 
110
110
110
 
010
111
111
 
111
111
111
 
010
111
101
 
111
111
110
 
111
011
111
 
011
010
111`)
		fmt.Println("DURLDULRUU")
		return
	}
 
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			mat := make([][]int, n)
			for r := 0; r < n; r++ {
				mat[r] = make([]int, n)
				for c := 0; c < n; c++ {
					mat[r][c] = 1
				}
			}
 
			if i == n {
				if j == n {
					mat[n-2][n-1] = 0
					mat[n-3][n-1] = 0
					mat[n-4][n-1] = 0
				} else {
					mat[n-2][j-1] = 0
					mat[n-3][j-1] = 0
					mat[n-3][j] = 0
				}
			} else if i == n-1 {
				mat[i-1][j-1] = 0
				mat[i-3][j-1] = 0
				if j > 1 {
					mat[i-2][j-2] = 0
				} else {
					mat[i-3][j] = 0
				}
			} else {
				mat[i+1][j-1] = 0
				if j > 1 {
					mat[i-1][j-2] = 0
					if i > 1 {
						mat[i-2][j-1] = 0
					}
				} else if i > 1 {
					mat[i-2][0] = 0
					mat[i-2][1] = 0
				}
			}
 
			for r := 0; r < n; r++ {
				for c := 0; c < n; c++ {
					fmt.Print(mat[r][c])
				}
				fmt.Println()
			}
			fmt.Println()
		}
	}
 
	panjang := n*n + 1
	s := make([]byte, panjang)
	for i := 0; i < panjang; i++ {
		if i == 0 {
			s[i] = 'U'
		} else if i == 1 {
			s[i] = 'U'
		} else if i == 2 {
			s[i] = 'D'
		} else if i < 3*n {
			if i%3 == 0 {
				s[i] = 'R'
			} else if i%3 == 1 {
				s[i] = 'L'
			} else {
				s[i] = 'R'
			}
		} else {
			if i%n >= n-2 {
				s[i] = 'D'
			} else {
				s[i] = 'U'
			}
		}
	}
	fmt.Println(string(s))
}
