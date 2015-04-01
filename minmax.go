package math



func Avg(xs []float64) float64{
total:=float64(0)
for _, x:=range xs{
total+=x
}
return total/float64(len(xs))
}

func Min(xs []float64) float64{
min:=float64(100)
for _,x:=range xs{
if min>x{
min=x
} 
} 
return min
}

func Max(xs []float64) float64{
max:=float64(0)
for _,x:=range xs{
if max<x{
max=x
} 
} 
return max
}
