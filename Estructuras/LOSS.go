package Estructuras
import (
	"fmt"
	"strings"
	"container/list"
	"os"
	"encoding/binary"
)

func EjecutarComandoLOSS(nombreComando string,propiedadesTemp []Propiedad,ListaDiscos *list.List)(ParamValidos bool){
	fmt.Println("----------------- Ejecutando Loss -----------------")
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
	    ExecuteLoss(propiedades[0],ListaDiscos)
	    return ParamValidos
	}else{
		ParamValidos = false
		return ParamValidos
	}
}
func ExecuteLoss(idParticion string,ListaDiscos *list.List)(bool){
	avd := AVD{}
	dd:= DD{}
	inodo:= Inodo{}
	bloque:= Bloque{}
	sb:=SB{}
	var otro int8 =0
	sbCopia := SB{}
    pathDisco,nombreParticion,_:=RecorrerListaDisco(idParticion,ListaDiscos)
    sb,_= DevolverSuperBlque(pathDisco,nombreParticion)
    _,_,InicioParticion:= ReturnMBR(pathDisco,nombreParticion)
    f, err := os.OpenFile(pathDisco,os.O_RDWR,0755)
	if err != nil {
		fmt.Println("No existe la ruta"+pathDisco)
		return false
	}
	defer f.Close()
	f.Seek(sb.InicioCopiaSB,0)
    err = binary.Read(f, binary.BigEndian, &sbCopia)
    fmt.Println(sbCopia.Sb_magic_num)

    f.Seek(InicioParticion,0)
    err = binary.Write(f, binary.BigEndian, &sbCopia)
    //Escribir Bit Map Arbol Virtual de Directorio
    f.Seek(sbCopia.Sb_ap_bitmap_arbol_directorio,0)
    i := 0
    for i=0;i<int(sbCopia.Sb_arbol_virtual_count);i++{
    	err = binary.Write(f, binary.BigEndian, &otro)
    }
    //Escribir Arbol de Directorio
    f.Seek(sbCopia.Sb_ap_arbol_directorio,0)
    i=0
    for i=0;i<int(sbCopia.Sb_arbol_virtual_count);i++{
    	err = binary.Write(f, binary.BigEndian, &avd)
    }
    //Escribir Bitmap Detalle Directorio
    f.Seek(sbCopia.Sb_ap_bitmap_detalle_directorio,0)
    i=0
    for i=0;i<int(sbCopia.Sb_detalle_directorio_count);i++{
    	err = binary.Write(f, binary.BigEndian, &otro)
    }
    //Escribir Detalle Directorio
    f.Seek(sbCopia.Sb_ap_detalle_directorio,0)
    i=0
    dd.Dd_ap_detalle_directorio = -1
    for i=0;i<int(sbCopia.Sb_detalle_directorio_count);i++{
    	err = binary.Write(f, binary.BigEndian, &dd)
    }
    //Escribir Bitmap Tabla Inodo
    f.Seek(sbCopia.Sb_ap_bitmap_tabla_inodo,0)
    i=0
    for i=0;i<int(sbCopia.Sb_inodos_count);i++{
    	err = binary.Write(f, binary.BigEndian, &otro)
    }
    //Escribir Tabla Inodos
	f.Seek(sbCopia.Sb_ap_tabla_inodo,0)
	    i=0
	    inodo.I_count_inodo = -1
	    for i=0;i<int(sbCopia.Sb_inodos_count);i++{
	    	err = binary.Write(f, binary.BigEndian, &inodo)
	    }
    //Escribir Bitmap BLoque de datos
    f.Seek(sbCopia.Sb_ap_bitmap_bloques,0)
    i=0
    for i=0;i<int(sbCopia.Sb_bloques_count);i++{
    	err = binary.Write(f, binary.BigEndian, &otro)
    }
    //Escribir Bloque de datos
    f.Seek(sbCopia.Sb_ap_bloques,0)
	    i=0
	    copy(bloque.Db_data[:],"")
	    for i=0;i<int(sbCopia.Sb_bloques_count);i++{
	    	err = binary.Write(f, binary.BigEndian, &bloque)
	    }
    //Escribir Bitacoras
    /*f.Seek(InicioBitacora,0)
    i=0
    bitacora.Size = -1
    for i=0;i<Bitacoras;i++{
    	err = binary.Write(f, binary.BigEndian, &bitacora)
    }
    //Escribir Copia Super Boot
    f.Seek(InicioCopiaSB,0)
    err = binary.Write(f, binary.BigEndian, &superBloque)

    /*f.Seek(InicioCopiaSB,0)
    err = binary.Read(f, binary.BigEndian, &superBloque)
    fmt.Println("----------")
    fmt.Println(superBloque)
    fmt.Println("----------")*/
    //Crear Raiz  -----> /  y  archivo con usuarios
    CrearRaiz(pathDisco,InicioParticion)
	fmt.Println("Particion a perder",nombreParticion)
    return false
}