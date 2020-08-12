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
        Status_particion bool
        TipoParticion [1]byte
        TipoAjuste [1]byte
        Inicio_particion int64
        TamanioTotal int64
        NombreParticion [10]byte
}
    //Struct para el MBR
 type  MBR struct{
        MbrTamanio int64
        MbrFechaCreacion [16]byte
        NoIdentificador uint8
        TipoAjuste [1]byte
        Particiones [4]Particion
}


type Integers struct {
	I1 uint16
	I2 int32
	I3 int64
	DOS byte
}

