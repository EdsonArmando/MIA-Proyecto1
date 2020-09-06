package Estructuras
import (
	"fmt"
	"os"
	"container/list"
	"strings"
	"encoding/binary"
)

func EjecutarComandoLogin(nombreComando string,propiedadesTemp []Propiedad,ListaDiscos *list.List)(bool,string){
	fmt.Println("----------------- Ejecutando Login -----------------")
	ParamValidos := true
	usuario :=""
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
	    ParamValidos,usuario = ExecuteLogin(propiedades[0],propiedades[1],propiedades[2],ListaDiscos)
	    return ParamValidos,usuario
	}else{
		ParamValidos = false
		return ParamValidos,usuario
	}
}
func ExecuteLogin(usuario string,password string,id string,ListaDiscos *list.List)(bool,string){
	idValido := IdValido(id,ListaDiscos);
	if idValido == false{
		fmt.Println("El id no existe la particion no esta montada")
		return false,""
	}else if global!=""{
		fmt.Println("Ya hay una sesion iniciada")
		return false,""
	}
	pathDisco,nombreParticion,nombreDisco:=RecorrerListaDisco(id,ListaDiscos)
	mbr,sizeParticion,InicioParticion:= ReturnMBR(pathDisco,nombreParticion)
	superBloque := SB{}
	f, err := os.OpenFile(pathDisco,os.O_RDONLY,0755)
	if err != nil {
		fmt.Println("No existe la ruta"+pathDisco)
		return false,""
	}
	defer f.Close()
    f.Seek(InicioParticion,0)
    err = binary.Read(f, binary.BigEndian, &superBloque)
    /*
    Obtener avd raiz
    */
    avd := AVD{}
    dd := DD{}
    inodo := Inodo{}
    bloque := Bloque{}
    f.Seek(superBloque.Sb_ap_arbol_directorio,0)
    err = binary.Read(f, binary.BigEndian, &avd)
    apuntadorDD := avd.Avd_ap_detalle_directorio
    f.Seek(superBloque.Sb_ap_detalle_directorio,0)
    for i:=0;i<int(superBloque.Sb_arbol_virtual_free);i++{
    	err = binary.Read(f, binary.BigEndian, &dd)
    	if i == int(apuntadorDD){
    		break
    	}
    }
    apuntadorInodo:=dd.Dd_array_files[0].Dd_file_ap_inodo
    f.Seek(superBloque.Sb_ap_tabla_inodo,0)
     for i:=0;i<int(superBloque.Sb_inodos_free);i++{
    	err = binary.Read(f, binary.BigEndian, &inodo)
    	if i == int(apuntadorInodo){
    		break
    	}
    }
	var userstxt string = ""

	//Leer Users.txt
	posicion:=0
	f.Seek(superBloque.Sb_ap_bloques,0)
	for i:=0;i<int(superBloque.Sb_inodos_free);i++{
    	err = binary.Read(f, binary.BigEndian, &bloque)

    	if int(inodo.I_array_bloques[posicion])!=-1 && int(inodo.I_array_bloques[posicion])==i{
    		userstxt += ConvertData(bloque.Db_data)
    	}else if int(inodo.I_array_bloques[posicion])==-1{
    		break
    	}else{
    		break
    	}
    	if posicion <4{
    		posicion++
    	}else if posicion==4{
    		posicion=0
    	}
    }
    lineaUsuarioTxt := strings.Split(userstxt, "\n")
    for i:=0;i<len(lineaUsuarioTxt);i++{
    	if len(lineaUsuarioTxt[i])!=17{
    		usuario_grupo := strings.Split(lineaUsuarioTxt[i], ",")
    		if usuario_grupo[1] == "U"{
    			if usuario_grupo[3]==usuario && usuario_grupo[4]==password{
    				fmt.Println("Bienvenido al sistema")
    				return true,usuario
    			}
    		}
    	}
    }
    fmt.Println(nombreDisco,mbr.MbrTamanio,sizeParticion)
	return false,""
}
func RecorrerListaDisco(id string,ListaDiscos *list.List)(string,string,string){
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
					return pathDisco,nombreParticion,nombreDisco
					break
				}
			}

		}
		element.Value = disco
	}
	fmt.Println(NoParticion)
	return "","",""
}
