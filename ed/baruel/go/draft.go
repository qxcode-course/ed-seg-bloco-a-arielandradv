package main
import "fmt"
func main() {
    var totalAlbum, possui int
    fmt.Scan(&totalAlbum, &possui)

    // Parte 1
    figurinhas := make([]int, possui)
    for i := range figurinhas {
        fmt.Scan(&figurinhas[i])
    }
    
    var repetidas []int
    for i :=  0; i < len(figurinhas)-1; i++ {
        if figurinhas[i] == figurinhas[i+1]{
            repetidas = append(repetidas, figurinhas[i])
        }
    }

    if len(repetidas) == 0 {
        fmt.Print("N")
    } else {
        for i := range repetidas {
            fmt.Print(repetidas[i])

            if i < len(repetidas)-1 {
                fmt.Print(" ")
            }
        }
    }
    fmt.Println() 

    // Parte 2
    album := make([]bool, totalAlbum+1)
    for _, fig := range figurinhas {
        album[fig] = true
    }

    var faltantes []int
    for i := 1; i <= totalAlbum; i++ {
        if !album[i] {
            faltantes = append(faltantes, i)
        }
    }

    if len(faltantes) == 0 {
        fmt.Print("N")
    } else {
        for i := range faltantes {
            fmt.Print(faltantes[i])

            if i < len(faltantes)-1 {
                fmt.Print(" ")
            }
        }
    }
    fmt.Println()

}
