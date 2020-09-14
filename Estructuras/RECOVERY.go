package Estructuras
import (
	"fmt"
	"strings"
	"container/list"
	"os"
	"encoding/binary"
)

func EjecutarComandoRECOVERY(nombreComando string,propiedadesTemp []Propiedad,ListaDiscos *list.List)(ParamValidos bool){
	fmt.Println("----------------- Ejecutando RECOVERY -----------------")
	ParamValidos = true
	var propiedades [1]string
	if len(propiedadesTemp) >= 1{
	    //Recorrer la lista de propiedades
	    for i:=0; i < len(propiedadesTemp); i++{
	        var propiedadTemp = propiedadesTemp[i]
	        var nombrePropiedad string= propiedadTemp.Name
	        switch strings.ToLower(nombrePropiedad){
	        case "-id":
	        	propiedades [0]=propiedadTemp.Val
	    	default:
	    		fmt.Println("Error al Ejecutar el Comando")
	        }
	    }
	    ExecuteRecovery(propiedades[0],ListaDiscos)
	    return ParamValidos
	}else{
		ParamValidos = false
		return ParamValidos
	}
}
func ExecuteRecovery(idParticion string,ListaDiscos *list.List)(bool){
	sb:=SB{}
	var array [100]Bitacora
    pathDisco,nombreParticion,_:=RecorrerListaDisco(idParticion,ListaDiscos)
    sb,_= DevolverSuperBlque(pathDisco,nombreParticion)
    bitacora := Bitacora{}
    f, err := os.OpenFile(pathDisco,os.O_RDWR,0755)
	if err != nil {
		fmt.Println("No existe la ruta"+pathDisco)
		return false
	}
	defer f.Close()
	conteo:=0
	f.Seek(sb.Sb_ap_log,0)
	fmt.Println(sb.Sb_ap_log)
	for i:=0;i<int(sb.Sb_inodos_count);i++{
		err = binary.Read(f, binary.BigEndian, &bitacora)
		if bitacora.Size==-1{
			break
		}
		if convertName(bitacora.Log_tipo[:])=="0"{
			array[conteo] = bitacora
			conteo++ 
			//ExecuteMKDIR(idParticion,convertName(bitacora.Log_nombre[:]),"-p",ListaDiscos)
		}else if convertName(bitacora.Log_tipo[:])=="1"{
			//ExecuteMKFILE(idParticion,convertName(bitacora.Log_nombre[:]),"-p",int(bitacora.Size),convertBloqueData(bitacora.Log_Contenido[:]),ListaDiscos)
			array[conteo] = bitacora
			conteo++ 
		}
	}
	f.Seek(0,0)
	for i:=0;i<conteo;i++{
		if convertName(array[i].Log_tipo[:])=="0"{
			ExecuteMKDIR(idParticion,convertName(array[i].Log_nombre[:]),"-p",ListaDiscos)
		}else if convertName(array[i].Log_tipo[:])=="1"{
			ExecuteMKFILE(idParticion,convertName(array[i].Log_nombre[:]),"-p",int(array[i].Size),convertBloqueData(array[i].Log_Contenido[:]),ListaDiscos)
		}
	}
	f.Seek(sb.Sb_ap_log,0)
	bitacora2 :=Bitacora{}
	bitacora2.Size = -1
    for i:=0;i<int(sb.Sb_arbol_virtual_count);i++{
    	err = binary.Write(f, binary.BigEndian, &bitacora2)
    }
	return false
}