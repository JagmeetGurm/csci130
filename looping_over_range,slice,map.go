package main

import "fmt"


type contact struct {
	greeting string
	name     string
}


func main() {

	i := 0
	for i < 50 {
		if i>18{//loop break
break
}
if i%2==0{
i++
continue
}
fmt.Println(i)
i++
}
mySlice:=[]int{1,4,8,39,2}
for i, current:= range mySlice {
fmt.Println(i, current)
}
//loop range slice mas
mySlice2:=[]string{ //so it seems that slice acts as vectors; if we give size [4], then it will be aray of size 4
"Good Morning",
"Sat Sri Akaal",
"Jee Ayan Nu",
"Good Bye",
}
for _, current2:= range mySlice2{
fmt.Println(current2)
}

//loop range slice of struct mas
mySlice3:=[]contact{
   {"Tim", "Good Morning"},
   {"Jim", "Sat Sri Akaal"},
   {"Justin", "Good Bye"},
}
for i, current:=range mySlice3{
if i==2{
fmt.Println(current.greeting+", "+current.name)
}
}


//Maps- have keys and values, are like dictionaries, give key and get value
myGreeting:=map[int]string{
1: "Good Morning",
2: "Sat Sri Akaal",
3: "Good Bye",
}
fmt.Println(myGreeting[2])

//MAP INSERT 
myGreeting[5]="hello"
fmt.Println(myGreeting[5])
myGreeting[5]="No, I have changed" //key value updated
fmt.Println(myGreeting[5])

//deleting an entry
val:=myGreeting[5]
val, exists :=myGreeting[5]
fmt.Println("val: ", val)
fmt.Println("exists: " , exists)

delete(myGreeting, 5)
fmt.Println("check for existence after deletion")
if val, exists:=myGreeting[5]; exists{
fmt.Println("not deleted")
}else{
fmt.Println("deleted")
fmt.Println("val: ", val)
fmt.Println("exists: ", exists)
}

//MAPS LOOP RANGE
for keee, valee:=range myGreeting{
fmt.Println(keee, valee)
}

//make map
var famouspeople map[string] string
famouspeople= make(map[string]string)
famouspeople["india"]= "Jag"
fmt.Println(famouspeople["india"])

//SLICES 
myslice5:=[]int {4,2,6,1}
for i, current5:=range myslice5{
fmt.Println(i, "-", current5)
}

//SLICES MAS
famousPerson:=[]string{
"Tom",
"Jerry",
"Pluto",
}
greeting5:=[]string{
"Hello",
"Good",
"bye",
}
for j :=0; j<len(famousPerson); j++{
fmt.Println(famousPerson[j]+ ", "+greeting5[j])
}
//slicing slices
fmt.Println("[1:2]")
fmt.Println(greeting5[1:2])//will only print 1 index value

fmt.Println("[:2]")//means print all indices value upto 2 but not 2
fmt.Println(greeting5[:2])
fmt.Println("[1:]")//will print all index values from 1 onwards
fmt.Println(greeting5[1:])


//SLICING STRING
greeting6:="I am learning go language."
fmt.Println("[4:7]")
fmt.Println(greeting6[4:7])//spaces are also indices in the string

//slices long way
var studentid [] string
studentid=make([]string, 3, 5)//means length is 3 but capacity of underlying array 5

studentid[0]="ab1"
studentid[1]="ab2"
studentid[2]="ab3"
for k:=0; k<len(studentid); k++{
fmt.Println(studentid[k])
}

//SLICING APPENDING
studentid=append(studentid, "ab4")
studentid=append(studentid, "ab5")
fmt.Println(studentid[3])
fmt.Println(studentid[4])

//APPENDING SLICE TO SLICE
myslice7:=[]int{7, 3, 2, 9}
myslice8:=[]int{4,5,6}
myslice7=append(myslice7, myslice8...)

for x, current8:=range myslice7{
fmt.Println(x, "- ",current8)
}
myslice7=append(myslice7[:3], myslice7[5:]...) //slices delete
for y, current9:=range myslice7{
fmt.Println(y, "- ",current9)
}



}
