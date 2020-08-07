package Estructuras
import (
	"fmt"
	"strings"
)

func EjecutarComandoMount(nombreComando string,propiedadesTemp []Propiedad)(ParamValidos bool){
	fmt.Println("----------------- Ejecutando MOUNT -----------------")
	ParamValidos = true
	if len(propiedadesTemp) >= 2{
	    //Recorrer la lista de propiedades
	    for i:=0; i < len(propiedadesTemp); i++{
	        var propiedadTemp = propiedadesTemp[i]
	        var nombrePropiedad string= propiedadTemp.Name
	        var propiedades [2]string
	        switch strings.ToLower(nombrePropiedad){
	        case "-name":
	        	propiedades [0]=propiedadTemp.Val
	        	fmt.Println(propiedadTemp.Val)
	        case "-path":
	        	propiedades [1]=propiedadTemp.Val
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