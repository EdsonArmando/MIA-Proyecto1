package Estructuras
import (
	"fmt"
	"container/list"
	"strings"
	"os"
	"time"
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
	mbr,sizeParticion,InicioParticion:= ReturnMBR(pathDisco,nombreParticion)
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
	//Bitmaps
	var InicioBitmapAVD int64 = InicioParticion + int64(unsafe.Sizeof(superBloque))
	var InicioAVD int64 = InicioBitmapAVD + cantidadAVD
	var InicioBitmapDD int64 = InicioAVD + (int64(unsafe.Sizeof(avd))*cantidadAVD)
	var InicioDD int64 = InicioBitmapDD + cantidadDD
	var InicioBitmapInodo int64 = InicioDD + (int64(unsafe.Sizeof(dd))*cantidadDD)
	var InicioInodo int64 = InicioBitmapInodo + cantidadInodos
	var InicioBitmapBloque int64 = InicioInodo + (int64(unsafe.Sizeof(inodo))*cantidadInodos)
	var InicioBLoque int64 = InicioBitmapBloque + cantidadBloques
	var InicioBitacora int64 = InicioBLoque + (int64(unsafe.Sizeof(bloque))*cantidadBloques)
	var InicioCopiaSB int64 = InicioBitacora + (int64(unsafe.Sizeof(bitacora))*Bitacoras)
	
	fmt.Println("----------")
	fmt.Println("pesoSB",unsafe.Sizeof(superBloque),"pesoAvd",unsafe.Sizeof(avd),"PesoDD",unsafe.Sizeof(dd),"PesoInodos",unsafe.Sizeof(inodo),"PesoBloques",unsafe.Sizeof(bloque),"PesoBitacoras",unsafe.Sizeof(bitacora))
	fmt.Println("----------")
	fmt.Println("CantidadAVD",cantidadAVD,"CantidadDD",cantidadDD,"CantidadInodos",cantidadInodos,"CantidadBloques",cantidadBloques,"CantidadBitacoras",Bitacoras)
	//Inicializando SuperBloque
	copy(superBloque.Sb_nombre_hd[:],nombreDisco)
	superBloque.Sb_arbol_virtual_count = cantidadAVD
	superBloque.Sb_detalle_directorio_count = cantidadDD
	superBloque.Sb_inodos_count = cantidadInodos
	superBloque.Sb_bloques_count = cantidadBloques
	//
	superBloque.Sb_arbol_virtual_free  = cantidadAVD
    superBloque.Sb_detalle_directorio_free = cantidadDD
    superBloque.Sb_inodos_free = cantidadInodos
    superBloque.Sb_bloques_free = cantidadBloques
    copy(superBloque.Sb_date_creacion[:],"")
    copy(superBloque.Sb_date_ultimo_montaje[:],"")
    superBloque.Sb_montajes_count = 0
    //Bitmaps
    superBloque.Sb_ap_bitmap_arbol_directorio = InicioBitmapAVD
    superBloque.Sb_ap_arbol_directorio  = InicioAVD
    superBloque.Sb_ap_bitmap_detalle_directorio = InicioBitmapDD
    superBloque.Sb_ap_detalle_directorio = InicioDD
    superBloque.Sb_ap_bitmap_tabla_inodo = InicioBitmapInodo
    superBloque.Sb_ap_tabla_inodo  = InicioInodo
    superBloque.Sb_ap_bitmap_bloques = InicioBitmapBloque
    superBloque.Sb_ap_bloques = InicioBLoque
    superBloque.Sb_ap_log = InicioBitacora
    superBloque.Sb_size_struct_arbol_directorio = int64(unsafe.Sizeof(avd))
    superBloque.Sb_size_struct_Detalle_directorio  = int64(unsafe.Sizeof(dd))
    superBloque.Sb_size_struct_inodo = int64(unsafe.Sizeof(inodo))
    superBloque.Sb_size_struct_bloque = int64(unsafe.Sizeof(bloque))
    superBloque.Sb_first_free_bit_arbol_directorio = InicioBitmapAVD
    superBloque.Sb_first_free_bit_detalle_directoriio = InicioBitmapDD
    superBloque.Sb_dirst_free_bit_tabla_inodo = InicioBitmapInodo
    superBloque.Sb_first_free_bit_bloques = InicioBitmapBloque
    superBloque.Sb_magic_num = 201701029
    superBloque.ConteoAVD = 0
    //Escribir en Particion
    f, err := os.OpenFile(pathDisco,os.O_RDWR,0755)
	if err != nil {
		fmt.Println("No existe la ruta"+pathDisco)
		return false
	}
	defer f.Close()
    //Escribir Super Boot
    f.Seek(InicioParticion,0)
    err = binary.Write(f, binary.BigEndian, &superBloque)
    //Escribir Bit Map Arbol Virtual de Directorio
    f.Seek(InicioBitmapAVD,0);
    var otro int8 = 0
    var i int64 = 0
    for i=0;i<cantidadAVD;i++{
    	err = binary.Write(f, binary.BigEndian, &otro)
    }
    //Escribir Arbol de Directorio
    f.Seek(InicioAVD,0)
    i=0
    for i=0;i<cantidadAVD;i++{
    	err = binary.Write(f, binary.BigEndian, &avd)
    }
    //Escribir Bitmap Detalle Directorio
    f.Seek(InicioBitmapDD,0)
    i=0
    for i=0;i<cantidadDD;i++{
    	err = binary.Write(f, binary.BigEndian, &otro)
    }
    //Escribir Detalle Directorio
    f.Seek(InicioDD,0)
    i=0
    dd.Dd_ap_detalle_directorio = -1
    for i=0;i<cantidadDD;i++{
    	err = binary.Write(f, binary.BigEndian, &dd)
    }
    //Escribir Bitmap Tabla Inodo
    f.Seek(InicioBitmapDD,0)
    i=0
    for i=0;i<cantidadInodos;i++{
    	err = binary.Write(f, binary.BigEndian, &otro)
    }
    //Escribir Tabla Inodos
	f.Seek(InicioInodo,0)
	    i=0
	    for i=0;i<cantidadInodos;i++{
	    	err = binary.Write(f, binary.BigEndian, &inodo)
	    }
    //Escribir Bitmap BLoque de datos
    f.Seek(InicioBitmapBloque,0)
    i=0
    for i=0;i<cantidadBloques;i++{
    	err = binary.Write(f, binary.BigEndian, &otro)
    }
    //Escribir Bloque de datos
    f.Seek(InicioBLoque,0)
	    i=0
	    for i=0;i<cantidadBloques;i++{
	    	err = binary.Write(f, binary.BigEndian, &bloque)
	    }
    //Escribir Bitacoras
    f.Seek(InicioBitacora,0)
	    i=0
	    for i=0;i<Bitacoras;i++{
	    	err = binary.Write(f, binary.BigEndian, &bitacora)
	    }
    //Escribir Copia Super Boot
    f.Seek(InicioCopiaSB,0)
    err = binary.Write(f, binary.BigEndian, &superBloque)

    /*f.Seek(InicioCopiaSB,0)
    err = binary.Read(f, binary.BigEndian, &superBloque)
    fmt.Println("----------")
    fmt.Println(superBloque)
    fmt.Println("----------")*/
    //Crear Raiz  -----> /  y  archivo con usuarios
    CrearRaiz(pathDisco,InicioParticion)
	fmt.Println("NO estructuras:",noEstructuras);
	fmt.Println("Particion a formatear",nombreParticion,NoParticion)
	fmt.Println(sizeParticion)	
	fmt.Printf("Fecha: %s\n",mbr.MbrFechaCreacion)
	return false
}

func ReturnMBR(path string,nombreParticion string) (MBR,int64,int64){
	mbr := MBR{}
    var Particiones [4]Particion 
    var nombre2 [15]byte
    var size int64
    copy(nombre2[:],nombreParticion)
	f, err := os.OpenFile(path,os.O_RDONLY,0755)
	if err != nil {
		fmt.Println("No existe la ruta"+path)
		return mbr,0,0
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
			return mbr,size, Particiones[i].Inicio_particion
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
							fmt.Println("Logica Encontrada")
							return mbr,ebr.TamanioTotal,ebr.Inicio_particion
						}
						
					}
				}
			}
		}
	return mbr,0,0
}

func CrearRaiz(pathDisco string,InicioParticion int64)(bool){
	dt := time.Now()
	f, err := os.OpenFile(pathDisco,os.O_RDWR,0755)
	if err != nil {
		fmt.Println("No existe la ruta"+pathDisco)
		return false
	}
	defer f.Close()
	f.Seek(InicioParticion,0)
	sb := SB{}
	err = binary.Read(f, binary.BigEndian, &sb)
	/*
	Escribir 1 en bitmap avd y escribir avd
	*/
	f.Seek(sb.Sb_ap_bitmap_arbol_directorio,0)
	var otro int8 = 0
	otro = 1
	err = binary.Write(f, binary.BigEndian, &otro)
	bitLibre,_:=f.Seek(0, os.SEEK_CUR)
	sb.Sb_first_free_bit_arbol_directorio = bitLibre
	avd := AVD{}
	copy(avd.Avd_fecha_creacion[:], dt.String())
	copy(avd.Avd_nomre_directotrio[:],"/")
	for j:=0;j<6;j++{
		avd.Avd_ap_array_subdirectoios[j]=-1
	}
	avd.Avd_ap_detalle_directorio = 0
	avd.Avd_ap_arbol_virtual_directorio = -1
	copy(avd.Avd_proper[:],"root")
	f.Seek(sb.Sb_ap_arbol_directorio,0)
	err = binary.Write(f, binary.BigEndian, &avd)

	sb.Sb_arbol_virtual_free  = sb.Sb_arbol_virtual_free - 1
	/*
	Escribir 1 en bitmap detalleDirectorio y escribir detalleDirectorio
	*/
	f.Seek(sb.Sb_ap_bitmap_detalle_directorio,0)
	otro = 1
	err = binary.Write(f, binary.BigEndian, &otro)
	otro=0
	bitLibre,_ =f.Seek(0, os.SEEK_CUR)
	sb.Sb_first_free_bit_detalle_directoriio = bitLibre
	detalleDirectorio :=DD{}
	arregloDD :=ArregloDD{}
	copy(arregloDD.Dd_file_nombre[:],"users.txt")
	copy(arregloDD.Dd_file_date_creacion[:],dt.String())
	copy(arregloDD.Dd_file_date_modificacion[:],dt.String())
	arregloDD.Dd_file_ap_inodo = 0
	detalleDirectorio.Dd_array_files[0] = arregloDD
	f.Seek(sb.Sb_ap_detalle_directorio,0)
	err = binary.Write(f, binary.BigEndian, &detalleDirectorio)

	sb.Sb_detalle_directorio_free = sb.Sb_detalle_directorio_free - 1
	/*
	Escribir 1 en bitmap tablaInodo y escribir Inodo
	*/
	var cantidadBloque int64=CantidadBloqueUsar("1,G,root\n1,U,root,root,201701029\n")
	f.Seek(sb.Sb_ap_bitmap_tabla_inodo,0)
	otro = 1
	err = binary.Write(f, binary.BigEndian, &otro)
	otro=0
	bitLibre,_ =f.Seek(0, os.SEEK_CUR)
	sb.Sb_dirst_free_bit_tabla_inodo = bitLibre
	inodo := Inodo{}
	for j:=0;j<4;j++{
		inodo.I_array_bloques[j]=-1
	}
	inodo.I_count_inodo = 0
	inodo.I_size_archivo = 10
	inodo.I_count_bloques_asignados = cantidadBloque
	for h:=0;h<int(cantidadBloque);h++{
		inodo.I_array_bloques[h] = int64(h) 
	}
	inodo.I_ao_indirecto = -1
	inodo.I_id_proper = 201701029
	f.Seek(sb.Sb_ap_tabla_inodo,0)
	err = binary.Write(f, binary.BigEndian, &inodo)

	sb.Sb_inodos_free = sb.Sb_inodos_free - 1	
	/*
	Escribir 1 en bitmap bloqueDatos y escribir el bloque datos
	*/
	f.Seek(sb.Sb_ap_bitmap_bloques,0)
	otro = 1
	for k:=0;k<int(cantidadBloque);k++{
		err = binary.Write(f, binary.BigEndian, &otro)
	}
	otro=0
	bitLibre,_ =f.Seek(0, os.SEEK_CUR)
	sb.Sb_first_free_bit_bloques = bitLibre
	f.Seek(sb.Sb_ap_bloques,0)
	usesTxt := []byte("1,G,root\n1,U,root,root,201701029\n")
	for k:=0;k<int(cantidadBloque);k++{
		if k==0{
			bloque := Bloque{}
			copy(bloque.Db_data[:],string([]byte(usesTxt[0:25])))
			err = binary.Write(f, binary.BigEndian, &bloque)
		}else{
			bloque := Bloque{}
			copy(bloque.Db_data[:],string([]byte(usesTxt[k*25:len(usesTxt)])))
			err = binary.Write(f, binary.BigEndian, &bloque)
		}
		sb.Sb_bloques_free = sb.Sb_bloques_free - 1
	}
	/*
	Actualizar SB
	*/
	f.Seek(0,0)
	f.Seek(InicioParticion,0)
	err = binary.Write(f, binary.BigEndian, &sb)
	return false
}
func CantidadBloqueUsar(data string)(int64){
	var noBloque int64 = 0
    cont := 0
	var dataX []byte = []byte(data)
	for i:=0;i<len(dataX);i++{
		if cont == 25{
			noBloque = noBloque + 1
			cont = 0
		}
		cont++
	}
	if len(dataX)%5 != 0{
		noBloque = noBloque + 1
	}
	return noBloque
}
func DevolverSuperBlque(path string,nombreParticion string) (SB,int64){
	mbr := MBR{}
	sb := SB{}
    var Particiones [4]Particion 
    var nombre2 [15]byte
    copy(nombre2[:],nombreParticion)
	f, err := os.OpenFile(path,os.O_RDONLY,0755)
	if err != nil {
		fmt.Println("No existe la ruta"+path)
		return sb,0
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
			f.Seek(Particiones[i].Inicio_particion,0)
			err = binary.Read(f, binary.BigEndian, &sb)
			return sb , Particiones[i].Inicio_particion
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
							fmt.Println("Logica Encontrada")
							f.Seek(ebr.Inicio_particion,0)
							err = binary.Read(f, binary.BigEndian, &sb)
							return sb ,ebr.Inicio_particion
						}
						
					}
				}
			}
		}
	return sb,0
}