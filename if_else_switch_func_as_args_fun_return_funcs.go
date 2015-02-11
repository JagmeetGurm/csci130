package main

import "fmt"
 
type Contact struct{
name string
greeting string
}

//declaring a function type
type myPrintType func(string)

func Greet(person Contact, myfun myPrintType){
myGreet, myname:=CreateMessage(person.name, person.greeting, "howdy")
myfun(myGreet)
myfun(myname)
}

func CreateMessage(nm string, grt ...string) (myG string, myN string){
myG=grt[1]+" " +nm
myN="\nHey, " +nm+"\n"
return
}

//creating some functions to pass
func myPrint(s string){
fmt.Print(s)
}

func myPrint2(s string){
fmt.Print(s)
}

//FUNCTION FACTORY-function that returns a function
func myPrintFunction(custom string) myPrintType{
return func(s string){
fmt.Println(s +custom)
}
}

//use of if struct
func myCondition(mycheck bool){
if mycheck{
fmt.Println("this ran")
}
fmt.Println("wassup")
}

//if variable
func myConditional(myCheck bool){
if myFavName:="Joe"; myCheck{
fmt.Println("this ran,"+myFavName)
}
fmt.Println("wassup")
}

//switching type
func SwitchOnType(x interface{}){
switch x.(type){
case int:
fmt.Println("its an int")
case string:
 fmt.Println("its a string")
}
}


func main(){
var t=Contact{"good to see you", "tim"}
Greet(t,myPrint)
u:=Contact{}
u.greeting="Good You are here,"
u.name="jen"
Greet(u, func(s string){
fmt.Println(s)
})
v:=Contact{}
v.greeting="we are learning"
v.name="Jule"
Greet(v, myPrintFunction("!!"))
Greet(u, myPrintFunction("^^"))
myCondition(true)
myConditional(true)
if(2==3){
fmt.Println("this runs b/c true")
} else {
fmt.Println("this if false")
}

switch "Jen" {
case "Yo":
fmt.Println("hello Yo")
case "Jen":
fmt.Println("hello Jen")
fallthrough //will execute the next case
case "Joe":
fmt.Println("hello Joe")
fallthrough //executes the default as well
default:  //default use is optionald
fmt.Println("no friends")//optional
}

//multiple switch evals
switch "jen" {
case "tim", "joe":
fmt.Println("hello tim or joe")
case "marcus", "moe":
fmt.Println("hello marcus or moe")
case "Jule", "jen": //either one is true so gets executed
fmt.Println("hello Jule or jen")
}

//switch expression not needed

myFriend:="jag"
switch {
case len(myFriend)==2:
fmt.Println("lenght of name is 2")
case myFriend=="joe":
fmt.Println("this is the one")
case myFriend=="jule", myFriend=="jag":
fmt.Println("jag or jule")
}

SwitchOnType(4)
}


