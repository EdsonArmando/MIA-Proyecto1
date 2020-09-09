package Estructuras

import (
	"fmt"
	"encoding/binary"
	"container/list"
	"os"
	"strings"
	"strconv"
	"log"
	"bytes"
)
func EjecutarComandoReporte(nombreComando string,propiedadesTemp []Propiedad,ListaDiscos *list.List)(ParamValidos bool){
	fmt.Println("----------------- Ejecutando Reporte -----------------")
	ParamValidos = true
	var propiedades [3]string
	if len(propiedadesTemp) >= 1{
	    //Recorrer la lista de propiedades
	    for i:=0; i < len(propiedadesTemp); i++{
	        var propiedadTemp = propiedadesTemp[i]
	        var nombrePropiedad string= propiedadTemp.Name
	        switch strings.ToLower(nombrePropiedad){
	        case "-id":
	        	propiedades [0]=propiedadTemp.Val
        	case "-path":
        		propiedades [1]=propiedadTemp.Val
    		case "-dest":
    			propiedades [1]=propiedadTemp.Val
	    	default:
	    		fmt.Println("Error al Ejecutar el Comando")
	        }
	    }
	    Reporte(propiedades[0],ListaDiscos)
	    return ParamValidos
	}else{
		ParamValidos = false
		return ParamValidos
	}
}

func Reporte(idParticion string,ListaDiscos *list.List)(bool){
	GraficarMKFS(idParticion,ListaDiscos)
	/*sb:=SB{}
	avd := AVD{}
	conteo:=0
	//var InicioParticion int64 =0
	pathDisco,nombreParticion,_:=RecorrerListaDisco(idParticion,ListaDiscos)
	sb,_= DevolverSuperBlque(pathDisco,nombreParticion)
	f, err := os.OpenFile(pathDisco,os.O_RDWR,0755)
	if err != nil {
		fmt.Println("No existe la ruta"+ pathDisco)
		return false
	}
    defer f.Close()
    var otro int8 = 0
    f.Seek(sb.Sb_ap_arbol_directorio,0)
    for i:=0;i<10;i++{
    	err = binary.Read(f, binary.BigEndian, &avd)
    	fmt.Println("Puntero DD",avd.Avd_ap_detalle_directorio)
    }
    fmt.Println("---------------Bitmap AVD------------------------")
    dd:=DD{}
	f.Seek(sb.Sb_ap_detalle_directorio,0)
    for i:=0;i<20;i++{
    	err = binary.Read(f, binary.BigEndian, &dd)
    	if i == 8 || i == 9{
    		for j:=0;j<5;j++{
    			fmt.Println(BytesNombreParticion(dd.Dd_array_files[j].Dd_file_nombre),":",dd.Dd_array_files[j].Dd_file_ap_inodo,dd.Dd_ap_detalle_directorio)
    		}
    	}
    }
    fmt.Print(conteo)
    conteo=0
    fmt.Println("---------------Bitmap DD------------------------")
	f.Seek(sb.Sb_ap_bitmap_tabla_inodo,0)
    for i:=0;i<20;i++{
    	err = binary.Read(f, binary.BigEndian, &otro)
    	fmt.Print(otro)
    	if otro==1{
    		conteo++
    	}
    }
    fmt.Print(conteo)
    conteo=0
    fmt.Println("---------------Bitmap Inodo------------------------")
	f.Seek(sb.Sb_ap_tabla_inodo,0)
	inodo :=Inodo{}
    for i:=0;i<19;i++{
    	err = binary.Read(f, binary.BigEndian, &inodo)
    	fmt.Println(inodo)
    }
    fmt.Println("---------------Bitmap Bloque------------------------",sb.ConteoBloque)
	*/
    return false
}

func GraficarMKFS(idParticion string,ListaDiscos *list.List)(bool){
	var buffer bytes.Buffer
	buffer.WriteString("digraph grafica{\nrankdir=TB;\nnode [shape = record, style=filled, fillcolor=seashell2];\n")
	sb:=SB{}
	var dos [15]byte
	avd := AVD{}
	//var InicioParticion int64 =0
	pathDisco,nombreParticion,_:=RecorrerListaDisco(idParticion,ListaDiscos)
	sb,_= DevolverSuperBlque(pathDisco,nombreParticion)
	f, err := os.OpenFile(pathDisco,os.O_RDWR,0755)
	if err != nil {
		fmt.Println("No existe la ruta"+ pathDisco)
		return false
	}
    defer f.Close()

    f.Seek(sb.Sb_ap_arbol_directorio,0)
    for i:=0;i<int(sb.Sb_arbol_virtual_count);i++{
    	err = binary.Read(f, binary.BigEndian, &avd)
    	if avd.Avd_nomre_directotrio == dos{
    		break
    	}
    	fmt.Println(BytesNombreParticion(avd.Avd_nomre_directotrio))
    	buffer.WriteString("nodo" + strconv.Itoa(i) + "[ shape=record, label =\"" + "{"+ CToGoString(avd.Avd_nomre_directotrio[:])+ "|{<f0> "+strconv.Itoa(int(avd.Avd_ap_array_subdirectoios[0])) +"|<f1>"+strconv.Itoa(int(avd.Avd_ap_array_subdirectoios[1])) + "|<f2> " + strconv.Itoa(int(avd.Avd_ap_array_subdirectoios[2])) + "|<f3> " + strconv.Itoa(int(avd.Avd_ap_array_subdirectoios[3])) + "|<f4> "+ strconv.Itoa(int(avd.Avd_ap_array_subdirectoios[4])) + "|<f5>" +  strconv.Itoa(int(avd.Avd_ap_array_subdirectoios[5])) + "|<f6>" +  strconv.Itoa(int(avd.Avd_ap_detalle_directorio)) + "|<f7> "+  strconv.Itoa(int(avd.Avd_ap_arbol_virtual_directorio)) + "}}\"];\n")

    }
    buffer.WriteString("\n}")
    var datos string
    datos = string(buffer.String())
    CreateArchivo("/home/edson/Escritorio/Proyecto/Proyecto1/MKFS.dot",datos)
    return false
}

func CreateArchivo(path string,data string){
	f, err := os.Create("/home/edson/Escritorio/Proyecto/Proyecto1/MKFS.dot")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    _, err2 := f.WriteString(data)

    if err2 != nil {
        log.Fatal(err2)
    }
    executeComand("dot -Tpdf MKFS.dot -o outfile.pdf")
}	

func CToGoString(c []byte) string {
    n := -1
    for i, b := range c {
        if b == 0 {
            break
        }
        n = i
    }
    return string(c[:n+1])
}