package main
 
import (
	"bufio"
	"fmt"
	"os"
)
 
func solve(reader *bufio.Reader, writer *bufio.Writer) {
	var n int
	fmt.Fscan(reader, &n)
 
	fmt.Fprintf(writer, "? 1 %d\n", n)
	writer.Flush()
 
	var tO int
	fmt.Fscan(reader, &tO)
	if tO == -1 {
		return
	}
 
	if tO == 0 || tO == n {
		fmt.Fprintln(writer, "! NO")
		writer.Flush()
		return
	}
 
	fmt.Fprintf(writer, "? 1 %d\n", tO)
	writer.Flush()
 
	var pO int
	fmt.Fscan(reader, &pO)
	if pO == -1 {
		return
	}
 
	if pO == tO {
		fmt.Fprintln(writer, "! NO")
	} else {
		fmt.Fprintln(writer, "! YES")
	}
	writer.Flush()
}
 
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
 
	var t int
	fmt.Fscan(reader, &t)
 
	for tc := 0; tc < t; tc++ {
		solve(reader, writer)
	}
}
