package main

import "fmt"

//use of IOTA
const (
PI=3.14
Language="go"
A=iota//2
B=iota//3
C=iota//4
)
//iota short-hand
const (
X=iota
Y
)
//no classes in go but there are types
type Contact struct{
name string
greeting string
age int
}

func Intro(person Contact) {
fmt.Println(person.name)
fmt.Println(person.age)
fmt.Println(person.greeting)
}

//calling function in another function
func Greet(per Contact) {
greet,_:=(CreateMessage(per.greeting,per.name))//two values are returned but 1 is used by use of _
fmt.Println(greet)
//fmt.Println(name)//ignoring unused variable by use of _
//fmt.Println(num)
}
func CreateMessage(greeting,name string) (string, string){
return greeting  , name

}
//multiple returns
func Greet2(person Contact){
greeting,name,num:=(Create(person.greeting,person.name))
fmt.Println(greeting)
fmt.Println(name)
fmt.Println(num) //if we make any one of these as comments, it will give us error because we are not using //the returned value

}
func Create(greeting, name string) (string, string,int){
num:=100
return greeting +" " + name ,"\nHi"+name+"\n", num  //use of escape or newline character
}


func Greet3(per Contact) {
myg, myn:=Message(per.greeting, per.name)
fmt.Println(myg)
fmt.Println(myn)
}
//naming return values
func Message(greet, name string) (mygreet, myname string){
mygreet=greet + " "+ name
myname="\nHey," + name
return
}

//VARIADIC functions , allows to have many parameters passed by of ssame type
func Greet4(per Contact){
myg,myn:=CreateMsg(per.greeting,per.name,"Jen")
fmt.Println(myg, myn)

}
func CreateMsg(name string, greeting ...string) (myg string, myn string){
fmt.Println(len(greeting))//tells no.of variadic parameters
myg=greeting[0] //variadic must be last parameter passed
myn=" Hey,"+name //can be only one varidic per function
return
}
//passing function as a parameter
func Greet5(per Contact, myfun func(string)){
myg, myn:=CreateMsg(per.name,per.greeting,"howndy")
myfun(myg)
myfun(myn)
}
func myprint(s string){
fmt.Print(s)
}

func main() {
fmt.Println(PI)
fmt.Println(Language)
fmt.Println(A)
fmt.Println(B)
fmt.Println(C)
fmt.Println(X)
fmt.Println(Y)
var c=Contact{
name: "RAvi",
age: 24,
greeting: "Hi",
}

//or this way can also be declared
var d=Contact{}
d.name="JIm"
d.greeting="Yo"//by default age gets assigned 0 when d is created so taht gets used

//and also this way
var t=Contact{"Joe", "Great", 40}
Intro(t)

//or this way
v:=Contact{}
v.name="Lue"
v.age=40
v.greeting="haa"
Intro(c)
Intro(d)
Intro(v)

var x=Contact{}
x.greeting="Hi"
x.name="Justing"
Greet(x)
Greet2(x)
Greet3(x)
var m=Contact{}
m.greeting="Yes"
m.name="Rob"
Greet4(m)
Greet5(m, myprint)
}
