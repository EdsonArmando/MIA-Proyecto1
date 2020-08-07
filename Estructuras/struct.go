package Estructuras
import ( 

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
