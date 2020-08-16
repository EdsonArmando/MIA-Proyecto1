package Estructuras
import (
	"fmt"
	"strings"
	"encoding/binary"
	"unsafe"
	"os"
	"strconv"
)
func EjecutarComandoFDISK(nombreComando string,propiedadesTemp []Propiedad)(ParamValidos bool){
	fmt.Println("----------------- Ejecutando FDISK -----------------")
	ParamValidos = true
	mbr := MBR{}
	particion := Particion{}
	var startPart int64=int64(unsafe.Sizeof(mbr))
	var propiedades [8]string
	if len(propiedadesTemp) >= 2{
	    //Recorrer la lista de propiedades
	    for i:=0; i < len(propiedadesTemp); i++{
	        var propiedadTemp = propiedadesTemp[i]
	        var nombrePropiedad string= propiedadTemp.Name
	        switch strings.ToLower(nombrePropiedad){
	        case "-size":
	        	propiedades [0]=propiedadTemp.Val
	    	case "-fit":
	    		propiedades [1]=propiedadTemp.Val
	        case "-unit":
	        	propiedades [2]=propiedadTemp.Val
	        case "-path":
	        	propiedades [3]=propiedadTemp.Val
        	case "-type":
        		propiedades [4]=propiedadTemp.Val
        	case "-delete":
        		propiedades [5]=propiedadTemp.Val
        	case "-name":
        		propiedades [6]=propiedadTemp.Val
        		fmt.Println(propiedades[6])
        	case "-add":
        		propiedades [7]=propiedadTemp.Val
	    	default:
	    		fmt.Println("Error al Ejecutar el Comando")
	        }
	    }
	    //Tamanio Particion
	   	var TamanioTotalParticion int64 = 0
	    if strings.ToLower(propiedades [2])=="b"{
	    	TamanioParticion,_  := strconv.ParseInt(propiedades[0], 10, 64)
	    	TamanioTotalParticion = TamanioParticion
	    }else if strings.ToLower(propiedades [2])=="k"{
	    	TamanioParticion,_  := strconv.ParseInt(propiedades[0], 10, 64)
	    	TamanioTotalParticion = TamanioParticion*1000
	    }else if strings.ToLower(propiedades [2])=="m"{
	    	TamanioParticion,_  := strconv.ParseInt(propiedades[0], 10, 64)
	    	TamanioTotalParticion = TamanioParticion*1000000
	    }else{
	    	TamanioParticion,_  := strconv.ParseInt(propiedades[0], 10, 64)
	    	TamanioTotalParticion = TamanioParticion*1000
	    }
	    //Obtener el MBR
	    switch strings.ToLower(propiedades [4]){
	    case "p":
	    	var Particiones [4]Particion 
	    	f, err := os.OpenFile(propiedades[3],os.O_RDWR,0755)
			if err != nil {
				fmt.Println("No existe la ruta"+propiedades[3])
				return false
			}
			defer f.Close()
			f.Seek(0,0)
			err = binary.Read(f, binary.BigEndian, &mbr)
			Particiones = mbr.Particiones
			if err != nil {
				fmt.Println("No existe el archivo en la ruta")
			}
		//El mbr ya se a leido,2.Verificar si existe espacion disponible o que no lo rebase
			if  (HayEspacio(TamanioTotalParticion,mbr.MbrTamanio)){
				return false
			}//Verificar si ya hay particiones
			if BytesToString(Particiones[0].Status_particion)  == "1" {
				fmt.Println("Ya existe una particion")
				for i:=0;i<4;i++{
					//Posicion en bytes del partstar de la n particion
					startPart+=Particiones[i].TamanioTotal
					if BytesToString(Particiones[i].Status_particion)  == "0"{
						fmt.Println(startPart)
						break
					}
				}
			}
			if(HayEspacio(startPart+TamanioTotalParticion,mbr.MbrTamanio)){
				return false
			}
		//dando valores a la particion
			copy(particion.Status_particion[:],"1")
			copy(particion.TipoParticion[:], propiedades [4])
	        copy(particion.TipoAjuste[:], propiedades [1])
	        particion.Inicio_particion = startPart
	       	particion.TamanioTotal = TamanioTotalParticion
		    copy(particion.NombreParticion[:], propiedades [6])
		    //Particion creada
		    for i:=0;i<4;i++{
					if  BytesToString(Particiones[i].Status_particion)  == "0"{
		    			Particiones[i]=particion
		    			break;
		    		}	
				}
		    f.Seek(0,0)
		    mbr.Particiones = Particiones
			err = binary.Write(f, binary.BigEndian, mbr)
			ReadFile(propiedades [3])
	    case "l":
	    	fmt.Println("Logica")
	    case "e":
	    	//Particiones Extendidas
	    	var Particiones [4]Particion 
	    	f, err := os.OpenFile(propiedades[3],os.O_RDWR,0755)
			if err != nil {
				fmt.Println("No existe la ruta"+propiedades[3])
				return false
			}
			defer f.Close()
			f.Seek(0,0)
			err = binary.Read(f, binary.BigEndian, &mbr)
			Particiones = mbr.Particiones
			if err != nil {
				fmt.Println("No existe el archivo en la ruta")
			}
		//El mbr ya se a leido,2.Verificar si existe espacion disponible o que no lo rebase
			if  (HayEspacio(TamanioTotalParticion,mbr.MbrTamanio)){
				return false
			}//Verificar si ya hay particiones
			if BytesToString(Particiones[0].Status_particion)  == "1" {
				fmt.Println("Ya existe una particion")
				for i:=0;i<4;i++{
					//Posicion en bytes del partstar de la n particion
					startPart+=Particiones[i].TamanioTotal
					if BytesToString(Particiones[i].Status_particion)  == "0"{
						fmt.Println(startPart)
						break
					}
				}
			}
			if(HayEspacio(startPart+TamanioTotalParticion,mbr.MbrTamanio)){
				return false
			}
		//dando valores a la particion
			copy(particion.Status_particion[:],"1")
			copy(particion.TipoParticion[:], propiedades [4])
	        copy(particion.TipoAjuste[:], propiedades [1])
	        particion.Inicio_particion = startPart
	       	particion.TamanioTotal = TamanioTotalParticion
		    copy(particion.NombreParticion[:], propiedades [6])
		    //Particion creada
		    for i:=0;i<4;i++{
					if  BytesToString(Particiones[i].Status_particion)  == "0"{
		    			Particiones[i]=particion
		    			break;
		    		}	
				}
		    f.Seek(0,0)
		    mbr.Particiones = Particiones
		    err = binary.Write(f, binary.BigEndian, mbr)
			ReadFile(propiedades [3])
		    ebr :=EBR{}
		    copy(ebr.Status_particion[:],"1")
		    copy(ebr.TipoAjuste[:], propiedades [1])
		    ebr.Inicio_particion = startPart
		    ebr.Particion_Siguiente = 0
		    ebr.TamanioTotal = TamanioTotalParticion
		    copy(ebr.NombreParticion[:], propiedades [6])
		    f.Seek(ebr.Inicio_particion,0)
		    err = binary.Write(f, binary.BigEndian, ebr)
		    //fmt.Println("******************EBR de la extendida")
	    	fmt.Println("Extendida","Leendo EBR")
	    	ReadFileEBR(propiedades [3])
	    default:
	    	fmt.Println("Ocurrio un error")
	    }
	    //ReadFile(propiedades[3])
	    return ParamValidos
	}else{
		ParamValidos = false
		return ParamValidos
	}
}
func ReadFileEBR(path string) (funciona bool){
	f, err := os.OpenFile(path,os.O_RDONLY,0755)
	if err != nil {
		fmt.Println("No existe la ruta"+path)
		return false
	}
	defer f.Close()
	ebr := EBR{}
	f.Seek(72200,0)
	err = binary.Read(f, binary.BigEndian, &ebr)
	if err != nil {
		fmt.Println("No existe el archivo en la ruta")
	}
	fmt.Println("Tamanio del EBR")
	fmt.Println(ebr)	
	fmt.Printf("NombreExtendida: %s\n",ebr.NombreParticion)
	return true
}
func ReadFile(path string) (funciona bool){
	f, err := os.OpenFile(path,os.O_RDONLY,0755)
	if err != nil {
		fmt.Println("No existe la ruta"+path)
		return false
	}
	defer f.Close()
	mbr := MBR{}
	f.Seek(0,0)
	err = binary.Read(f, binary.BigEndian, &mbr)
	if err != nil {
		fmt.Println("No existe el archivo en la ruta")
	}
	fmt.Println("Tamanio del MBR")
	fmt.Println(mbr)	
	fmt.Printf("Fecha: %s\n",mbr.MbrFechaCreacion)
	return true
}
func HayEspacio(TamanioTotalParticion int64,tamanioDisco int64)(bool){
	if  ((TamanioTotalParticion) > tamanioDisco) || (TamanioTotalParticion<0){
			fmt.Println("ERROR ---->EL Tamanio de la particion es mayor a el tamanio del disco o el tamanio es incorrecto")
			return true
		}
	return false
}