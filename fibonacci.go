package main

import "fmt"

func fib(num int) int{

if num==0 {
return 0
}else if num==1 {
return 1
}else {
return fib(num-1) + fib(num-2) 
}
}
func main(){//0 1 1 2 3 5
fmt.Println(fib(4))
}
