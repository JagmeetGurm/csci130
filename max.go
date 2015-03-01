package main

import "fmt"

func main() {

x := []int{ 48,96,86,68, 57,82,63,70, 37,34,83,27, 19,97, 9,17, }

i:=0
max:=0
for i<len(x) {
if max<x[i] {
max=x[i]
}
i=i+1
}
fmt.Println(max)
}
