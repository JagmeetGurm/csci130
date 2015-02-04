package main
import "fmt"

//USER DEFINED TYPE
type salutation string

//using struct like class(outside main function)
type sal struct{
name string
greet string
}

//CAPITALIZATION matters(outside main function)
type Sal struct{
nm string
gr string
}

//CONSTANTS(outside main function);use of round brackets instead of curly 
const( 
Pi=3.14
Language="go"
)

func main() {
//declaring a variable of string type
var message string 
message = "hello go world"
fmt.Println(message)// printing the variable

//intializing many varibles(declaring and assigning them)
var a, b, c int = 1, 2, 3
fmt.Println(a, b, c)

//inferring type
//go infers the type of variable on its own
var msg = "hi there"
var x,y,z=1,2,3
fmt.Println(msg,x,y,z)

//mixing the type inferences 
var u, v, w=1, true, "hi"
fmt.Println(u,v,w)

//getting rid of var
d,e,f:=4,true,4.5
fmt.Println(e,f,d)

//POINTERS- a variable that points to a memory location. has to have a data type
var greeting *string=&message// a pointer pointing to memory location which holds a string
fmt.Println(greeting, *greeting)//1st one simply prints the address of message, 2nd one the value at mesaage
fmt.Println(message, *greeting)//both print the same thing value of message
*greeting="new value "//greeting still points to message but the value stored by mesage has been changed
fmt.Println(*greeting, message)

//USER DEFINED TYPE being used
var new salutation= "Hello there"
fmt.Println(new)

//USING STRUCT LIKE A CLASS
var s=sal{"Joe", "Hi there"}
fmt.Println(s.name, s.greet)

//explicit pass of arguments
var s2=sal{name:"Jill", greet: "Helo"}
fmt.Println(s2.name, s2.greet)

//can use this notation too
var s3=sal{}
s3.name="Rob"
s3.greet="yo"
fmt.Println(s3.name, s3.greet)

//IMPORTANCE OF CAPITALIZATION
//capitalized- public used to put in a package and then import the package
//lowercase-private

var s4=Sal{}
s4.nm="Bob"
s4.gr="haha"
fmt.Println(s4.nm, s4.gr)

fmt.Println(Pi)
fmt.Println(Language)

}
