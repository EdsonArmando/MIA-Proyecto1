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

type Particion struct{
        status_particion bool
        tipoParticion string
        tipoAjuste string
        inicio_particion uint16
        tamanioTotal uint16
        nombreParticion string
}
    //Struct para el MBR
 type  MBR struct{
        MbrTamanio int64
        MbrFechaCreacion [16]byte
        NoIdentificador uint8
        TipoAjuste [1]byte
}


type Integers struct {
	I1 uint16
	I2 int32
	I3 int64
	DOS byte
}

