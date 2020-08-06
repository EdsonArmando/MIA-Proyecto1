package Estructuras
import ( 
    "fmt"
) 
//Estructura para cada Comando y sus Propiedades
type Propiedad struct{
    Name string
    Val  string
}
type Comando struct {
    Name string
    Propiedades []Propiedad
}
// This func must be Exported, Capitalized, and comment added.
func Demo(n int) {
    fmt.Println("HI")
    fmt.Println(n)

}
func Check(e error) {
    if e != nil {
        fmt.Println("Error")
    }
}