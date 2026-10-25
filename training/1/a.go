package main
import"fmt"
func main(){
var has, a,b,c int
fmt.Scan(&a, &b, &c)
	if (a % c == 0) && (b % c == 0){ 
		has = (b / c) * (a / c)
	}else {
		if a % c == 0 {
			has = a / c
			if b % c == 0{
			has = has * (b/c)
			}else {
			has = has * (b/c+1)
			}
		}else{
			has = a/c +1
			if b % c == 0{
			has = has * (b/c)
			}else {
			has = has * (b/c+1)
			}
		}
	}
fmt.Println(has)
}