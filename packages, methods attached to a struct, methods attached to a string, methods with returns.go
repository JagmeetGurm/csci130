package main

import (                //grouping packages
 "fmt"
 m "go_workspace/math"//package alias
"math"                    //math package
)// notice the import path begins in the folder AFTER the src folder

//STRUCT 
type contact struct{
name string
greeting string
}


type mysentence string

func (s mysentence) eatchoc(){
fmt.Println("Mthod: eat more ")
}

func (c contact) hello(){
fmt.Println("method: "+c.name)
}

//methods with return

func (c contact) helloo() string {
return "method: " +c.name+"sya"
}

//methods with pointers
func (c *contact) rename(newname string){
c.name=newname
}

func main() {
 friend:= contact{"joe", "hello"}
	numberSlice := []int{5, 10, 15}

	sumOfNumbers := m.Sum(numberSlice)

	fmt.Println(sumOfNumbers)
avgofnos:=m.Average(numberSlice)
fmt.Println(avgofnos)
sq:=math.Sqrt(9)
fmt.Println(sq)
friend.hello()
var message mysentence="hello friends"
fmt.Println(message)
message.eatchoc()   //method mas
var mes2 mysentence
mes2.eatchoc()
fmt.Println(friend.helloo())
fmt.Println("earlier name:"+ friend.name)
friend.rename("jen")//by reference
friend.hello()

}

