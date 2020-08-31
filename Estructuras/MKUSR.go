package Estructuras
import (
	"fmt"
	"container/list"
	"strings"
)

func EjecutarComandoMKUSR(nombreComando string,propiedadesTemp []Propiedad,ListaDiscos *list.List)(ParamValidos bool){
	fmt.Println("----------------- Ejecutando MKGRP -----------------")
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
        	case "-usr":
        		propiedades [1]=propiedadTemp.Val
        	case "-pwd":
	        	propiedades [2]=propiedadTemp.Val
        	case "-grp":
        		propiedades [3]=propiedadTemp.Val
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
func ExecuteMKUSR(){

}