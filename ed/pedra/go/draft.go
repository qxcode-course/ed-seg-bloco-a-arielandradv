package main
import "fmt"
func main() {
    n := 0
    fmt.Scan(&n)

    type Competidor struct{
        pedraA int
        pedraB int
        diferenca int
        valido bool
    }

    competidores := make([]Competidor, n)

    for i:= range competidores {
        fmt.Scan(&competidores[i].pedraA, &competidores[i].pedraB)

        if competidores[i].pedraA >= 10 && competidores[i].pedraB >= 10 {
            competidores[i].valido = true

            dif:= competidores[i].pedraA - competidores[i].pedraB
            if dif < 0 {
                dif = -dif
            }
            competidores[i].diferenca = dif
        
        } else {
            competidores[i].valido = false
        }
    }

    idVencedor:= -1
    melhorDiferenca:= 101
    for i:= range competidores {
        if competidores[i].valido == true{
            if competidores[i].diferenca < melhorDiferenca {
                melhorDiferenca = competidores[i].diferenca
                idVencedor= i
            }
        }
    }

    if idVencedor == -1{
        fmt.Println("sem ganhador")
    } else {
        fmt.Println(idVencedor)
    }


}