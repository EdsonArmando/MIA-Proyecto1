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
            for i:=0;i<len(disco.Particiones);i++{
                var mountTemp = disco.Particiones[i]
                if mountTemp.NombreParticion != ""{
                    fmt.Println("->",mountTemp.Id,"->",disco.Path,"->",mountTemp.NombreParticion)
                }
            }
        }
    }
}
func IdValido(id string,ListaDiscos *list.List)(bool){
	esta:=false
	for element := ListaDiscos.Front(); element != nil; element = element.Next() {
        var disco  DISCO
        disco = element.Value.(DISCO)
        if disco.NombreDisco != ""{
            for i:=0;i<len(disco.Particiones);i++{
                var mountTemp = disco.Particiones[i]
                if mountTemp.NombreParticion != ""{
                	if mountTemp.Id == id{
                		return true
                	}
                }
            }
        }
    }
    return esta
}
func EjecutarComando(path string,NombreParticion [15]byte,ListaDiscos *list.List)(bool){
	var encontrada = false
	lineaComando := strings.Split(path, "/")
	nombreDisco:= lineaComando[len(lineaComando)-1]
	f, err := os.OpenFile(path,os.O_RDONLY,0755)
	if err != nil {
		fmt.Println("No existe la ruta"+path)
		return false
	}
	defer f.Close()
	mbr := MBR{}
	f.Seek(0,0)
	err = binary.Read(f, binary.BigEndian, &mbr)
	Particiones := mbr.Particiones
	for i:=0;i<4;i++{
		if string(Particiones[i].NombreParticion[:])== string(NombreParticion[:]){
			encontrada = true
			if strings.ToLower(BytesToString(Particiones[i].TipoParticion)) == "e"{
				fmt.Println("Error no se puede Montar una particion Extendida")
			}else{
				ParticionMontar(ListaDiscos,string(NombreParticion[:]),string(nombreDisco),path)
			}
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
					encontrada = true
					//Montar Particion
					ParticionMontar(ListaDiscos,string(NombreParticion[:]),string(nombreDisco),path)
				}
			}
		}
	}
	if encontrada == false{
		fmt.Println("Error no se encontro la particion")
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
			//#id->vda1
			for i:=0;i<len(disco.Particiones);i++{
				var mountTemp = disco.Particiones[i]
				if BytesToString(mountTemp.Estado) == "0"{
					mountTemp.Id = "vd"+BytesToString(disco.Id)+mountTemp.Id
					mountTemp.NombreParticion = nombreParticion
					copy(mountTemp.Estado[:],"1")
					copy(mountTemp.EstadoMKS[:],"0")
					disco.Particiones[i] = mountTemp
					break
				}else if BytesToString(mountTemp.Estado) == "1" && mountTemp.NombreParticion == nombreParticion{
					fmt.Println("La Particion ya esta montada")
					break
				}
			}
			element.Value = disco
			break
		}else if BytesToString(disco.Estado) == "1" && ExisteDisco(ListaDiscos,nombreDisco) && nombreDisco==disco.NombreDisco{
			fmt.Println("Otra particion montada en el disco ", BytesToString(disco.Id))
			for i:=0;i<len(disco.Particiones);i++{
				var mountTemp = disco.Particiones[i]
				if BytesToString(mountTemp.Estado) == "0"{
					mountTemp.Id = "vd"+BytesToString(disco.Id)+mountTemp.Id
					mountTemp.NombreParticion = nombreParticion
					copy(mountTemp.Estado[:],"1")
					copy(mountTemp.EstadoMKS[:],"0")
					disco.Particiones[i] = mountTemp
					break
				}else if BytesToString(mountTemp.Estado) == "1" && mountTemp.NombreParticion == nombreParticion{
					fmt.Println("La Particion ya esta montada")
					break
				}
			}
			element.Value = disco
			break
		}
	}
}
func ExisteDisco(ListaDiscos *list.List,nombreDisco string)(bool){
	Existe := false
	for element := ListaDiscos.Front(); element != nil; element = element.Next() {
		var disco  DISCO
		disco = element.Value.(DISCO)
		if disco.NombreDisco == nombreDisco{
			return true
		}else{
			Existe = false
		}
	}
	return Existe
}