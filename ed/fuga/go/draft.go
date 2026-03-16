package main
import "fmt"

func AbsInt(n int) int {
    if n < 0 {
        return -n
    }
    return n
}

func main() {
    H, P, D, F:= 0
    

    fmt.Scan(&H)
    fmt.Scan(&P)
    fmt.Scan(&F)
    fmt.Scan(&D)

    distanciaH := 0
    distanciaP := 0

    distanciaH = AbsInt(H - F)
    distanciaP = AbsInt(P - F)

    
}
