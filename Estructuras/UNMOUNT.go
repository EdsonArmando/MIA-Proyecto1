package Estructuras
import (
	"fmt"
	"strings"
	"container/list"
)
func EjecutarComandoUnmount(nombreComando string,propiedadesTemp []Propiedad,ListaDiscos *list.List)(ParamValidos bool){
	fmt.Println("----------------- Ejecutando UNMOUNT -----------------")
	ParamValidos = true
	if len(propiedadesTemp) >= 1{
	    //Recorrer la lista de propiedades
	    for i:=0; i < len(propiedadesTemp); i++{
	        var propiedadTemp = propiedadesTemp[i]
	        var nombrePropiedad string= propiedadTemp.Name
	       	Id := strings.ReplaceAll(nombrePropiedad, "-id", "")
	        switch strings.ToLower(nombrePropiedad){
	        case "-id"+Id:
	        	ejecutarUnmount(propiedadTemp.Val,ListaDiscos)
	    	default:
	    		fmt.Println("Error al Ejecutar el Comando")
	        }
	    }
	    return ParamValidos
	}else{
		ParamValidos = false
		return ParamValidos
	}
}

func ejecutarUnmount(id string,ListaDiscos *list.List)(bool){
	Id:= strings.ReplaceAll(id, "vd","")
	eliminada := false
	NoParticion := Id[1:]
	IdDisco := Id[:1]
	for element := ListaDiscos.Front(); element != nil; element = element.Next() {
		var disco  DISCO
		disco = element.Value.(DISCO)
		if BytesToString(disco.Id) == IdDisco{
		   for i:=0;i<len(disco.Particiones);i++{
				var mountTemp = disco.Particiones[i]
				if mountTemp.Id == id{
					eliminada = true
					fmt.Println("Encontrada",NoParticion)
					mountTemp.Id = NoParticion
					mountTemp.NombreParticion = ""
					copy(mountTemp.Estado[:],"0")
					copy(mountTemp.EstadoMKS[:],"0")
					disco.Particiones[i] = mountTemp
					fmt.Println("Particion Desmontada ", id)
					break
				}
			}

		}
		element.Value = disco
	}
	if eliminada == false{
		fmt.Println("No se pudo Desmontar la particion con id",id)
	}
	return false
}

//unmount -id5->vda1