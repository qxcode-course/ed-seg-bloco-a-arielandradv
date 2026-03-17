package main
import "fmt"
func main() {
    qtd := 0
    dir := 0
    fmt.Scan(&qtd, &dir)
    cobra := make([]Gomo, qtd)
    for i := range cobra {
        fmt.Scan(&cobra[i].x, &cobra[i].y)
    }

    fmt.Println(qtd, dir, cobra)
    for i:= qtd -1; i > 0; i-- {
        cobra[i] = cobra [i-1]
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
     for_, gomo := range cobra {
        fmt.Printf("%v %v\ngomo.x, gomo.y")
     }
}
