package Estructuras
import (
	"fmt"
	"container/list"
	"encoding/binary"
	"strings"
	"os"
)

func EjecutarComandoMKGRP(nombreComando string,propiedadesTemp []Propiedad,ListaDiscos *list.List)(ParamValidos bool){
	fmt.Println("----------------- Ejecutando MKGRP -----------------")
	ParamValidos = true
	var propiedades [2]string
	if len(propiedadesTemp) >= 1{
	    //Recorrer la lista de propiedades
	    for i:=0; i < len(propiedadesTemp); i++{
	        var propiedadTemp = propiedadesTemp[i]
	        var nombrePropiedad string= propiedadTemp.Name
	        switch strings.ToLower(nombrePropiedad){
	        case "-id":
	        	propiedades [0]=propiedadTemp.Val
        	case "-name":
        		propiedades [1]=propiedadTemp.Val
	    	default:
	    		fmt.Println("Error al Ejecutar el Comando")
	        }
	    }
	    ExecuteMKGRP(propiedades[0],propiedades[1],ListaDiscos)
	    return ParamValidos
	}else{
		ParamValidos = false
		return ParamValidos
	}
}
func ExecuteMKGRP(id string,name string,ListaDiscos *list.List){
	pathDisco,nombreParticion,_:=RecorrerListaDisco(id,ListaDiscos)
	if global == "root"{
		ModificarArchivo(pathDisco,nombreParticion,"users.txt",name)
	}else{
		fmt.Println("NO existe usuario/usuario incorrecto")
	}
	
}
func ModificarArchivo(pathDisco string,nombreParticion string,nombreArchivo string,grupo string)(bool){
	sb := SB{}
	sb,_= DevolverSuperBlque(pathDisco,nombreParticion)
	avd :=AVD{}
	dd := DD{}
	inodo := Inodo{}
	f, err := os.OpenFile(pathDisco,os.O_RDONLY,0755)
	if err != nil {
		fmt.Println("No existe la ruta"+pathDisco)
		return false
	}
	defer f.Close()
	f.Seek(sb.Sb_ap_arbol_directorio,0)
	err = binary.Read(f, binary.BigEndian, &avd)
	/*
	AVD ya esta Inicializado
	*/
	apuntadorDD := avd.Avd_ap_detalle_directorio
	f.Seek(sb.Sb_ap_detalle_directorio,0)
	for i:=0;i<int(sb.Sb_detalle_directorio_count);i++{
		err = binary.Read(f, binary.BigEndian, &dd)
		if i==int(apuntadorDD){
			break
		}
	}
	arrgloDD := ArregloDD{}
	arrgloDD = dd.Dd_array_files[0]
	apuntadorInodo := arrgloDD.Dd_file_ap_inodo
	f.Seek(sb.Sb_ap_tabla_inodo,0)
	for i:=0;i<int(sb.Sb_inodos_count);i++{
		err = binary.Read(f, binary.BigEndian, &inodo)
		if i==int(apuntadorInodo){
			break
		}
	}
	//fmt.Printf("Archivo: %s\n",arrgloDD.Dd_file_nombre)
	return false
}