package Estructuras
import (
	"fmt"
	"strings"
	"os/exec"
	"strconv"
	"encoding/binary"
	"log"
	"os"
)

func EjecutarComandoMKDISK(nombreComando string,propiedadesTemp []Propiedad)(ParamValidos bool){
	mbr1 := MBR{}
	copy(mbr1.MbrFechaCreacion[:], "11/08/2020")
	mbr1.NoIdentificador = 25
	copy(mbr1.TipoAjuste[:], "f")
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
	    		copy(mbr1.TipoAjuste[:],propiedades [1])
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
	    	mbr1.MbrTamanio=((tamanioTotal)-1)*1024
	    }else{
	    	comandos +=" bs=" + strconv.Itoa(int(tamanioTotal)-1) + "M"+ " count=1"
	    	mbr1.MbrTamanio=tamanioTotal*1048576
	    }	   
	    //com := "dd if=/dev/zero of=/home/edson/Escritorio/Proyecto/Proyecto1/dico1.disk count=1 bs=1M"
	    executeComand(comandos)
	    //Escribir MBR
	    f, err := os.OpenFile(propiedades [3],os.O_WRONLY,0755)
		if err != nil {
			log.Fatalln(err)
		}
		defer func() {
			if err := f.Close(); err != nil {
				log.Fatalln(err)
			}
		}()
		err = binary.Write(f, binary.LittleEndian, mbr1)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Disco Creado Exitosamente")
		ReadFile()
	    return ParamValidos
	}else{
		ParamValidos = false
		return ParamValidos
	}
}
func ReadFile(){
	f, err := os.Open("/home/edson/Escritorio/Proyecto/Proyecto1/Vivo/dico1.disk")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	mbr := MBR{}
	err = binary.Read(f, binary.LittleEndian, &mbr)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("-----La fecha de creacion es")
	fmt.Println(BytesToString(mbr.MbrFechaCreacion))
}
func executeComand(comandos string){
	args:= strings.Split(comandos," ")
    cmd := exec.Command(args[0],args[1:]...)
    cmd.CombinedOutput()
}
func BytesToString(data [16]byte) string {
	return string(data[:])
}
func CheckError(e error) {
    if e != nil {
    	fmt.Println("Error - ----------")
        fmt.Println(e)
    }
}