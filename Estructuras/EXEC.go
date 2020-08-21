package Estructuras
import (
	"fmt"
	"strings"
	"io/ioutil"
	"container/list"
)
func Demo(n int) {
    fmt.Println("HI")
    fmt.Println(n)

}
func Check(e error) {
    if e != nil {
        fmt.Println("Error")
    }
}
func EjecutarComandoExec(nombreComando string,propiedadesTemp []Propiedad,ListaDiscos *list.List)(ParamValidos bool){
	fmt.Println("----------------- Ejecutando EXEC -----------------")
	ParamValidos = true
	if len(propiedadesTemp) >= 1{
	    //Recorrer la lista de propiedades
	    for i:=0; i < len(propiedadesTemp); i++{
	        var propiedadTemp = propiedadesTemp[i]
	        var nombrePropiedad string= propiedadTemp.Name
	        switch strings.ToLower(nombrePropiedad){
	        case "-path":
	        	fmt.Println(propiedadTemp.Val)
	        	dat, err := ioutil.ReadFile(propiedadTemp.Val)
			    Check(err)
			    LeerTexto(string(dat),ListaDiscos)
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