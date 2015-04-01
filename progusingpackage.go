package main

import "fmt"
import "golang-book/ch11/math"

func main() {
    xs := []float64{1,2,3,4}
    avg := math.Avg(xs)
    fmt.Println(avg)
min:=math.Min(xs)
fmt.Println(min)
max:=math.Max(xs)
fmt.Println(max)
}
