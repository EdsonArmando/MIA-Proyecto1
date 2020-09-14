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
	var propiedades [4]string
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
    		case "-nombre":
    			propiedades [2]=propiedadTemp.Val
            case "-ruta":
                propiedades [3]=propiedadTemp.Val
            case "-sigue":
                propiedades [1]+=propiedadTemp.Val
	    	default:
	    		fmt.Println("Error al Ejecutar el Comando",nombrePropiedad)
	        }
	    }
        EsComilla :=  propiedades[1][0:1]
        if EsComilla == "\""{
            if propiedades[3]!=""{
                propiedades[3] = propiedades[3][1 : len(propiedades[3])-1]
            }
            propiedades[1] = propiedades[1][1 : len(propiedades[1])-1]
        }
        carpetas_Graficar := strings.Split(propiedades[1], "/")
        var comando = ""
        for i:=1;i<len(carpetas_Graficar)-1;i++{
            comando+=carpetas_Graficar[i]+"/"
        }
        executeComand("mkdir " + comando[0:len(comando)-1])
        switch strings.ToLower(propiedades[2]){
        case "mbr":
            GraficarMBR_EBR(propiedades[0],ListaDiscos,propiedades[1])
        case "disk":
            GraficarDisco(propiedades[0],ListaDiscos,propiedades[1])
        case "sb":
            GraficarSuperBloque(propiedades[0],ListaDiscos,propiedades[1])
        case "bm_arbdir":
            Reporteb_m_arbdir(propiedades[0],propiedades[1],propiedades[3],ListaDiscos)
        case "bm_detdir":
            Reporteb_m_detdir(propiedades[0],propiedades[1],propiedades[3],ListaDiscos)
        case "bm_inode":
            Reporte_bm_inode(propiedades[0],propiedades[1],propiedades[3],ListaDiscos)
        case "bm_block":
            Reporte_bm_block(propiedades[0],propiedades[1],propiedades[3],ListaDiscos)
        case "bitacora":
            GraficarBitacora(propiedades[0],propiedades[1],propiedades[3],ListaDiscos)
        case "directorio":
            Reporte_directorio(propiedades[0],propiedades[1],propiedades[3],ListaDiscos)
        case "tree_file":
            Reporte_tree_file(propiedades[0],propiedades[1],propiedades[3],ListaDiscos)
        case "tree_directorio":
            Reporte_tree_directorio(propiedades[0],propiedades[1],propiedades[3],ListaDiscos)
        case "tree_complete":
            GraficarTreeFull(propiedades[0],propiedades[1],propiedades[3],ListaDiscos)
        default:
            fmt.Println("Nombre Incorrecto")
               
        }
	    return ParamValidos
	}else{
		ParamValidos = false
		return ParamValidos
	}
}
/*
Reporte Bitacora
*/
func GraficarBitacora(idParticion string,pathCarpeta string,ruta string,ListaDiscos *list.List)(bool){
    sb:=SB{}
    var buffer bytes.Buffer
    buffer.WriteString("digraph G{\nrankdir=TB\n")
    pathDisco,nombreParticion,_:=RecorrerListaDisco(idParticion,ListaDiscos)
    sb,_= DevolverSuperBlque(pathDisco,nombreParticion)
    f, err := os.OpenFile(pathDisco,os.O_RDWR,0755)
    if err != nil {
        fmt.Println("No existe la ruta"+ pathDisco)
        return false
    }
    f.Seek(sb.Sb_ap_log,0)
    bitacora := Bitacora{}
    for i:=0;i<3000;i++{
        err = binary.Read(f, binary.BigEndian, &bitacora)
        if bitacora.Size==-1{
            break
        }
        buffer.WriteString("tbl"+strconv.Itoa(i+1)+" [\nshape=box\nlabel=<\n")
        buffer.WriteString("<table border='5'  height='30' WIDTH='15.0' cellborder='1'>\n")
        buffer.WriteString("<tr>  <td height='60' colspan='6'>Log: "+ strconv.Itoa(i+1) +"</td>  </tr>\n")
        buffer.WriteString("<tr>  <td width='150'> <b>TipoOperacion</b> </td> <td width='150'> <b>Tipo</b> </td> <td width='150'> <b>Path</b> </td> <td width='150'> <b>Contenido</b> </td>  <td width='150'> <b>Fecha Log</b> </td> <td width='150'> <b>Size</b> </td></tr>\n")
        buffer.WriteString("<tr>  <td width='150'> <b>"+convertName(bitacora.Log_tipo_operacion[:])+"</b> </td> <td width='150'> <b>"+convertName(bitacora.Log_tipo[:])+"</b> </td> <td width='150'> <b>"+convertName(bitacora.Log_nombre[:])+"</b> </td> <td width='150'> <b> "+convertName(bitacora.Log_Contenido[:])+" </b> </td>  <td width='150'> <b>"+convertName(bitacora.Log_fecha[:])+"</b> </td> <td width='150'> <b>"+strconv.Itoa(int(bitacora.Size))+"</b> </td></tr>\n")
        buffer.WriteString("</table>\n")
        buffer.WriteString(">];\n")
        if BuscarBitacora(idParticion,pathCarpeta,ruta,ListaDiscos,i+1){
            buffer.WriteString("tbl" + strconv.Itoa(i+1) + "-> tbl"+ strconv.Itoa(i+2) + " [color=white];\n")
        }
    }
    buffer.WriteString("}")
    var datos string
    datos = string(buffer.String())
    CreateArchivo(pathCarpeta,datos)
    return false
}
func BuscarBitacora(idParticion string,pathCarpeta string,ruta string,ListaDiscos *list.List,posicion int)(bool){
    sb:=SB{}
    pathDisco,nombreParticion,_:=RecorrerListaDisco(idParticion,ListaDiscos)
    sb,_= DevolverSuperBlque(pathDisco,nombreParticion)
    f, err := os.OpenFile(pathDisco,os.O_RDWR,0755)
    if err != nil {
        fmt.Println("No existe la ruta"+ pathDisco)
        return false
    }
    f.Seek(sb.Sb_ap_log,0)
    bitacora := Bitacora{}
    for i:=0;i<3000;i++{
        err = binary.Read(f, binary.BigEndian, &bitacora)
        if bitacora.Size==-1 && posicion==i{
            return false
        }else if bitacora.Size!=-1 && posicion==i{
            return true
        }
    }
    return false
}
/*
    Reporte del bitmap Arbol de directorio
*/
func Reporteb_m_arbdir(idParticion string,pathCarpeta string,ruta string,ListaDiscos *list.List)(bool){
    sb:=SB{}
    var buffer bytes.Buffer
    pathDisco,nombreParticion,_:=RecorrerListaDisco(idParticion,ListaDiscos)
    sb,_= DevolverSuperBlque(pathDisco,nombreParticion)
    f, err := os.OpenFile(pathDisco,os.O_RDWR,0755)
    if err != nil {
        fmt.Println("No existe la ruta"+ pathDisco)
        return false
    }
    cont:=0
    var otro int8 = 0
    f.Seek(sb.Sb_ap_bitmap_arbol_directorio,0)
    for i:=0;i<int(sb.Sb_arbol_virtual_count);i++{
        err = binary.Read(f, binary.BigEndian, &otro)
        buffer.WriteString(strconv.Itoa(int(otro)) + "|")
        cont++
        if i == 600{
            break
        }else if cont == 20{
            buffer.WriteString("\n")
            cont=0
        }
    }
    buffer.WriteString("\n}")
    var datos string
    datos = string(buffer.String())
    CreateArchivo(pathCarpeta,datos)
    return false
}
/*
    Reporte BItmap detalle de directorio
*/
func Reporteb_m_detdir(idParticion string,pathCarpeta string,ruta string,ListaDiscos *list.List)(bool){
    sb:=SB{}
    var buffer bytes.Buffer
    //var InicioParticion int64 =0
    pathDisco,nombreParticion,_:=RecorrerListaDisco(idParticion,ListaDiscos)
    sb,_= DevolverSuperBlque(pathDisco,nombreParticion)
    f, err := os.OpenFile(pathDisco,os.O_RDWR,0755)
    if err != nil {
        fmt.Println("No existe la ruta"+ pathDisco)
        return false
    }
    cont:=0
    var otro int8 = 0
    f.Seek(sb.Sb_ap_bitmap_detalle_directorio,0)
    for i:=0;i<int(sb.Sb_detalle_directorio_count);i++{
        err = binary.Read(f, binary.BigEndian, &otro)
        buffer.WriteString(strconv.Itoa(int(otro)) + "|")
        cont++
        if i == 600{
            break
        }else if cont == 20{
            buffer.WriteString("\n")
            cont=0
        }
    }
    buffer.WriteString("\n}")
    var datos string
    datos = string(buffer.String())
    CreateArchivo(pathCarpeta,datos)
    return false
}
/*
    Reporte Bitmap de Inodo
*/
func Reporte_bm_inode(idParticion string,pathCarpeta string,ruta string,ListaDiscos *list.List)(bool){
    sb:=SB{}
    var buffer bytes.Buffer
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
    cont:=0
    f.Seek(sb.Sb_ap_bitmap_tabla_inodo,0)
    for i:=0;i<int(sb.Sb_inodos_count);i++{
        err = binary.Read(f, binary.BigEndian, &otro)
        buffer.WriteString(strconv.Itoa(int(otro)) + "|")
        cont++
        if i == 600{
            break
        }else if cont == 20{
            buffer.WriteString("\n")
            cont=0
        }
    }
    buffer.WriteString("\n}")
    var datos string
    datos = string(buffer.String())
    CreateArchivo(pathCarpeta,datos)
    return false
}
/*
    Reporte Bitmap Bloque
*/
func Reporte_bm_block(idParticion string,pathCarpeta string,ruta string,ListaDiscos *list.List)(bool){
    sb:=SB{}
    var buffer bytes.Buffer
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
    cont:=0
    f.Seek(sb.Sb_ap_bitmap_bloques,0)
    for i:=0;i<int(sb.Sb_bloques_count);i++{
        err = binary.Read(f, binary.BigEndian, &otro)
        buffer.WriteString(strconv.Itoa(int(otro)) + "|")
        cont++
        if i == 600{
            break
        }else if cont == 20{
            buffer.WriteString("\n")
            cont=0
        }
    }
    buffer.WriteString("\n}")
    var datos string
    datos = string(buffer.String())
    CreateArchivo(pathCarpeta,datos)
    return false
}
/*
    Reporte Directorio
*/
func Reporte_directorio(idParticion string,pathCarpeta string,ruta string,ListaDiscos *list.List)(bool){
    var buffer bytes.Buffer
    buffer.WriteString("digraph grafica{\nrankdir=TB;\nnode [shape = record, style=filled, fillcolor=seashell2];\n")
    sb:=SB{}
    var dos [15]byte
    avd := AVD{}
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
        buffer.WriteString("nodo" + strconv.Itoa(i) + "[ shape=record, label =\"" + "{"+ convertName(avd.Avd_nomre_directotrio[:])+ "|{<f0> "+strconv.Itoa(int(avd.Avd_ap_array_subdirectoios[0])) +"|<f1>"+strconv.Itoa(int(avd.Avd_ap_array_subdirectoios[1])) + "|<f2> " + strconv.Itoa(int(avd.Avd_ap_array_subdirectoios[2])) + "|<f3> " + strconv.Itoa(int(avd.Avd_ap_array_subdirectoios[3])) + "|<f4> "+ strconv.Itoa(int(avd.Avd_ap_array_subdirectoios[4])) + "|<f5>" +  strconv.Itoa(int(avd.Avd_ap_array_subdirectoios[5])) + "|<f6>" +  strconv.Itoa(int(avd.Avd_ap_detalle_directorio)) + "|<f7> "+  strconv.Itoa(int(avd.Avd_ap_arbol_virtual_directorio)) + "}}\"];\n")
    }
    buffer.WriteString("\n}")
    var datos string
    datos = string(buffer.String())
    CreateArchivo(pathCarpeta,datos)
    return false
}
func Reporte_tree_directorio(idParticion string,pathCarpeta string,ruta string,ListaDiscos *list.List)(bool){
    carpetas_Graficar := strings.Split(ruta, "/")
    var buffer bytes.Buffer
    var noDirectorio int64 =0
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
    f.Seek(sb.Sb_ap_arbol_directorio,0)
    for i:=0;i<int(sb.Sb_arbol_virtual_count);i++{
        err = binary.Read(f, binary.BigEndian, &avd)
        if avd.Avd_nomre_directotrio == dos{
            break
        }
        if convertName(avd.Avd_nomre_directotrio[:]) == carpetas_Graficar[len(carpetas_Graficar)-1]{
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
            noDirectorio = avd.Avd_ap_detalle_directorio
            if EstaLlenoDD(avd.Avd_ap_detalle_directorio,sb.Sb_ap_detalle_directorio,sb.Sb_detalle_directorio_count,pathDisco){
                strArray[i] = convertName(avd.Avd_nomre_directotrio[:])
                buffer.WriteString("nodo"+ strconv.Itoa(i) + ":f6 -> node"+strconv.Itoa(int( avd.Avd_ap_detalle_directorio))+ "\n")
            }
            buffer.WriteString("nodo" + strconv.Itoa(i) + "[ shape=record, label =\"" + "{"+ convertName(avd.Avd_nomre_directotrio[:])+ "|{<f0> "+strconv.Itoa(int(avd.Avd_ap_array_subdirectoios[0])) +"|<f1>"+strconv.Itoa(int(avd.Avd_ap_array_subdirectoios[1])) + "|<f2> " + strconv.Itoa(int(avd.Avd_ap_array_subdirectoios[2])) + "|<f3> " + strconv.Itoa(int(avd.Avd_ap_array_subdirectoios[3])) + "|<f4> "+ strconv.Itoa(int(avd.Avd_ap_array_subdirectoios[4])) + "|<f5>" +  strconv.Itoa(int(avd.Avd_ap_array_subdirectoios[5])) + "|<f6>" +  strconv.Itoa(int(avd.Avd_ap_detalle_directorio)) + "|<f7> "+  strconv.Itoa(int(avd.Avd_ap_arbol_virtual_directorio)) + "}}\"];\n")
        }
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
        if noDirectorio == int64(i){
            if EstaLlenoDD(int64(i),sb.Sb_ap_detalle_directorio,sb.Sb_detalle_directorio_count,pathDisco){
                /*for j:=0;j<5;j++{
                    if convertName(dd.Dd_array_files[j].Dd_file_nombre[:]) != convertName(dos[:]){
                        buffer.WriteString("node"+ strconv.Itoa(i) + ":f"+strconv.Itoa(j+1) +"->  nodex" + strconv.Itoa(int(dd.Dd_array_files[j].Dd_file_ap_inodo))+ "\n")
                    }
                }*/
                buffer.WriteString("node"+  strconv.Itoa(i) + "[shape=record, label=\"" + "{ dd " + strArray[i] + "|")
                for j:=0;j<5;j++{
                    if convertName(dd.Dd_array_files[j].Dd_file_nombre[:]) != convertName(dos[:]){
                        buffer.WriteString("{<f" + strconv.Itoa(j) + "> " + convertName(dd.Dd_array_files[j].Dd_file_nombre[:]) + "| <f" + strconv.Itoa(j+1) + "> " + strconv.Itoa(int(dd.Dd_array_files[j].Dd_file_ap_inodo)) + "} |")
                    }else {
                        buffer.WriteString("{-1 | } |")
                    }

                }
                if dd.Dd_ap_detalle_directorio != -1{
                    noDirectorio = dd.Dd_ap_detalle_directorio
                    buffer.WriteString("{" + strconv.Itoa(int(dd.Dd_ap_detalle_directorio))+ " | <f10>  }}\"];\n")
                    buffer.WriteString("node"+  strconv.Itoa(i)+":f10 -> " + "node"+strconv.Itoa(int(dd.Dd_ap_detalle_directorio)))
                }else{
                    buffer.WriteString("{*1 | <f10>  }}\"];\n")
                }
                buffer.WriteString("\n")
            }
        }
    }
    //Crear Archivo
    buffer.WriteString("\n}")
    var datos string
    datos = string(buffer.String())
    CreateArchivo(pathCarpeta,datos)
    return true
}
func Reporte_tree_file(idParticion string,pathCarpeta string,ruta string,ListaDiscos *list.List)(bool){
    var bloquesGraficar [100]int
    carpetas_Graficar := strings.Split(ruta, "/")
    var buffer bytes.Buffer
    var noDirectorio int64 =0
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
    f.Seek(sb.Sb_ap_arbol_directorio,0)
    for i:=0;i<int(sb.Sb_arbol_virtual_count);i++{
        err = binary.Read(f, binary.BigEndian, &avd)
        if avd.Avd_nomre_directotrio == dos{
            break
        }
        if convertName(avd.Avd_nomre_directotrio[:]) == carpetas_Graficar[len(carpetas_Graficar)-2]{
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
            noDirectorio = avd.Avd_ap_detalle_directorio
            if EstaLlenoDD(avd.Avd_ap_detalle_directorio,sb.Sb_ap_detalle_directorio,sb.Sb_detalle_directorio_count,pathDisco){
                strArray[i] = convertName(avd.Avd_nomre_directotrio[:])
                buffer.WriteString("nodo"+ strconv.Itoa(i) + ":f6 -> node"+strconv.Itoa(int( avd.Avd_ap_detalle_directorio))+ "\n")
            }
            buffer.WriteString("nodo" + strconv.Itoa(i) + "[ shape=record, label =\"" + "{"+ convertName(avd.Avd_nomre_directotrio[:])+ "|{<f0> "+strconv.Itoa(int(avd.Avd_ap_array_subdirectoios[0])) +"|<f1>"+strconv.Itoa(int(avd.Avd_ap_array_subdirectoios[1])) + "|<f2> " + strconv.Itoa(int(avd.Avd_ap_array_subdirectoios[2])) + "|<f3> " + strconv.Itoa(int(avd.Avd_ap_array_subdirectoios[3])) + "|<f4> "+ strconv.Itoa(int(avd.Avd_ap_array_subdirectoios[4])) + "|<f5>" +  strconv.Itoa(int(avd.Avd_ap_array_subdirectoios[5])) + "|<f6>" +  strconv.Itoa(int(avd.Avd_ap_detalle_directorio)) + "|<f7> "+  strconv.Itoa(int(avd.Avd_ap_arbol_virtual_directorio)) + "}}\"];\n")
        }
    }
    /*
    Graficar DD's
    */
    noInodoGraficar:=0
    f.Seek(sb.Sb_ap_detalle_directorio,0)
    dd := DD{}
    for i:=0;i<int(sb.Sb_detalle_directorio_count);i++{
        err = binary.Read(f, binary.BigEndian, &dd)
        if dd.Ocupado == 0{
            break
        }
        if noDirectorio == int64(i){
            if EstaLlenoDD(int64(i),sb.Sb_ap_detalle_directorio,sb.Sb_detalle_directorio_count,pathDisco){
                for j:=0;j<5;j++{
                    if convertName(dd.Dd_array_files[j].Dd_file_nombre[:]) == carpetas_Graficar[len(carpetas_Graficar)-1]{
                        noInodoGraficar = int(dd.Dd_array_files[j].Dd_file_ap_inodo)
                        buffer.WriteString("node"+ strconv.Itoa(i) + ":f"+strconv.Itoa(j+1) +"->  nodex" + strconv.Itoa(int(dd.Dd_array_files[j].Dd_file_ap_inodo))+ "\n")
                    }
                }
                buffer.WriteString("node"+  strconv.Itoa(i) + "[shape=record, label=\"" + "{ dd " + strArray[i] + "|")
                for j:=0;j<5;j++{
                    if convertName(dd.Dd_array_files[j].Dd_file_nombre[:]) != convertName(dos[:]){
                        buffer.WriteString("{<f" + strconv.Itoa(j) + "> " + convertName(dd.Dd_array_files[j].Dd_file_nombre[:]) + "| <f" + strconv.Itoa(j+1) + "> " + strconv.Itoa(int(dd.Dd_array_files[j].Dd_file_ap_inodo)) + "} |")
                    }else {
                        buffer.WriteString("{-1 | } |")
                    }

                }
                if dd.Dd_ap_detalle_directorio != -1{
                    noDirectorio = dd.Dd_ap_detalle_directorio
                    buffer.WriteString("{" + strconv.Itoa(int(dd.Dd_ap_detalle_directorio))+ " | <f10>  }}\"];\n")
                    buffer.WriteString("node"+  strconv.Itoa(i)+":f10 -> " + "node"+strconv.Itoa(int(dd.Dd_ap_detalle_directorio)))
                }else{
                    buffer.WriteString("{*1 | <f10>  }}\"];\n")
                }
                buffer.WriteString("\n")
            }
        }
    }
    /*
    Graficar Inodo's
    X para identificarlos
    */
    f.Seek(sb.Sb_ap_tabla_inodo,0)
    inodo := Inodo{}
    cont1 :=0
    for i:=0;i<int(sb.Sb_inodos_count);i++{
        err = binary.Read(f, binary.BigEndian, &inodo)
        if inodo.I_count_inodo == -1{
            break
        }
        if noInodoGraficar == i{
                if inodo.I_ao_indirecto != -1{
                noInodoGraficar = i+1
                buffer.WriteString("nodex"+strconv.Itoa(int(inodo.I_count_inodo)) + "[shape=record, label=\"{Inodo"+ strconv.Itoa(int(inodo.I_count_inodo)) +"|{"+ strconv.Itoa(int(inodo.I_array_bloques[0])) +"| <f0> }|{" + strconv.Itoa(int(inodo.I_array_bloques[1])) + "| <f1> }|{" + strconv.Itoa(int(inodo.I_array_bloques[2])) +" | <f2> }|{" + strconv.Itoa(int(inodo.I_array_bloques[3])) +"| <f3> }|{" + strconv.Itoa(int(inodo.I_ao_indirecto)) +" | <f4> }}\"];"+ "\n")
                buffer.WriteString("nodex"+strconv.Itoa(int(inodo.I_count_inodo)) + " :f4 ->" + "nodex"+strconv.Itoa(int(inodo.I_ao_indirecto))+"\n")
                for h:=0;h<4;h++{
                    if inodo.I_array_bloques[h]==-1{
                        break
                    }else{
                        buffer.WriteString("nodex"+strconv.Itoa(int(inodo.I_count_inodo)) + " :f" + strconv.Itoa(h)  +"-> data"+strconv.Itoa(int(inodo.I_array_bloques[h]))+"\n")
                        bloquesGraficar[cont1]=int(inodo.I_array_bloques[h])
                        cont1++
                    }
                }
            }else{
                buffer.WriteString("nodex"+strconv.Itoa(int(inodo.I_count_inodo)) + "[shape=record, label=\"{Inodo"+ strconv.Itoa(int(inodo.I_count_inodo)) +"|{"+ strconv.Itoa(int(inodo.I_array_bloques[0])) +"| <f0> }|{" + strconv.Itoa(int(inodo.I_array_bloques[1])) + "| <f1> }|{" + strconv.Itoa(int(inodo.I_array_bloques[2])) +" | <f2> }|{" + strconv.Itoa(int(inodo.I_array_bloques[3])) +"| <f3> }|{*" + strconv.Itoa(int(inodo.I_ao_indirecto)) +" | <f4> }}\"];"+ "\n")
                for h:=0;h<4;h++{
                    if inodo.I_array_bloques[h]==-1{
                        break
                    }else{
                        bloquesGraficar[cont1]=int(inodo.I_array_bloques[h])
                        buffer.WriteString("nodex"+strconv.Itoa(int(inodo.I_count_inodo)) + " :f" + strconv.Itoa(h)  +"-> data"+strconv.Itoa(int(inodo.I_array_bloques[h]))+"\n")
                        cont1++
                    }
                }
            }
        }
    }
    cont1 =0
    f.Seek(sb.Sb_ap_bloques,0)
    data := Bloque{}
    for i:=0;i<int(sb.Sb_bloques_count);i++{
        err = binary.Read(f, binary.BigEndian, &data)
        if data.Db_data[0]==0{
            break
        }
        if bloquesGraficar[cont1]==i{
            buffer.WriteString("data" + strconv.Itoa(i) +"[shape=record, label=\"{data| <f1> "+ convertBloqueData(data.Db_data[:]) + "}}\"];\n")
            cont1++
        }   
    }
    //Crear Archivo
    buffer.WriteString("\n}")
    var datos string
    datos = string(buffer.String())
    CreateArchivo(pathCarpeta,datos)
    return false
}
func Reporte(idParticion string,ListaDiscos *list.List)(bool){
	//GraficarTreeFull(idParticion,ListaDiscos)
    //GraficarMBR_EBR(idParticion,ListaDiscos)
    //GraficarDisco(idParticion,ListaDiscos)
    //GraficarSuperBloque(idParticion,ListaDiscos)
    return false
}
func buscarDD(idParticion string,ListaDiscos *list.List,carpeta string,archivo string)(int64){
    sb:=SB{}
    var dos [15]byte
    avd := AVD{}
    //var InicioParticion int64 =0
    pathDisco,nombreParticion,_:=RecorrerListaDisco(idParticion,ListaDiscos)
    sb,_= DevolverSuperBlque(pathDisco,nombreParticion)
    f, err := os.OpenFile(pathDisco,os.O_RDWR,0755)
    if err != nil {
        fmt.Println("No existe la ruta"+ pathDisco)
        return 0
    }
    defer f.Close()
    f.Seek(sb.Sb_ap_arbol_directorio,0)
    for i:=0;i<int(sb.Sb_arbol_virtual_count);i++{
        err = binary.Read(f, binary.BigEndian, &avd)
        if avd.Avd_nomre_directotrio == dos{
            break
        }
        if convertName(avd.Avd_nomre_directotrio[:]) == carpeta{
           ddPoscion:= avd.Avd_ap_detalle_directorio
           f.Seek(sb.Sb_ap_detalle_directorio,0)
            dd := DD{}
            for i:=0;i<int(sb.Sb_detalle_directorio_count);i++{
                err = binary.Read(f, binary.BigEndian, &dd)
                if dd.Ocupado == 0{
                    break
                }
                if int64(i) == ddPoscion{
                    for l:=0;l<5;l++{
                        if convertName(dd.Dd_array_files[l].Dd_file_nombre[:])==archivo{
                            return ddPoscion
                        }
                    }
                }
            }
        }
    }
    return 0
}

func GraficarTreeFull(idParticion string,pathCarpeta string,ruta string,ListaDiscos *list.List)(bool){
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
            strArray[i] = convertName(avd.Avd_nomre_directotrio[:])
    		buffer.WriteString("nodo"+ strconv.Itoa(i) + ":f6 -> node"+strconv.Itoa(int( avd.Avd_ap_detalle_directorio))+ "\n")
    	}
    	buffer.WriteString("nodo" + strconv.Itoa(i) + "[ shape=record, label =\"" + "{"+ convertName(avd.Avd_nomre_directotrio[:])+ "|{<f0> "+strconv.Itoa(int(avd.Avd_ap_array_subdirectoios[0])) +"|<f1>"+strconv.Itoa(int(avd.Avd_ap_array_subdirectoios[1])) + "|<f2> " + strconv.Itoa(int(avd.Avd_ap_array_subdirectoios[2])) + "|<f3> " + strconv.Itoa(int(avd.Avd_ap_array_subdirectoios[3])) + "|<f4> "+ strconv.Itoa(int(avd.Avd_ap_array_subdirectoios[4])) + "|<f5>" +  strconv.Itoa(int(avd.Avd_ap_array_subdirectoios[5])) + "|<f6>" +  strconv.Itoa(int(avd.Avd_ap_detalle_directorio)) + "|<f7> "+  strconv.Itoa(int(avd.Avd_ap_arbol_virtual_directorio)) + "}}\"];\n")
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
                if convertName(dd.Dd_array_files[j].Dd_file_nombre[:]) != convertName(dos[:]){
                    buffer.WriteString("node"+ strconv.Itoa(i) + ":f"+strconv.Itoa(j+1) +"->  nodex" + strconv.Itoa(int(dd.Dd_array_files[j].Dd_file_ap_inodo))+ "\n")
                }
            }
    		buffer.WriteString("node"+  strconv.Itoa(i) + "[shape=record, label=\"" + "{ dd " + strArray[i] + "|")
    		for j:=0;j<5;j++{
                if convertName(dd.Dd_array_files[j].Dd_file_nombre[:]) != convertName(dos[:]){
                    buffer.WriteString("{<f" + strconv.Itoa(j) + "> " + convertName(dd.Dd_array_files[j].Dd_file_nombre[:]) + "| <f" + strconv.Itoa(j+1) + "> " + strconv.Itoa(int(dd.Dd_array_files[j].Dd_file_ap_inodo)) + "} |")
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
        err = binary.Read(f, binary.BigEndian, &data)
        if data.Db_data[0]==0{
            break
        }
        buffer.WriteString("data" + strconv.Itoa(i) +"[shape=record, label=\"{data| <f1> "+ convertBloqueData(data.Db_data[:]) + "}}\"];\n")

    }
    buffer.WriteString("\n}")
    var datos string
    datos = string(buffer.String())
    CreateArchivo(pathCarpeta,datos)
    return false
}
func GraficarMBR_EBR(idParticion string,ListaDiscos *list.List,path string)(bool){
    var NombreParticion [15]byte
    var buffer bytes.Buffer
    buffer.WriteString("digraph G{\nsubgraph cluster{\nlabel=\"MBR\"\ntbl[shape=box,label=<\n<table border='0' cellborder='1' cellspacing='0' width='300'  height='200' >\n<tr>  <td width='150'> <b>Nombre</b> </td> <td width='150'> <b>Valor</b> </td>  </tr>\n")
    pathDisco,_,_:=RecorrerListaDisco(idParticion,ListaDiscos)
    f, err := os.OpenFile(pathDisco,os.O_RDWR,0755)
    if err != nil {
        fmt.Println("No existe la ruta"+ pathDisco)
        return false
    }
    defer f.Close()
    mbr := MBR{}
    f.Seek(0,0)
    err = binary.Read(f, binary.BigEndian, &mbr)
    buffer.WriteString("<tr>  <td><b>mbr_tamaño</b></td><td>"+ strconv.Itoa(int(mbr.MbrTamanio)) +"</td>  </tr>\n")
    buffer.WriteString("<tr>  <td><b>mbr_Fecha_Creacion</b></td><td>" + convertName(mbr.MbrFechaCreacion[:]) + "</td>  </tr>\n")
    buffer.WriteString("<tr>  <td><b>mbr_Disk_Signature</b></td><td>" + strconv.Itoa(int(mbr.NoIdentificador)) + "</td>  </tr>\n")
    buffer.WriteString("<tr>  <td><b>mbr_Disk_Fit</b></td><td>" + "F" + "</td>  </tr>")
    Particiones := mbr.Particiones
    for i:=0;i<4;i++{
            if convertName(Particiones[i].NombreParticion[:]) != convertName(NombreParticion[:]){
                    buffer.WriteString("<tr>  <td><b>part_Status_"+ strconv.Itoa(i+1)+ "</b></td><td>" + convertName(Particiones[i].Status_particion[:]) + "</td>  </tr>\n")
                    buffer.WriteString("<tr>  <td><b>Tipo_Particion_"+ strconv.Itoa(i+1)+ "</b></td><td>" + convertName(Particiones[i].TipoParticion[:]) + "</td>  </tr>\n")
                    buffer.WriteString("<tr>  <td><b>Tipo_Ajuste_"+ strconv.Itoa(i+1)+ "</b></td><td>" + convertName(Particiones[i].TipoAjuste[:]) + "</td>  </tr>\n")
                    buffer.WriteString("<tr>  <td><b>Inicio_Particion_"+ strconv.Itoa(i+1)+ "</b></td><td>" + strconv.Itoa(int(Particiones[i].Inicio_particion)) + "</td>  </tr>\n")
                    buffer.WriteString("<tr>  <td><b>TamanioTotal_"+strconv.Itoa(i+1) + "</b></td><td>" + strconv.Itoa(int(Particiones[i].TamanioTotal)) + "</td>  </tr>\n")
                    buffer.WriteString("<tr>  <td><b>NombreParticion_"+ strconv.Itoa(i+1)+ "</b></td><td>" + convertName(Particiones[i].NombreParticion[:]) + "</td>  </tr>\n")
            }   
        }
    buffer.WriteString("</table>\n>];\n}")
    //Verificar si hay EBR
    for i:=0;i<4;i++{
            if  strings.ToLower(BytesToString(Particiones[i].TipoParticion)) == "e"{
                buffer.WriteString("subgraph cluster_1{\n label=\"EBR_Inicial\"\ntbl_1[shape=box, label=<\n<table border='0' cellborder='1' cellspacing='0'  width='300' height='160' >\n<tr>  <td width='150'><b>Nombre</b></td> <td width='150'><b>Valor</b></td>  </tr>\n")
                var InicioExtendida int64=Particiones[i].Inicio_particion
                f.Seek(InicioExtendida,0)
                ebr:=EBR{}
                err = binary.Read(f, binary.BigEndian, &ebr)
                    buffer.WriteString("<tr>  <td><b>Tipo_Ajuste_"+ strconv.Itoa(i+1)+ "</b></td><td>" + convertName(ebr.TipoAjuste[:]) + "</td>  </tr>\n")                                                                           
                    buffer.WriteString("<tr>  <td><b>Inicio_Particion_"+ strconv.Itoa(i+1)+ "</b></td><td>" + strconv.Itoa(int(ebr.Inicio_particion)) + "</td>  </tr>\n")
                    buffer.WriteString("<tr>  <td><b>Particion_Siguiente_"+ strconv.Itoa(i+1)+ "</b></td><td>" + strconv.Itoa(int(ebr.Particion_Siguiente)) + "</td>  </tr>\n")
                    buffer.WriteString("<tr>  <td><b>TamanioTotal_"+strconv.Itoa(i+1) + "</b></td><td>" + strconv.Itoa(int(ebr.TamanioTotal)) + "</td>  </tr>\n")
                    buffer.WriteString("<tr>  <td><b>NombreParticion_"+ strconv.Itoa(i+1)+ "</b></td><td>" + convertName(ebr.NombreParticion[:]) + "</td>  </tr>\n")
                    buffer.WriteString("</table>\n>];\n}")
                if ebr.Particion_Siguiente == -1{
                    fmt.Println("No Hay particiones Logicas")
                }else{
                    cont:=0
                    f.Seek(InicioExtendida,0)
                    err = binary.Read(f, binary.BigEndian, &ebr)
                    for {
                        if ebr.Particion_Siguiente == -1{
                            break
                        }else{
                            f.Seek(ebr.Particion_Siguiente,0)
                            err = binary.Read(f, binary.BigEndian, &ebr)
                            cont++
                        }
                        buffer.WriteString("subgraph cluster_"+ strconv.Itoa(cont+1) + "{\n label=\"EBR_"+strconv.Itoa(cont)+"\"\ntbl_"+ strconv.Itoa(cont+1) + "[shape=box, label=<\n<table border='0' cellborder='1' cellspacing='0'  width='300' height='160' >\n<tr>  <td width='150'><b>Nombre</b></td> <td width='150'><b>Valor</b></td>  </tr>\n")
                        buffer.WriteString("<tr>  <td><b>Tipo_Ajuste_"+ strconv.Itoa(cont)+ "</b></td><td>" + convertName(ebr.TipoAjuste[:]) + "</td>  </tr>\n")                                                                           
                        buffer.WriteString("<tr>  <td><b>Inicio_Particion_"+ strconv.Itoa(cont)+ "</b></td><td>" + strconv.Itoa(int(ebr.Inicio_particion)) + "</td>  </tr>\n")
                        buffer.WriteString("<tr>  <td><b>Particion_Siguiente_"+ strconv.Itoa(cont)+ "</b></td><td>" + strconv.Itoa(int(ebr.Particion_Siguiente)) + "</td>  </tr>\n")
                        buffer.WriteString("<tr>  <td><b>TamanioTotal_"+strconv.Itoa(cont) + "</b></td><td>" + strconv.Itoa(int(ebr.TamanioTotal)) + "</td>  </tr>\n")
                        buffer.WriteString("<tr>  <td><b>NombreParticion_"+ strconv.Itoa(cont)+ "</b></td><td>" + convertName(ebr.NombreParticion[:]) + "</td>  </tr>\n")
                        buffer.WriteString("</table>\n>];\n}")
                    }
                }
            }
        }
    buffer.WriteString("}\n")
    var datos string
    datos = string(buffer.String())
    CreateArchivo(path,datos)
    return false
}
//Graficar Disco y calcular Porcentajes
func GraficarDisco(idParticion string,ListaDiscos *list.List,path string)(bool){
    var NombreParticion [15]byte
    var buffer bytes.Buffer
    buffer.WriteString("digraph G{\ntbl [\nshape=box\nlabel=<\n<table border='0' cellborder='2' width='100' height=\"30\" color='lightblue4'>\n<tr>")
    pathDisco,_,_:=RecorrerListaDisco(idParticion,ListaDiscos)
    f, err := os.OpenFile(pathDisco,os.O_RDWR,0755)
    if err != nil {
        fmt.Println("No existe la ruta"+ pathDisco)
        return false
    }
    defer f.Close()
    PorcentajeUtilizao := 0.0
    var EspacioUtilizado int64=0
    mbr := MBR{}
    f.Seek(0,0)
    err = binary.Read(f, binary.BigEndian, &mbr)
    TamanioDisco := mbr.MbrTamanio
    Particiones := mbr.Particiones
    buffer.WriteString("<td height='30' width='75'> MBR </td>")
    for i:=0;i<4;i++{
        if convertName(Particiones[i].NombreParticion[:])!=convertName(NombreParticion[:]) && strings.ToLower(BytesToString(Particiones[i].TipoParticion)) == "p"{
            PorcentajeUtilizao =(float64(Particiones[i].TamanioTotal)/float64(TamanioDisco))*100
            buffer.WriteString("<td height='30' width='75.0'>PRIMARIA <br/>"+ convertName(Particiones[i].NombreParticion[:]) +" <br/> Ocupado: "+ strconv.Itoa(int(PorcentajeUtilizao))+"%</td>")
            EspacioUtilizado+=Particiones[i].TamanioTotal
        }else if convertName(Particiones[i].Status_particion[:])=="0"{
            buffer.WriteString("<td height='30' width='75.0'>Libre</td>")
        }
        if  strings.ToLower(BytesToString(Particiones[i].TipoParticion)) == "e"{
            EspacioUtilizado+=Particiones[i].TamanioTotal
            PorcentajeUtilizao =(float64(Particiones[i].TamanioTotal)/float64(TamanioDisco))*100
            buffer.WriteString("<td  height='30' width='15.0'>\n")
            buffer.WriteString("<table border='5'  height='30' WIDTH='15.0' cellborder='1'>\n")
            buffer.WriteString(" <tr>  <td height='60' colspan='100%'>EXTENDIDA <br/>"+ convertName(Particiones[i].NombreParticion[:]) +" <br/> Ocupado:" + strconv.Itoa(int(PorcentajeUtilizao)) +"%</td>  </tr>\n<tr>")
            var InicioExtendida int64=Particiones[i].Inicio_particion
            f.Seek(InicioExtendida,0)
            ebr:=EBR{}
            err = binary.Read(f, binary.BigEndian, &ebr)
            if ebr.Particion_Siguiente == -1{
                fmt.Println("No Hay particiones Logicas")
            }else{
                var EspacioUtilizado int64=0
                cont:=0
                f.Seek(InicioExtendida,0)
                err = binary.Read(f, binary.BigEndian, &ebr)
                for {
                    if ebr.Particion_Siguiente == -1{
                        break
                    }else{
                        f.Seek(ebr.Particion_Siguiente,0)
                        err = binary.Read(f, binary.BigEndian, &ebr)
                        EspacioUtilizado+=ebr.TamanioTotal
                        PorcentajeUtilizao =(float64(ebr.TamanioTotal)/float64(Particiones[i].TamanioTotal))*100
                        buffer.WriteString("<td height='30'>EBR</td><td height='30'> Logica:  "+ convertName(ebr.NombreParticion[:]) + " "+strconv.Itoa(int(PorcentajeUtilizao)) +"%</td>")
                        cont++
                    }
                }
                if (Particiones[i].TamanioTotal-EspacioUtilizado)>0{
                    PorcentajeUtilizao =(float64(TamanioDisco-EspacioUtilizado)/float64(TamanioDisco))*100
                    buffer.WriteString("<td height='30' width='100%'>Libre: "+ strconv.Itoa(int(PorcentajeUtilizao))+"%</td>")
                }
            }
            buffer.WriteString("</tr>\n")
            buffer.WriteString("</table>\n</td>")
        }
    } 
    if (TamanioDisco-EspacioUtilizado)>0{
        PorcentajeUtilizao =(float64(TamanioDisco-EspacioUtilizado)/float64(TamanioDisco))*100
        buffer.WriteString("<td height='30' width='75.0'>Libre: "+ strconv.Itoa(int(PorcentajeUtilizao))+"%</td>")
    }
    buffer.WriteString("     </tr>\n</table>\n>];\n}")
    var datos string
    datos = string(buffer.String())
    CreateArchivo(path,datos)
    return false
}
func GraficarSuperBloque(idParticion string,ListaDiscos *list.List,path string)(bool){
    sb:=SB{}
    pathDisco,nombreParticion,_:=RecorrerListaDisco(idParticion,ListaDiscos)
    sb,_= DevolverSuperBlque(pathDisco,nombreParticion)
    var buffer bytes.Buffer
    buffer.WriteString("digraph G{\nsubgraph cluster{\nlabel=\"Super Bloque\"\ntbl[shape=box,label=<\n<table border='0' cellborder='1' cellspacing='0' width='300'  height='200' >\n<tr>  <td width='150'> <b>Nombre</b> </td> <td width='150'> <b>Valor</b> </td>  </tr>\n")
    buffer.WriteString("<tr><td><b>Sb_nombre_hd</b></td><td>"+convertName(sb.Sb_nombre_hd[:])+"</td> </tr>*")
    buffer.WriteString("<tr><td><b>Sb_arbol_virtual_count</b></td><td>"+strconv.Itoa(int(sb.Sb_arbol_virtual_count))+"</td> </tr>\n")
    buffer.WriteString("<tr><td><b>Sb_detalle_directorio_count</b></td><td>"+ strconv.Itoa(int(sb.Sb_detalle_directorio_count)) + "</td> </tr>\n")
    buffer.WriteString("<tr><td><b>Sb_inodos_count</b></td><td>" + strconv.Itoa(int(sb.Sb_inodos_count)) + "</td> </tr>\n")
    buffer.WriteString("<tr><td><b>Sb_bloques_count</b></td><td>" +strconv.Itoa(int(sb.Sb_bloques_count)) + "</td> </tr>\n")
    buffer.WriteString("<tr><td><b>Sb_arbol_virtual_free</b></td><td>" + strconv.Itoa(int(sb.Sb_arbol_virtual_free)) + "</td> </tr>\n")
    buffer.WriteString("<tr><td><b>Sb_detalle_directorio_free</b></td><td>" + strconv.Itoa(int(sb.Sb_detalle_directorio_free))  + "</td> </tr>\n")
    buffer.WriteString("<tr><td><b>Sb_inodos_free</b></td><td>"+strconv.Itoa(int(sb.Sb_inodos_free)) +"</td> </tr>\n")
    buffer.WriteString("<tr><td><b>Sb_bloques_free</b></td><td>"+strconv.Itoa(int(sb.Sb_bloques_free)) +"</td> </tr>\n")
    buffer.WriteString("<tr><td><b>Sb_date_creacion</b></td><td>"+convertName(sb.Sb_date_creacion[:])+"</td> </tr>\n")
    buffer.WriteString("<tr><td><b>Sb_date_ultimo_montaje</b></td><td>"+convertName(sb.Sb_date_ultimo_montaje[:])+"</td> </tr>\n")
    buffer.WriteString("<tr><td><b>Sb_montajes_count</b></td><td>"+strconv.Itoa(int(sb.Sb_montajes_count)) +"</td> </tr>\n")
    buffer.WriteString("<tr><td><b>Sb_ap_bitmap_arbol_directorio</b></td><td>"+strconv.Itoa(int(sb.Sb_ap_bitmap_arbol_directorio)) +"</td> </tr>\n")
    buffer.WriteString("<tr><td><b>Sb_ap_arbol_directorio</b></td><td>"+strconv.Itoa(int(sb.Sb_ap_arbol_directorio)) +"</td> </tr>\n")
    buffer.WriteString("<tr><td><b>Sb_ap_bitmap_detalle_directorio</b></td><td>"+strconv.Itoa(int(sb.Sb_ap_bitmap_detalle_directorio)) +"</td> </tr>\n")
    buffer.WriteString("<tr><td><b>Sb_ap_detalle_directorio</b></td><td>"+strconv.Itoa(int(sb.Sb_ap_detalle_directorio)) +"</td> </tr>\n")
    buffer.WriteString("<tr><td><b>Sb_ap_bitmap_tabla_inodo</b></td><td>"+strconv.Itoa(int(sb.Sb_ap_bitmap_tabla_inodo)) +"</td> </tr>\n")
    buffer.WriteString("<tr><td><b>Sb_ap_tabla_inodo</b></td><td>"+strconv.Itoa(int(sb.Sb_ap_tabla_inodo)) +"</td> </tr>\n")
    buffer.WriteString("<tr><td><b>Sb_ap_bitmap_bloques</b></td><td>"+strconv.Itoa(int(sb.Sb_ap_bitmap_bloques)) +"</td> </tr>\n")
    buffer.WriteString("<tr><td><b>Sb_ap_bloques</b></td><td>"+strconv.Itoa(int(sb.Sb_ap_bloques)) +"</td> </tr>\n")
    buffer.WriteString("<tr><td><b>Sb_ap_log</b></td><td>"+strconv.Itoa(int(sb.Sb_ap_log)) +"</td> </tr>\n")
    buffer.WriteString("<tr><td><b>Sb_size_struct_arbol_directorio</b></td><td>"+strconv.Itoa(int(sb.Sb_size_struct_arbol_directorio)) +"</td> </tr>\n")
    buffer.WriteString("<tr><td><b>Sb_size_struct_Detalle_directorio</b></td><td>"+strconv.Itoa(int(sb.Sb_size_struct_Detalle_directorio)) +"</td> </tr>\n")
    buffer.WriteString("<tr><td><b>Sb_size_struct_inodo</b></td><td>"+strconv.Itoa(int(sb.Sb_size_struct_inodo)) +"</td> </tr>\n")
    buffer.WriteString("<tr><td><b>Sb_size_struct_bloque</b></td><td>"+strconv.Itoa(int(sb.Sb_size_struct_bloque)) +"</td> </tr>\n")
    buffer.WriteString("<tr><td><b>Sb_first_free_bit_arbol_directorio</b></td><td>"+strconv.Itoa(int(sb.Sb_first_free_bit_arbol_directorio)) +"</td> </tr>\n")
    buffer.WriteString("<tr><td><b>Sb_first_free_bit_detalle_directoriio</b></td><td>"+strconv.Itoa(int(sb.Sb_first_free_bit_detalle_directoriio)) +"</td> </tr>\n")
    buffer.WriteString("<tr><td><b>Sb_dirst_free_bit_tabla_inodo</b></td><td>"+strconv.Itoa(int(sb.Sb_dirst_free_bit_tabla_inodo)) +"</td> </tr>\n")
    buffer.WriteString("<tr><td><b>Sb_first_free_bit_bloques</b></td><td>"+strconv.Itoa(int(sb.Sb_first_free_bit_bloques)) +"</td> </tr>\n")
    buffer.WriteString("<tr><td><b>Sb_magic_num</b></td><td>"+strconv.Itoa(int(sb.Sb_magic_num)) +"</td> </tr>\n")
    buffer.WriteString("</table>\n>];\n}\n}")
    var datos string
    datos = string(buffer.String())
    CreateArchivo(path,datos)
    return false
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
                if len(convertName(dd.Dd_array_files[j].Dd_file_nombre[:])) > 0{
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
    propiedades := strings.Split(path, "/")
    nombreArchivo := propiedades[len(propiedades)-1]
	f, err := os.Create(path)

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    _, err2 := f.WriteString(data)

    if err2 != nil {
        log.Fatal(err2)
    }
    executeComand("dot -Tpdf "+ path +" -o "+ nombreArchivo[0:len(nombreArchivo)-4] +".pdf")
}	

func convertName(c []byte) string {
    n := -1
    for i, b := range c {
        if b == 0 {
            break
        }
        n = i
    }
    return string(c[:n+1])
}
func convertBloqueData(c []byte) string {
    if c[0]==32{
        return " "
    }
    n := -1
    for i, b := range c {
        if b == 32 || b==0{
            break
        }
        n = i
    }
    return string(c[:n+1])
}