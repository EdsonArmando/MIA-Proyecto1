package Estructuras
import (
	"fmt"
	"strings"
	"os"
	"encoding/binary"
	"container/list"
)

func EjecutarComandoMount(nombreComando string,propiedadesTemp []Propiedad,ListaDiscos *list.List)(ParamValidos bool){
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
	    EjecutarComando(propiedades [1],nombre,ListaDiscos)
	    return ParamValidos
	}else{
		ParamValidos = false
		return ParamValidos
	}
}
func EjecutarReporteMount(ListaDiscos *list.List){
	fmt.Println("----------------- Ejecutando REPORTES MOUNT -----------------")
	for element := ListaDiscos.Front(); element != nil; element = element.Next() {
        var disco  DISCO
        disco = element.Value.(DISCO)
        if disco.NombreDisco != ""{
            fmt.Println("El Disco Montado Es: ",disco.NombreDisco," Identificado con: ",BytesToString(disco.Id))
            fmt.Println("Las particiones Montadas son: ")
            for i:=0;i<len(disco.Particiones);i++{
                var mountTemp = disco.Particiones[i]
                if mountTemp.NombreParticion != ""{
                    fmt.Println(mountTemp.Id,". ",mountTemp.NombreParticion)
                }
            }
            fmt.Println("----------------------")
        }
    }
}
func EjecutarComando(path string,NombreParticion [15]byte,ListaDiscos *list.List)(bool){
	lineaComando := strings.Split(path, "/")
	nombreDisco:= lineaComando[len(lineaComando)-1]
	fmt.Println(nombreDisco)
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
			ParticionMontar(ListaDiscos,string(NombreParticion[:]),string(nombreDisco),path)
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
						fmt.Printf("NombreLogica: %s\n",ebr.NombreParticion)
						//Montar Particion
						ParticionMontar(ListaDiscos,string(NombreParticion[:]),string(nombreDisco),path)
					}
				}
		}
	}
	if err != nil {
		fmt.Println("No existe el archivo en la ruta")
	}
	return true
}

func ParticionMontar(ListaDiscos *list.List,nombreParticion string,nombreDisco string,path string){
	for element := ListaDiscos.Front(); element != nil; element = element.Next() {
		var disco  DISCO
		disco = element.Value.(DISCO)
		if BytesToString(disco.Estado) == "0" && !ExisteDisco(ListaDiscos,nombreDisco){
			disco.NombreDisco = nombreDisco
			disco.Path = path
			copy(disco.Estado[:],"1")
			for i:=0;i<len(disco.Particiones);i++{
				var mountTemp = disco.Particiones[i]
				if BytesToString(mountTemp.Estado) == "0"{
					mountTemp.NombreParticion = nombreParticion
					copy(mountTemp.Estado[:],"1")
					copy(mountTemp.EstadoMKS[:],"0")
					disco.Particiones[i] = mountTemp
					break
				}
			}
			element.Value = disco
			break
		}else if BytesToString(disco.Estado) == "1" && ExisteDisco(ListaDiscos,nombreDisco){
			fmt.Println("Otra particion montada en el disco ", BytesToString(disco.Id))
			for i:=0;i<len(disco.Particiones);i++{
				var mountTemp = disco.Particiones[i]
				if BytesToString(mountTemp.Estado) == "0"{
					mountTemp.NombreParticion = nombreParticion
					copy(mountTemp.Estado[:],"1")
					copy(mountTemp.EstadoMKS[:],"0")
					disco.Particiones[i] = mountTemp
					break
				}
			}
			element.Value = disco
			break
		}
	}
}
func ExisteDisco(ListaDiscos *list.List,nombreDisco string)(bool){
	for element := ListaDiscos.Front(); element != nil; element = element.Next() {
		var disco  DISCO
		disco = element.Value.(DISCO)
		if disco.NombreDisco == nombreDisco{
			return true
		}else{
			return false
		}
	}
	return false
}