package main
import "fmt"

func main() {
    var H, P, F, D int
    fmt.Scan(&H, &P, &F, &D)
    
    var distanciaH, distanciaP int

    if D == -1 {
        distanciaH = ((F - H) + 16) % 16
        distanciaP = ((F - P) + 16) % 16
    } else {
        distanciaH = ((H - F) + 16) % 16
        distanciaP = ((P - F) + 16) % 16
    }

    if distanciaH < distanciaP {
        fmt.Println("S")
    } else {
        fmt.Println("N")
    }
}
