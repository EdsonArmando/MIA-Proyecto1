package Estructuras
import (
	"fmt"
	"strings"
)
func EjecutarComandoUnmount(nombreComando string,propiedadesTemp []Propiedad)(ParamValidos bool){
	fmt.Println("----------------- Ejecutando UNMOUNT -----------------")
	ParamValidos = true
	if len(propiedadesTemp) >= 1{
	    //Recorrer la lista de propiedades
	    for i:=0; i < len(propiedadesTemp); i++{
	        var propiedadTemp = propiedadesTemp[i]
	        var nombrePropiedad string= propiedadTemp.Name
	        switch strings.ToLower(nombrePropiedad){
	        case "-id":
	        	fmt.Println(propiedadTemp.Val)
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