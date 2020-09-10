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
    var strArray [100]string
	//var InicioParticion int64 =0
	pathDisco,nombreParticion,_:=RecorrerListaDisco(idParticion,ListaDiscos)
	sb,_= DevolverSuperBlque(pathDisco,nombreParticion)
	f, err := os.OpenFile(pathDisco,os.O_RDWR,0755)
	if err != nil {
		fmt.Println("No existe la ruta"+ pathDisco)
		return false
	}
    defer f.Close()
    /*
    Graficar AVD's
    */
    f.Seek(sb.Sb_ap_arbol_directorio,0)
    for i:=0;i<int(sb.Sb_arbol_virtual_count);i++{
    	err = binary.Read(f, binary.BigEndian, &avd)
    	if avd.Avd_nomre_directotrio == dos{
    		break
    	}
    	for j:=0;j<6;j++{
    		if avd.Avd_ap_array_subdirectoios[j]!=-1{
    			buffer.WriteString("nodo" + strconv.Itoa(i) + ":f"+ strconv.Itoa(j) + " -> nodo" + strconv.Itoa(int(avd.Avd_ap_array_subdirectoios[j])) + "\n")
    		}else{
    			break
    		}
    	} 
    	if avd.Avd_ap_arbol_virtual_directorio != -1{
    		buffer.WriteString("nodo" + strconv.Itoa(i) + ":f7" + " -> nodo" + strconv.Itoa(int(avd.Avd_ap_arbol_virtual_directorio)) + "\n")
    	}
    	if EstaLlenoDD(avd.Avd_ap_detalle_directorio,sb.Sb_ap_detalle_directorio,sb.Sb_detalle_directorio_count,pathDisco){
            strArray[i] = CToGoString(avd.Avd_nomre_directotrio[:])
    		buffer.WriteString("nodo"+ strconv.Itoa(i) + ":f6 -> node"+strconv.Itoa(int( avd.Avd_ap_detalle_directorio))+ "\n")
    	}
    	buffer.WriteString("nodo" + strconv.Itoa(i) + "[ shape=record, label =\"" + "{"+ CToGoString(avd.Avd_nomre_directotrio[:])+ "|{<f0> "+strconv.Itoa(int(avd.Avd_ap_array_subdirectoios[0])) +"|<f1>"+strconv.Itoa(int(avd.Avd_ap_array_subdirectoios[1])) + "|<f2> " + strconv.Itoa(int(avd.Avd_ap_array_subdirectoios[2])) + "|<f3> " + strconv.Itoa(int(avd.Avd_ap_array_subdirectoios[3])) + "|<f4> "+ strconv.Itoa(int(avd.Avd_ap_array_subdirectoios[4])) + "|<f5>" +  strconv.Itoa(int(avd.Avd_ap_array_subdirectoios[5])) + "|<f6>" +  strconv.Itoa(int(avd.Avd_ap_detalle_directorio)) + "|<f7> "+  strconv.Itoa(int(avd.Avd_ap_arbol_virtual_directorio)) + "}}\"];\n")
    }
    /*
    Graficar DD's
    */
    f.Seek(sb.Sb_ap_detalle_directorio,0)
    dd := DD{}
    for i:=0;i<int(sb.Sb_detalle_directorio_count);i++{
    	err = binary.Read(f, binary.BigEndian, &dd)
    	if dd.Ocupado == 0{
    		break
    	}
        //fmt.Println(EstaLlenoDD(int64(i),sb.Sb_ap_detalle_directorio,sb.Sb_detalle_directorio_count,pathDisco),i)
    	if EstaLlenoDD(int64(i),sb.Sb_ap_detalle_directorio,sb.Sb_detalle_directorio_count,pathDisco){
            for j:=0;j<5;j++{
                if CToGoString(dd.Dd_array_files[j].Dd_file_nombre[:]) != CToGoString(dos[:]){
                    buffer.WriteString("node"+ strconv.Itoa(i) + ":f"+strconv.Itoa(j+1) +"->  nodex" + strconv.Itoa(int(dd.Dd_array_files[j].Dd_file_ap_inodo))+ "\n")
                }
            }
    		buffer.WriteString("node"+  strconv.Itoa(i) + "[shape=record, label=\"" + "{ dd " + strArray[i] + "|")
    		for j:=0;j<5;j++{
                if CToGoString(dd.Dd_array_files[j].Dd_file_nombre[:]) != CToGoString(dos[:]){
                    buffer.WriteString("{<f" + strconv.Itoa(j) + "> " + CToGoString(dd.Dd_array_files[j].Dd_file_nombre[:]) + "| <f" + strconv.Itoa(j+1) + "> " + strconv.Itoa(int(dd.Dd_array_files[j].Dd_file_ap_inodo)) + "} |")
                }else {
                    buffer.WriteString("{-1 | } |")
                }

			}
            if dd.Dd_ap_detalle_directorio != -1{
                buffer.WriteString("{" + strconv.Itoa(int(dd.Dd_ap_detalle_directorio))+ " | <f10>  }}\"];\n")
                buffer.WriteString("node"+  strconv.Itoa(i)+":f10 -> " + "node"+strconv.Itoa(int(dd.Dd_ap_detalle_directorio)))
            }else{
                buffer.WriteString("{*1 | <f10>  }}\"];\n")
            }
            buffer.WriteString("\n")
    	}
    }
    /*
    Graficar Inodo's
    X para identificarlos
    */
    f.Seek(sb.Sb_ap_tabla_inodo,0)
    inodo := Inodo{}
    for i:=0;i<int(sb.Sb_inodos_count);i++{
        err = binary.Read(f, binary.BigEndian, &inodo)
        if inodo.I_count_inodo == -1{
            break
        }
        if inodo.I_ao_indirecto != -1{
            buffer.WriteString("nodex"+strconv.Itoa(int(inodo.I_count_inodo)) + "[shape=record, label=\"{Inodo"+ strconv.Itoa(int(inodo.I_count_inodo)) +"|{"+ strconv.Itoa(int(inodo.I_array_bloques[0])) +"| <f0> }|{" + strconv.Itoa(int(inodo.I_array_bloques[1])) + "| <f1> }|{" + strconv.Itoa(int(inodo.I_array_bloques[2])) +" | <f2> }|{" + strconv.Itoa(int(inodo.I_array_bloques[3])) +"| <f3> }|{" + strconv.Itoa(int(inodo.I_ao_indirecto)) +" | <f4> }}\"];"+ "\n")
            buffer.WriteString("nodex"+strconv.Itoa(int(inodo.I_count_inodo)) + " :f4 ->" + "nodex"+strconv.Itoa(int(inodo.I_ao_indirecto))+"\n")
            for h:=0;h<4;h++{
                if inodo.I_array_bloques[h]==-1{
                    break
                }else{
                    buffer.WriteString("nodex"+strconv.Itoa(int(inodo.I_count_inodo)) + " :f" + strconv.Itoa(h)  +"-> data"+strconv.Itoa(int(inodo.I_array_bloques[h]))+"\n")
                }
            }
        }else{
            buffer.WriteString("nodex"+strconv.Itoa(int(inodo.I_count_inodo)) + "[shape=record, label=\"{Inodo"+ strconv.Itoa(int(inodo.I_count_inodo)) +"|{"+ strconv.Itoa(int(inodo.I_array_bloques[0])) +"| <f0> }|{" + strconv.Itoa(int(inodo.I_array_bloques[1])) + "| <f1> }|{" + strconv.Itoa(int(inodo.I_array_bloques[2])) +" | <f2> }|{" + strconv.Itoa(int(inodo.I_array_bloques[3])) +"| <f3> }|{*" + strconv.Itoa(int(inodo.I_ao_indirecto)) +" | <f4> }}\"];"+ "\n")
            for h:=0;h<4;h++{
                if inodo.I_array_bloques[h]==-1{
                    break
                }else{
                    buffer.WriteString("nodex"+strconv.Itoa(int(inodo.I_count_inodo)) + " :f" + strconv.Itoa(h)  +"-> data"+strconv.Itoa(int(inodo.I_array_bloques[h]))+"\n")
                }
            }
        }
    }
    /*
    Graficar Bloque's
    */
    f.Seek(sb.Sb_ap_bloques,0)
    data := Bloque{}
    for i:=0;i<int(sb.Sb_bloques_count);i++{
        if data.Db_data[0]==0{
            break
        }
        buffer.WriteString("data" + strconv.Itoa(i) +"[shape=record, label=\"{data| <f1> }}\"];\n")

    }
    buffer.WriteString("\n}")
    var datos string
    datos = string(buffer.String())
    CreateArchivo("/home/edson/Escritorio/Proyecto/Proyecto1/MKFS.dot",datos)
    return false
}
func EstaLlenoDD(posicion int64,inicioDD int64,cantidadDD int64,pathDisco string)(bool){
    estaLleno := false
    f, err := os.OpenFile(pathDisco,os.O_RDWR,0755)
    if err != nil {
        fmt.Println("No existe la ruta"+ pathDisco)
        return false
    }
    defer f.Close()
    f.Seek(inicioDD,0)
    dd := DD{}
    for i:=0;i<int(cantidadDD);i++{
        err = binary.Read(f, binary.BigEndian, &dd)
        if dd.Ocupado == 0{
            break
        }
        if i == int(posicion){
            for j:=0;j<5;j++{
                if len(CToGoString(dd.Dd_array_files[j].Dd_file_nombre[:])) > 0{
                    estaLleno = true
                    break
                }else{
                    estaLleno= false
                }
            }
        }
    }
	return estaLleno
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