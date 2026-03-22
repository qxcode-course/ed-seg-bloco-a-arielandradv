 package main
import "fmt"

type Gomo struct{
   x int
   y int
}

func main() {
    var qtd int
    var dir string
    fmt.Scan(&qtd, &dir)

    cobra := make([]Gomo, qtd)
    for i := range cobra {
        fmt.Scan(&cobra[i].x, &cobra[i].y)
    }
    for i:= qtd -1; i > 0 ; i-- {
      cobra[i] = cobra[i-1]
    }

     switch dir {
     case "L":
        cobra[0].x--
     case "R":
        cobra[0].x++
     case "U":
        cobra[0].y--
     case "D":
        cobra[0].y++
     }

     for _, gomo := range cobra {
        fmt.Printf("%d %d\n", gomo.x, gomo.y)
     }
}