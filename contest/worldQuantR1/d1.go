package main
 
import (
	"bufio"
	"fmt"
	"os"
)
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<16)
	writer := bufio.NewWriterSize(os.Stdout, 1<<16)
	defer writer.Flush()
	var a int
	fmt.Fscan(reader, &a)
	for i := 0; i < a; i++ {
		var n, m int
		fmt.Fscan(reader, &n, &m)
		for j := 0; j < m; j++ {
			var b, c int
			fmt.Fscan(reader, &b, &c)
		}
		if n == 1 {
			fmt.Fprintln(writer, 0)
			continue
		}
 
		var ganjil int
		ganjil = (n + 1) / 2
		fmt.Fprintln(writer, ganjil)
 
		for j := 1; j <= ganjil; j++ {
			var genap, b int
			genap = j + 1
			b = j + ganjil
			if b > n {
				b = n
			}
			fmt.Fprintln(writer, genap, b)
		}
	}
}
