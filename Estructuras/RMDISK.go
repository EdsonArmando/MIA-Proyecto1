package Estructuras
import (
	"fmt"
	"strings"
)
func EjecutarComandoRMDISK(nombreComando string,propiedadesTemp []Propiedad)(ParamValidos bool){
	fmt.Println("----------------- Ejecutando RMDISK -----------------")
	ParamValidos = true
	if len(propiedadesTemp) >= 1{
	    //Recorrer la lista de propiedades
	    for i:=0; i < len(propiedadesTemp); i++{
	        var propiedadTemp = propiedadesTemp[i]
	        var nombrePropiedad string= propiedadTemp.Name
	        switch strings.ToLower(nombrePropiedad){
	        case "-path":
	        	executeComand("rm " + propiedadTemp.Val)
	        	fmt.Println("Disco Elimindo Correctamente")
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
