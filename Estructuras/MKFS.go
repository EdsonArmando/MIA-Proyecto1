package Estructuras
import (
	"fmt"
	"container/list"
	"strings"
	"os"
	"encoding/binary"
	"unsafe"
)

func EjecutarComandoMKFS(nombreComando string,propiedadesTemp []Propiedad,ListaDiscos *list.List)(ParamValidos bool){
	fmt.Println("----------------- Ejecutando MKFS -----------------")
	ParamValidos = true
	var propiedades [4]string
	if len(propiedadesTemp) >= 1{
	    //Recorrer la lista de propiedades
	    for i:=0; i < len(propiedadesTemp); i++{
	        var propiedadTemp = propiedadesTemp[i]
	        var nombrePropiedad string= propiedadTemp.Name
	        switch strings.ToLower(nombrePropiedad){
	        case "-id":
	        	propiedades [0]=propiedadTemp.Val
        	case "-type":
        		propiedades [1]=propiedadTemp.Val
        	case "-add":
	        	propiedades [2]=propiedadTemp.Val
        	case "-unit":
	        	propiedades [3]=propiedadTemp.Val
	    	default:
	    		fmt.Println("Error al Ejecutar el Comando")
	        }
	    }
	    ExecuteMKFS(propiedades[0],ListaDiscos)
	    return ParamValidos
	}else{
		ParamValidos = false
		return ParamValidos
	}
}
func ExecuteMKFS(id string,ListaDiscos *list.List)(bool){
	idValido := IdValido(id,ListaDiscos);
	if idValido == false{
		fmt.Println("El id no existe");
		return false
	}
	Id:= strings.ReplaceAll(id, "vd","")
	NoParticion := Id[1:]
	IdDisco := Id[:1]
	pathDisco := ""
	nombreParticion:=""
	nombreDisco:=""
	for element := ListaDiscos.Front(); element != nil; element = element.Next() {
		var disco  DISCO
		disco = element.Value.(DISCO)
		if BytesToString(disco.Id) == IdDisco{
		   for i:=0;i<len(disco.Particiones);i++{
				var mountTemp = disco.Particiones[i]
				if mountTemp.Id == id{
					copy(mountTemp.EstadoMKS[:],"1");
					nombreParticion = mountTemp.NombreParticion
					pathDisco = disco.Path
					nombreDisco = disco.NombreDisco
					break
				}
			}

		}
		element.Value = disco
	}
	mbr,sizeParticion:= ReturnMBR(pathDisco,nombreParticion)
	superBloque := SB{}
	avd := AVD{}
	dd:= DD{}
	inodo:= Inodo{}
	bloque:= Bloque{}
	bitacora := Bitacora{}
	noEstructuras := (sizeParticion-(2*int64(unsafe.Sizeof(superBloque))))/
	(27+int64(unsafe.Sizeof(avd))+int64(unsafe.Sizeof(dd))+(5*int64(unsafe.Sizeof(inodo))+
	(20*int64(unsafe.Sizeof(bloque)))+int64(unsafe.Sizeof(bitacora))))

	//NO estructuras
	var cantidadAVD int64= noEstructuras
	var cantidadDD int64= noEstructuras
	var cantidadInodos int64= noEstructuras*5
	var cantidadBloques int64= 4*cantidadInodos
	var Bitacoras int64= noEstructuras	
	fmt.Println(cantidadAVD,cantidadDD,cantidadInodos,cantidadBloques,Bitacoras)
	//Inicializando SuperBloque
	copy(superBloque.Sb_nombre_hd[:],nombreDisco)
	superBloque.Sb_arbol_virtual_count = cantidadAVD
	superBloque.Sb_detalle_directorio_count = cantidadDD
	superBloque.Sb_inodos_count = cantidadInodos
	superBloque.Sb_bloques_count = cantidadBloques

	fmt.Println(noEstructuras);
	fmt.Println("Tamanio del MBR",nombreParticion,NoParticion)
	fmt.Println(sizeParticion)	
	fmt.Printf("Fecha: %s\n",mbr.MbrFechaCreacion)
	return false
}

func ReturnMBR(path string,nombreParticion string) (MBR,int64){
	mbr := MBR{}
    var Particiones [4]Particion 
    var nombre2 [15]byte
    var size int64
    copy(nombre2[:],nombreParticion)
	f, err := os.OpenFile(path,os.O_RDONLY,0755)
	if err != nil {
		fmt.Println("No existe la ruta"+path)
		return mbr,0
	}
	defer f.Close()

	f.Seek(0,0)
	err = binary.Read(f, binary.BigEndian, &mbr)
	if err != nil {
		fmt.Println("No existe el archivo en la ruta")
	}
	Particiones = mbr.Particiones
	for i:=0;i<4;i++{
		if BytesNombreParticion(Particiones[i].NombreParticion)==BytesNombreParticion(nombre2){
			size = Particiones[i].TamanioTotal
			return mbr,size
		}
	}
	for i:=0;i<4;i++{
			if  strings.ToLower(BytesToString(Particiones[i].TipoParticion)) == "e"{
				var InicioExtendida int64=Particiones[i].Inicio_particion
				f.Seek(InicioExtendida,0)
				ebr:=EBR{}
				err = binary.Read(f, binary.BigEndian, &ebr)
				if ebr.Particion_Siguiente == -1{
					fmt.Println("No Hay particiones Logicas")
				}else{
					f.Seek(InicioExtendida,0)
					err = binary.Read(f, binary.BigEndian, &ebr)
					for {
						if ebr.Particion_Siguiente == -1{
							break
						}else{
							f.Seek(ebr.Particion_Siguiente,0)
							err = binary.Read(f, binary.BigEndian, &ebr)
						}
						if BytesNombreParticion(ebr.NombreParticion)==BytesNombreParticion(nombre2){
							fmt.Println("Loica Encontrada")
							return mbr,ebr.TamanioTotal
						}
						
					}
				}
			}
		}
	return mbr,0
}