package main
import "fmt"
func main() {
    solteiros := make(map[int]int)
    qtd := 0
    pares := 0
    fmt.Scan(&qtd)

    for i := 0; i < qtd; i++ {
        animal := 0
        fmt.Scan(&animal)
        v, existe := solteiros[-animal]
        if existe && v > 0{
            solteiros[-animal] = v -1
            pares ++
        } else {
            solteiros[animal] ++
        }
    }
    fmt.Println(pares)

}
