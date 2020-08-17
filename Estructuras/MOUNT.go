package Estructuras
import (
	"fmt"
	"strings"
	"os"
	"encoding/binary"
)

func EjecutarComandoMount(nombreComando string,propiedadesTemp []Propiedad)(ParamValidos bool){
	fmt.Println("----------------- Ejecutando MOUNT -----------------")
	var propiedades [2]string
	var nombre [15]byte
	ParamValidos = true
	if len(propiedadesTemp) >= 2{
	    //Recorrer la lista de propiedades
	    for i:=0; i < len(propiedadesTemp); i++{
	        var propiedadTemp = propiedadesTemp[i]
	        var nombrePropiedad string= propiedadTemp.Name
	        switch strings.ToLower(nombrePropiedad){
	        case "-name":
	        	propiedades [0]=propiedadTemp.Val
	        	copy(nombre[:], propiedades [0])
	        case "-path":
	        	propiedades [1]=propiedadTemp.Val
	    	default:
	    		fmt.Println("Error al Ejecutar el Comando")
	        }
	    }
	    //Empezar a montar las Particiones
	    EjecutarComando(propiedades [1],nombre)
	    fmt.Println("La particion Montada es " ,nombre)
	    return ParamValidos
	}else{
		ParamValidos = false
		return ParamValidos
	}
}

func EjecutarComando(path string,NombreParticion [15]byte)(bool){
	f, err := os.OpenFile(path,os.O_RDONLY,0755)
	if err != nil {
		fmt.Println("No existe la ruta"+path)
		return false
	}
	defer f.Close()
	mbr := MBR{}
	f.Seek(0,0)
	err = binary.Read(f, binary.BigEndian, &mbr)
	fmt.Println("Particiones del disco")
	Particiones := mbr.Particiones
	for i:=0;i<4;i++{
		if string(Particiones[i].NombreParticion[:])== string(NombreParticion[:]){
			fmt.Println("Encontrada")
		}
		if strings.ToLower(BytesToString(Particiones[i].TipoParticion)) == "e"{
			ebr:=EBR{}
			f.Seek(Particiones[i].Inicio_particion,0)
			err = binary.Read(f, binary.BigEndian, &ebr)
				for {
					if ebr.Particion_Siguiente == -1{
						break
					}else{
						f.Seek(ebr.Particion_Siguiente,0)
						err = binary.Read(f, binary.BigEndian, &ebr)
					}
					var nombre string = string(ebr.NombreParticion[:])
					var nombre2 string = string(NombreParticion[:])
					if nombre== nombre2{
						fmt.Println("Encontrada")
					}
				}
		}
	}
	if err != nil {
		fmt.Println("No existe el archivo en la ruta")
	}
	return true
}