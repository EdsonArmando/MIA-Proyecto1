package Estructuras
import (
	"fmt"
	"container/list"
	"encoding/binary"
	"strings"
	"os"
	"time"
)

func EjecutarComandoMKDIR(nombreComando string,propiedadesTemp []Propiedad,ListaDiscos *list.List)(ParamValidos bool){
	fmt.Println("----------------- Ejecutando Mkdir -----------------")
	ParamValidos = true
	var propiedades [3]string
	pathEspacio :=" "
	if len(propiedadesTemp) >= 1{
	    //Recorrer la lista de propiedades
	    for i:=0; i < len(propiedadesTemp); i++{
	        var propiedadTemp = propiedadesTemp[i]
	        var nombrePropiedad string= propiedadTemp.Name
	        switch strings.ToLower(nombrePropiedad){
	        case "-id":
	        	propiedades [0]=propiedadTemp.Val
        	case "-path":
        		propiedades [1]=propiedadTemp.Val
    		case "-p":
    			propiedades [2]=propiedadTemp.Val
    		case "-sigue":
    			pathEspacio+=propiedadTemp.Val+" "
	    	default:
	    		fmt.Println("Error al Ejecutar el Comando")
	        }
	    }
	    if pathEspacio!=" "{
	    	pathCompleta := propiedades[1]+pathEspacio
	    	ExecuteMKDIR(propiedades[0],pathCompleta[1 : len(pathCompleta)-2],propiedades[2],ListaDiscos)
	    }else{
	    	ExecuteMKDIR(propiedades[0],propiedades[1],propiedades[2],ListaDiscos)
	    }

	    return ParamValidos
	}else{
		ParamValidos = false
		return ParamValidos
	}
}
func ExecuteMKDIR(id string,path string, p string,ListaDiscos *list.List){
	/*
	Si no existen las carpetas se crean
	*/
	if p=="-p"{
		pathDisco,nombreParticion,_:=RecorrerListaDisco(id,ListaDiscos)
		RecorrePath(path,nombreParticion,pathDisco)
		//CrearCarpeta(pathDisco,nombreParticion,path)
	}
}
func RecorrePath(path string,nombreParticion string,pathDisco string){
	/*
	Quitar las comillas al path si tiene
	*/
	EsComilla :=  path[0:1]
	if EsComilla == "\""{
		path = path[1 : len(path)-1]
	}
	//Ver si hay mas de una carpeta
	if strings.Contains(path, "/"){
		carpetas:= strings.Split(path, "/")
		if len(carpetas)==2{
			if ExisteCarpeta(pathDisco,nombreParticion,carpetas[1])==false{
				otroAvd, _ := ModificarCarpeta(pathDisco,nombreParticion,"/","")
				if otroAvd == true{
						ModificarCarpeta(pathDisco,nombreParticion,"/","/")
						CrearCarpeta(pathDisco,nombreParticion,carpetas[1])
				}else{
						if ExisteCarpeta(pathDisco,nombreParticion,carpetas[1])==false{
						CrearCarpeta(pathDisco,nombreParticion,carpetas[1])
					}
				}				
			}
		}else{
			//mkdir -p -id->vda1 -path->/home/user6/nueva
			for i:=1;i<len(carpetas);i++{
				if ExisteCarpeta(pathDisco,nombreParticion,carpetas[i])==false{
					if carpetas[i-1]==""{
						carpetas[i-1]="/"
					}
					otroAvd, _ := ModificarCarpeta(pathDisco,nombreParticion,carpetas[i-1],"")
					if otroAvd == true{
						//fmt.Println("Necesita modificar el otro avd",noCarpeta2,carpetas[i-1])
						ModificarCarpeta(pathDisco,nombreParticion,carpetas[i-1],carpetas[i-1])
						CrearCarpeta(pathDisco,nombreParticion,carpetas[i])
					}else{
						CrearCarpeta(pathDisco,nombreParticion,carpetas[i])
					}
				}
			}
		}
	}
}
func ExisteCarpeta(pathDisco string,nombreParticion string,carpetaBuscar string)(bool){
	sb:=SB{}
	var nombre2 [15]byte
	copy(nombre2[:],carpetaBuscar)
	avd := AVD{}
	sb,_= DevolverSuperBlque(pathDisco,nombreParticion)
	f, err := os.OpenFile(pathDisco,os.O_RDWR,0755)
	if err != nil {
		fmt.Println("No existe la ruta"+pathDisco)
		return false
	}
	defer f.Close()
	f.Seek(sb.Sb_ap_arbol_directorio,0)
	for i:=0;i<int(sb.Sb_arbol_virtual_count);i++{
		err = binary.Read(f, binary.BigEndian, &avd)
		if BytesNombreParticion(avd.Avd_nomre_directotrio)==BytesNombreParticion(nombre2){
			return true
		}
	}
	return false
}
/*
	Funcion para modifica Puntero de avd
*/
func ModificarCarpeta(pathDisco string,nombreParticion string,carpetaModificar string,nombreOpcional string)(bool,int64){
	puntero_avd := true
	sb := SB{}
	avd := AVD{}
	var nombre2 [15]byte
	copy(nombre2[:],carpetaModificar)
	var bitLibre int64
	//var InicioParticion int64
	sb,_= DevolverSuperBlque(pathDisco,nombreParticion)
	f, err := os.OpenFile(pathDisco,os.O_RDWR,0755)
	if err != nil {
		fmt.Println("No existe la ruta"+pathDisco)
		return false,0
	}
	defer f.Close()
	f.Seek(sb.Sb_ap_arbol_directorio,0)
	bitLibre = sb.Sb_ap_arbol_directorio
	for i:=0;i<int(sb.Sb_arbol_virtual_count);i++{
		err = binary.Read(f, binary.BigEndian, &avd)
		if BytesNombreParticion(avd.Avd_nomre_directotrio)==BytesNombreParticion(nombre2){
				if avd.Avd_ap_arbol_virtual_directorio!=-1{
					bitLibre,_=f.Seek(0, os.SEEK_CUR)
					continue
				}
			for i:=0;i<len(avd.Avd_ap_array_subdirectoios);i++{
				if avd.Avd_ap_array_subdirectoios[i]==-1{
					avd.Avd_ap_array_subdirectoios[i]=sb.ConteoAVD + 1
					//fmt.Println(avd.Avd_ap_array_subdirectoios,avd.Avd_ap_detalle_directorio)
					puntero_avd = false
					break
				}
			}
			if puntero_avd != true{
				f.Seek(bitLibre,0)
				err = binary.Write(f, binary.BigEndian, &avd)
				bitLibre=0
				break
			}else {
					if estaLlenoAVD(pathDisco,nombreParticion,carpetaModificar)==false{
						avd.Avd_ap_arbol_virtual_directorio = sb.ConteoAVD + 1
						f.Seek(bitLibre,0)
						err = binary.Write(f, binary.BigEndian, &avd)
						bitLibre=0
						CrearCarpeta(pathDisco,nombreParticion,carpetaModificar)
						return true,avd.Avd_ap_arbol_virtual_directorio
					}
					break
				}
		}
		bitLibre,_=f.Seek(0, os.SEEK_CUR)
	}
	return false,0

}
func estaLlenoAVD(pathDisco string,nombreParticion string,carpeta string)(bool){
	sb:=SB{}
	avd := AVD{}
	estaLleno:=true
	var nombre2 [15]byte
	copy(nombre2[:],carpeta)
	sb,_= DevolverSuperBlque(pathDisco,nombreParticion)
	f, err := os.OpenFile(pathDisco,os.O_RDWR,0755)
	if err != nil {
		fmt.Println("No existe la ruta"+pathDisco)
		return false
	}
	defer f.Close()
	f.Seek(sb.Sb_ap_arbol_directorio,0)
	for i:=0;i<int(sb.Sb_arbol_virtual_count);i++{
		err = binary.Read(f, binary.BigEndian, &avd)
		if BytesNombreParticion(avd.Avd_nomre_directotrio)==BytesNombreParticion(nombre2){
			if avd.Avd_ap_array_subdirectoios[5]==-1{
				estaLleno=true
			}else if avd.Avd_ap_array_subdirectoios[5]!=-1{
				estaLleno=false
			}
		}
	}
	return estaLleno
}
func CrearCarpeta(pathDisco string,nombreParticion string,carpetaHija string)(bool){
	dt := time.Now()
	var nombre2 [15]byte
	copy(nombre2[:],"")
	sb := SB{}
	avd := AVD{}
	var InicioParticion int64
	sb,InicioParticion= DevolverSuperBlque(pathDisco,nombreParticion)
	f, err := os.OpenFile(pathDisco,os.O_RDWR,0755)
	if err != nil {
		fmt.Println("No existe la ruta"+pathDisco)
		return false
	}
	defer f.Close()
	var bitLibre int64 =0
	var bitLibreDD int64 = 0
	//bitLibre,_:=f.Seek(0, os.SEEK_CUR)
	f.Seek(sb.Sb_ap_arbol_directorio,0)
	for i:=0;i<int(sb.Sb_arbol_virtual_count);i++{
		err = binary.Read(f, binary.BigEndian, &avd)
		if BytesNombreParticion(avd.Avd_nomre_directotrio)==BytesNombreParticion(nombre2){
			avdTemp := AVD{}
			copy(avdTemp.Avd_fecha_creacion[:], dt.String())
			copy(avdTemp.Avd_nomre_directotrio[:],carpetaHija)
			for j:=0;j<6;j++{
				avdTemp.Avd_ap_array_subdirectoios[j]=-1
			}
			avdTemp.Avd_ap_detalle_directorio = sb.ConteoDD + 1
			avdTemp.Avd_ap_arbol_virtual_directorio = -1
			copy(avdTemp.Avd_proper[:],global)
			f.Seek(bitLibre,0)
			/*
			Escribir AVD
			*/
			err = binary.Write(f, binary.BigEndian, &avdTemp)
			sb.Sb_arbol_virtual_free  = sb.Sb_arbol_virtual_free - 1
			sb.ConteoAVD = sb.ConteoAVD + 1
			sb.ConteoDD = sb.ConteoDD +1 
			/*
			Marcar en bitmap
			*/
			f.Seek(sb.Sb_first_free_bit_arbol_directorio,0)
			var otro int8 = 0
			otro = 1
			err = binary.Write(f, binary.BigEndian, &otro)
			bitLibre,_:=f.Seek(0, os.SEEK_CUR)
			sb.Sb_first_free_bit_arbol_directorio = bitLibre
			/*
				Escribir DD y marcar en bitmap
			*/
			f.Seek(sb.Sb_first_free_bit_detalle_directoriio,0)
			otro = 1
			err = binary.Write(f, binary.BigEndian, &otro)
			otro=0
			bitLibre,_ =f.Seek(0, os.SEEK_CUR)
			sb.Sb_first_free_bit_detalle_directoriio = bitLibre
			detalleDirectorio :=DD{}
			f.Seek(sb.Sb_ap_detalle_directorio,0)
			for i:=0;i<int(sb.Sb_detalle_directorio_count);i++{
				err = binary.Read(f, binary.BigEndian, &detalleDirectorio)
				if detalleDirectorio.Ocupado == 0{
					detalleDirectorioTemp := DD{}
					arregloDD := ArregloDD{}
					arregloDD.Dd_file_ap_inodo = -1
					for j:=0;j<5;j++{
						detalleDirectorioTemp.Dd_array_files[j] = arregloDD
					}
					detalleDirectorioTemp.Ocupado = 1
					detalleDirectorioTemp.Dd_ap_detalle_directorio = -1
					f.Seek(bitLibreDD,0)
					err = binary.Write(f, binary.BigEndian, &detalleDirectorioTemp)
					/*for j:=0;j<5;j++{
						fmt.Println(detalleDirectorioTemp.Dd_array_files[j].Dd_file_ap_inodo)
					}*/
					sb.Sb_detalle_directorio_free = sb.Sb_detalle_directorio_free - 1
					bitLibreDD = 0
					break
				}
				bitLibreDD,_=f.Seek(0, os.SEEK_CUR)
			}
			/*
			Actualizar SB
			*/
			f.Seek(InicioParticion,0)
			err = binary.Write(f, binary.BigEndian, &sb)
			bitLibre = 0
			break;
		}
		bitLibre,_=f.Seek(0, os.SEEK_CUR)
	}
	/*f.Seek(sb.Sb_ap_arbol_directorio,0)
	for i:=0;i<13;i++{
		err = binary.Read(f, binary.BigEndian, &avd)
		fmt.Printf("carpeta: %s\n",avd.Avd_nomre_directotrio)
		fmt.Println(avd.Avd_ap_detalle_directorio)
	}
	fmt.Println(sb.ConteoDD)
	
	f.Seek(sb.Sb_ap_detalle_directorio,0)
	dd := DD{}
	for i:=0;i<13;i++{
		err = binary.Read(f, binary.BigEndian, &dd)
		for j:=0;j<5;j++{
			fmt.Print(dd.Dd_array_files[j].Dd_file_ap_inodo==-1)
		}
		fmt.Println("----------")
	}*/
	return false
}
func EscribirDetalleDirectorio(){

}