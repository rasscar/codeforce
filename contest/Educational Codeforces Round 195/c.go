
package main
import "fmt"
func main(){
var t int
fmt.Scan(&t)
for tc:=0;tc<t;tc++{
var n,k int
fmt.Scan(&n,&k)
a:=make([]int,n)
for i:=0;i<n;i++{fmt.Scan(&a[i])}
p:=make([]int,n+1)
for i:=0;i<n;i++{p[i+1]=(p[i]+a[i])%k}
freq:=make([]int,k)
ans:=0
for r:=k+1;r<=n;r++{
validJ:=r-k-1
freq[p[validJ]]++
ans+=freq[p[r]]
}
fmt.Println(ans)
}
}
