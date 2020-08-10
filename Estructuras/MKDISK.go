package Estructuras
import (
	"fmt"
	"strings"
	"os/exec"
	"strconv"
)

func EjecutarComandoMKDISK(nombreComando string,propiedadesTemp []Propiedad)(ParamValidos bool){
	fmt.Println("----------------- Ejecutando MKDISK -----------------")
	comandos := "dd if=/dev/zero ";
	ParamValidos = true
	var propiedades [4]string
	if len(propiedadesTemp) >= 2{
	    //Recorrer la lista de propiedades
	    for i:=0; i < len(propiedadesTemp); i++{
	        var propiedadTemp = propiedadesTemp[i]
	        var nombrePropiedad string= propiedadTemp.Name
	        //Vector temporal de propiedades
	        switch strings.ToLower(nombrePropiedad){
	        case "-size":
	        	propiedades [0]=propiedadTemp.Val
	    	case "-fit":
	    		propiedades [1]=propiedadTemp.Val
	        case "-unit":
	        	propiedades [2]=strings.ToLower(propiedadTemp.Val)
	        case "-path":
	        	propiedades [3]=propiedadTemp.Val
	        	arr_path:=strings.Split(propiedades [3],"/")
	        	pathCompleta:=""
	      	  	for i:=0; i < len(arr_path)-1; i++{
	      	  		pathCompleta+=arr_path[i]+"/"
	        	}
	        	executeComand("mkdir "+pathCompleta)
	        	comandos +="of="+propiedades [3];
	    	default:
	    		fmt.Println("Error al Ejecutar el Comando")
	        }
	    }
	    tamanioTotal ,_ := strconv.ParseInt(propiedades[0], 10, 64)
	    if propiedades [2] == "k"{
	    	comandos +=" bs=" + strconv.Itoa((int(tamanioTotal)-1)*1024) + " count=1"
	    }else{
	    	comandos +=" bs=" + strconv.Itoa(int(tamanioTotal)-1) + "M"+ " count=1"
	    }	   
	    //com := "dd if=/dev/zero of=/home/edson/Escritorio/Proyecto/Proyecto1/dico1.disk count=1 bs=1M"
	    executeComand(comandos)
	    fmt.Println("Disco Creado Exitosamente")
	    return ParamValidos
	}else{
		ParamValidos = false
		return ParamValidos
	}
}
func executeComand(comandos string){
	fmt.Println(comandos)
	args:= strings.Split(comandos," ")
    cmd := exec.Command(args[0],args[1:]...)
    cmd.CombinedOutput()
}
func CheckError(e error) {
    if e != nil {
    	fmt.Println("Error - ----------")
        fmt.Println(e)
    }
}