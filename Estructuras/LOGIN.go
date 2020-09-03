package Estructuras
import (
	"fmt"
	"os"
	"container/list"
	"strings"
	"encoding/binary"
)

func EjecutarComandoLogin(nombreComando string,propiedadesTemp []Propiedad,ListaDiscos *list.List)(ParamValidos bool){
	fmt.Println("----------------- Ejecutando Login -----------------")
	ParamValidos = true
	var propiedades [3]string
	if len(propiedadesTemp) >= 1{
	    //Recorrer la lista de propiedades
	    for i:=0; i < len(propiedadesTemp); i++{
	        var propiedadTemp = propiedadesTemp[i]
	        var nombrePropiedad string= propiedadTemp.Name
	        switch strings.ToLower(nombrePropiedad){
	        case "-usr":
	        	propiedades [0]=propiedadTemp.Val
        	case "-pwd":
        		propiedades [1]=string(propiedadTemp.Val)
    		case "-id":
    			propiedades [2]=propiedadTemp.Val
	    	default:
	    		fmt.Println("Error al Ejecutar el Comando")
	        }
	    }
	    ExecuteLogin(propiedades[0],propiedades[1],propiedades[2],ListaDiscos)
	    return ParamValidos
	}else{
		ParamValidos = false
		return ParamValidos
	}
}
func ExecuteLogin(usuario string,password string,id string,ListaDiscos *list.List)(bool){
	idValido := IdValido(id,ListaDiscos);
	if idValido == false{
		fmt.Println("El id no existe la particion no esta montada")
		return false
	}
	fmt.Println("IdValido")
	Id:= strings.ReplaceAll(id, "vd","")
	NoParticion := Id[1:]
	IdDisco := Id[:1]
	pathDisco := ""
	nombreParticion:=""
	nombreDisco:=""
	for element := ListaDiscos.Front(); element != nil; element = element.Next() {
		var disco  DISCO
		disco = element.Value.(DISCO)
		if BytesToString(disco.Id) == IdDisco{
		   for i:=0;i<len(disco.Particiones);i++{
				var mountTemp = disco.Particiones[i]
				if mountTemp.Id == id{
					copy(mountTemp.EstadoMKS[:],"1");
					nombreParticion = mountTemp.NombreParticion
					pathDisco = disco.Path
					nombreDisco = disco.NombreDisco
					break
				}
			}

		}
		element.Value = disco
	}
	mbr,sizeParticion,InicioParticion:= ReturnMBR(pathDisco,nombreParticion)
	bloque := Bloque{}
	superBloque := SB{}
	f, err := os.OpenFile(pathDisco,os.O_RDONLY,0755)
	if err != nil {
		fmt.Println("No existe la ruta"+pathDisco)
		return false
	}
	defer f.Close()
    f.Seek(InicioParticion,0)
    err = binary.Read(f, binary.BigEndian, &superBloque)
    f.Seek(superBloque.Sb_ap_bloques,0)
	err = binary.Read(f, binary.BigEndian, &bloque)
	fmt.Printf("DataBloque: %s\n",bloque.Db_data)
    fmt.Println("------------")
    fmt.Println(NoParticion,nombreDisco,sizeParticion,mbr.MbrTamanio)
	return false
}