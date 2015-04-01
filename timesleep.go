package main

import("time")

func SleepProcess(s time.Duration){
<-time.After(s *time.Minute/60)
}

func main(){
//or i:=0; i<10;i++{
SleepProcess(5)
//}
}
