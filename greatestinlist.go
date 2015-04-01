package main
 
import "fmt"

func GreatestNoInList(args ...int) int {
greatest :=0
for _, val:=range args {
if greatest < val {
greatest=val
}
}
return greatest
}
func main(){
fmt.Println(GreatestNoInList(4,5,6,7,89,2,11))
}
