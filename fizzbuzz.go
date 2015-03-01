package main

import "fmt"

func main() {
    i := 1
    for i <= 20 {
     if (i%3==0 && i%5==0) {
fmt.Println("FizzBuzz")
}else if i%3==0 {
   fmt.Println("Fizz")
     }else if i%5==0 {
   fmt.Println("Buzz")
} 

   i = i + 1
    }
}
