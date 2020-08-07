package Estructuras
import (
	"fmt"
	"strings"
)
func EjecutarComandoFDISK(nombreComando string,propiedadesTemp []Propiedad)(ParamValidos bool){
	fmt.Println("----------------- Ejecutando FDISK -----------------")
	ParamValidos = true
	if len(propiedadesTemp) >= 2{
	    //Recorrer la lista de propiedades
	    for i:=0; i < len(propiedadesTemp); i++{
	        var propiedadTemp = propiedadesTemp[i]
	        var nombrePropiedad string= propiedadTemp.Name
	        var propiedades [8]string
	        switch strings.ToLower(nombrePropiedad){
	        case "-size":
	        	propiedades [0]=propiedadTemp.Val
	        	fmt.Println(propiedadTemp.Val)
	    	case "-fit":
	    		propiedades [1]=propiedadTemp.Val
	    		fmt.Println(propiedadTemp.Val)
	        case "-unit":
	        	propiedades [2]=propiedadTemp.Val
	        	fmt.Println(propiedadTemp.Val)
	        case "-path":
	        	propiedades [3]=propiedadTemp.Val
	        	fmt.Println(propiedadTemp.Val)
        	case "-type":
        		propiedades [4]=propiedadTemp.Val
        		fmt.Println(propiedadTemp.Val)
        	case "-delete":
        		propiedades [5]=propiedadTemp.Val
        		fmt.Println(propiedadTemp.Val)
        	case "-name":
        		propiedades [6]=propiedadTemp.Val
        		fmt.Println(propiedadTemp.Val)
        	case "-add":
        		propiedades [7]=propiedadTemp.Val
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