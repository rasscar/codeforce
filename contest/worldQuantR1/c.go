package main
import (
	"bufio"
	"fmt"
	"os"
)
func main(){
	reader := bufio.NewReaderSize(os.Stdin, 1<<16)
	writer := bufio.NewWriterSize(os.Stdout, 1<<16)
	defer writer.Flush()
	var a int
	fmt.Fscan(reader, &a)
	for i := 0; i < a; i++ {
		var ganjil, genap, n int
		fmt.Fscan(reader, &n)
		for j := 0; j < n; j++ {
			var b int
			fmt.Fscan(reader, &b)
			if b%2 == 0 {
				genap++
			} else {
				ganjil++
			}
		}
		if (ganjil == 1 && genap == 1) || (ganjil == 3 && genap == 0) {
			fmt.Fprintln(writer, "Larpity")
		} else {
			fmt.Fprintln(writer, "Moggity")
		}
	}
}
