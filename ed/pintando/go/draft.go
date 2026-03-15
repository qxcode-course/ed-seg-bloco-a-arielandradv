package main
import (
    "fmt"
    "math"
)

func main() {
    
    var l1, l2, l3 float64
    
    fmt.Scan(&l1)
    fmt.Scan(&l2)
    fmt.Scan(&l3)

    p:= (l1+l2+l3)/2
    resultado:= math.Sqrt(p* (p-l1) * (p-l2) * (p-l3))
    fmt.Printf("%.2f\n", resultado)
}
