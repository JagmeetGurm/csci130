package main
import "fmt"

func half( num int) (int, bool){
var b bool
if num%2==0 {
b=true
} else if num%2==1 {
b=false 
}
return num/2, b
}

func main() {
num:=8
a, b:=half(num)
fmt.Println(a,  b)
}
