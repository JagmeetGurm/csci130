package main

import "fmt"

func swap(x *int, y *int){
temp:=*x
*x=*y
*y=temp
}

func main(){
a:=5
b:=4

swap(&a,&b)
fmt.Println(a, b)
}
