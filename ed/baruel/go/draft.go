package main
import "fmt"
func main() {
    var totalAlbum, possui int
    fmt.Scan(&totalAlbum, &possui)

    figurinhas := make([]int, possui)
    for i := range figurinhas {
        fmt.Scan(&figurinhas[i])
    }
    
    var repetidas []int
    for i := 0; i < len(figurinhas)-1; i++ {
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
    
}
