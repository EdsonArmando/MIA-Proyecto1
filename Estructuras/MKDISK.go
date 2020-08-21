package Estructuras
import (
	"fmt"
	"strings"
	"os/exec"
	"strconv"
	"encoding/binary"
	"time"
	"log"
	"math/rand"
	"os"
)
/*
Vector 

*/
func EjecutarComandoMKDISK(nombreComando string,propiedadesTemp []Propiedad,cont int)(ParamValidos bool){
	dt := time.Now()
	mbr1 := MBR{}
	copy(mbr1.MbrFechaCreacion[:], dt.String())
	mbr1.NoIdentificador = int64(rand.Intn(100)+cont)
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
	    	comandos +=" bs=" + strconv.Itoa((int(tamanioTotal))*1000) + " count=1"
	    	mbr1.MbrTamanio=((tamanioTotal)-1)*1000
	    }else{
	    	comandos +=" bs=" + strconv.Itoa(int(tamanioTotal)) + "mb"+ " count=1"
	    	mbr1.MbrTamanio=tamanioTotal*1000000
	    }
	    //Inicializando Particiones
	   	for i:=0; i < 4; i++{
	   		copy(mbr1.Particiones[i].Status_particion[:], "0")
	   		copy(mbr1.Particiones[i].TipoParticion[:], "")
	   		copy(mbr1.Particiones[i].TipoAjuste[:], "")
	   		mbr1.Particiones[i].Inicio_particion = 0
	   		mbr1.Particiones[i].TamanioTotal = 0
	   		copy(mbr1.Particiones[i].NombreParticion[:], "")
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
		f.Seek(0,0)
		err = binary.Write(f, binary.BigEndian, mbr1)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Disco Creado Exitosamente")
	    return ParamValidos
	}else{
		ParamValidos = false
		return ParamValidos
	}
}

func executeComand(comandos string){
	args:= strings.Split(comandos," ")
    cmd := exec.Command(args[0],args[1:]...)
    cmd.CombinedOutput()
}
func BytesToString(data [1]byte) string {
	return string(data[:])
}
func CheckError(e error) {
    if e != nil {
    	fmt.Println("Error - ----------")
        fmt.Println(e)
    }
}