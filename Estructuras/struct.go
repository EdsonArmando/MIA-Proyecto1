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
//Estructuras para el Disco y Particiones
type Particion struct{
        Status_particion [1]byte
        TipoParticion [1]byte
        TipoAjuste [2]byte
        Inicio_particion int64
        TamanioTotal int64
        NombreParticion [10]byte
}
    //Struct para el MBR
 type  MBR struct{
        MbrTamanio int64
        MbrFechaCreacion [19]byte
        NoIdentificador int64
        Particiones [4]Particion
}
//Struct para las particiones Logicas
type EBR struct{
        Status_particion [1]byte
        TipoAjuste [2]byte
        Inicio_particion int64
        Particion_Siguiente int64
        TamanioTotal int64
        NombreParticion [15]byte
}

type Integers struct {
	I1 uint16
	I2 int32
	I3 int64
	DOS byte
}

