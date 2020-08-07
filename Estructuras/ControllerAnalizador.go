package Estructuras
import (
	"container/list"
	"strings"
	"fmt"
)
//Funcion para leer y reconocer los comandos lleno la lista de comandos
func LeerTexto(dat string){
	//Leendo la cadena de entrada
	ListaComandos := list.New()
	lineaComando := strings.Split(dat, "\n")
	var c Comando
	    for i:=0; i < len(lineaComando); i++{
        EsComentario :=  lineaComando[i][0:1]
        if EsComentario != "#" {
            comando := lineaComando[i]
            propiedades := strings.Split(string(comando), " ")
            //Nombre Comando
            nombreComando := propiedades[0] 
            //Struct para el Comando
            c = Comando {Name: strings.ToLower(nombreComando)}
            propiedadesTemp := make([]Propiedad, len(propiedades)-1)
            for i:=1; i < len(propiedades); i++{
                if propiedades[i]==""{
                    continue
                }
                valor_propiedad_Comando := strings.Split(propiedades[i], "=")
                propiedadesTemp[i-1]= Propiedad{Name:valor_propiedad_Comando[0],
                Val:valor_propiedad_Comando[1]}
            }
            c.Propiedades = propiedadesTemp
            //Agregando el comando a la lista comandos
            ListaComandos.PushBack(c)
        }     
    }
    RecorrerListaComando(ListaComandos)
}
//Funcion para recorrer la Lista de Comandos
func RecorrerListaComando(ListaComandos *list.List){
    var ParamValidos bool = true
	 for element := ListaComandos.Front(); element != nil; element = element.Next() {
        var comandoTemp Comando
        comandoTemp = element.Value.(Comando)
        //Lista de propiedades del Comando
        switch strings.ToLower(comandoTemp.Name){
        case "mkdisk":
            ParamValidos = EjecutarComandoMKDISK(comandoTemp.Name,comandoTemp.Propiedades)
            if ParamValidos == false{
                fmt.Println("Parametros Invalidos")
            }
        case "rmdisk":
            ParamValidos = EjecutarComandoRMDISK(comandoTemp.Name,comandoTemp.Propiedades)
             if ParamValidos == false{
                fmt.Println("Parametros Invalidos")
            }
        case "fdisk":
            ParamValidos = EjecutarComandoFDISK(comandoTemp.Name,comandoTemp.Propiedades)
             if ParamValidos == false{
                fmt.Println("Parametros Invalidos")
            }
        case "mount":
            ParamValidos = EjecutarComandoMount(comandoTemp.Name,comandoTemp.Propiedades)
             if ParamValidos == false{
                fmt.Println("Parametros Invalidos")
            }
        case "unmount":
            ParamValidos = EjecutarComandoUnmount(comandoTemp.Name,comandoTemp.Propiedades)
             if ParamValidos == false{
                fmt.Println("Parametros Invalidos")
            }
        case "exec":
            ParamValidos = EjecutarComandoExec(comandoTemp.Name,comandoTemp.Propiedades)
             if ParamValidos == false{
                fmt.Println("Parametros Invalidos")
            }
        default:
            fmt.Println("Error al Ejecutar el Comando")
        }
    }
}

