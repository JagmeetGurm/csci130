package main

import("fmt"; "math")

type shape interface{
area() float64
perimeter()float64
}

type Circle struct{
x float64
y float64
r float64
}

type Rectangle struct{
x1,y1,x2,y2 float64
}

func(c *Circle) area() float64{
return math.Pi*c.r*c.r
}

func(c *Circle) perimeter() float64{
return math.Pi*c.r*2
}

func (r *Rectangle) perimeter() float64{
return(2* ((r.x2-r.x1)+(r.y2-r.y1)))
}
func (r *Rectangle) area()float64{
return((r.x2-r.x1)*(r.y2-r.y1))
}

func main(){
c:=Circle{0,0,5}
r:=Rectangle{0,0,10,10}
fmt.Println(c.perimeter())
fmt.Println(r.perimeter())
}

