package Estructuras
import ( 
) 
//Estructura para cada Comando y sus Propiedades
type Propiedad struct{
    Name string
    Val  string
}
type Comando struct {
    Name string
    Propiedades []Propiedad
}
//Estructuras para el Disco y Particiones
type Particion struct{
        Status_particion [1]byte
        TipoParticion [1]byte
        TipoAjuste [2]byte
        Inicio_particion int64
        TamanioTotal int64
        NombreParticion [15]byte
}
    //Struct para el MBR
 type  MBR struct{
        MbrTamanio int64
        MbrFechaCreacion [19]byte
        NoIdentificador int64
        Particiones [4]Particion
}
//Struct para las particiones Logicas
type EBR struct{
        Status_particion [1]byte
        TipoAjuste [2]byte
        Inicio_particion int64
        Particion_Siguiente int64
        TamanioTotal int64
        NombreParticion [15]byte
}
//EStruc de las particiones montadas
type MOUNT struct{
    NombreParticion string
    Id string
    Estado [1]byte
    EstadoMKS [1]byte
}

//Estruct Disco 
type DISCO struct{
    NombreDisco string
    Path string
    Id [1]byte 
    Estado [1]byte 
    Particiones [100]MOUNT
}
//57.51
type Integers struct {
	I1 uint16
	I2 int32
	I3 int64
	DOS byte
}

//Structuras Segunda Fase
//SuperBloque
type SB struct{
    Sb_nombre_hd [15]byte
    Sb_arbol_virtual_count int64
    Sb_detalle_directorio_count int64
    Sb_inodos_count int64
    Sb_bloques_count int64
    Sb_arbol_virtual_free int64
    Sb_detalle_directorio_free int64
    Sb_inodos_free int64
    Sb_bloques_free int64
    Sb_date_creacion [19]byte
    Sb_date_ultimo_montaje [19]byte
    Sb_montajes_count int64
    Sb_ap_bitmap_arbol_directorio int64
    Sb_ap_arbol_directorio int64
    Sb_ap_bitmap_detalle_directorio int64
    Sb_ap_detalle_directorio int64
    Sb_ap_bitmap_tabla_inodo int64
    Sb_ap_tabla_inodo int64
    Sb_ap_bitmap_bloques int64
    Sb_ap_bloques int64
    Sb_ap_log int64
    Sb_size_struct_arbol_directorio int64
    Sb_size_struct_Detalle_directorio int64
    Sb_size_struct_inodo int64
    Sb_size_struct_bloque int64
    Sb_first_free_bit_arbol_directorio int64
    Sb_first_free_bit_detalle_directoriio int64
    Sb_dirst_free_bit_tabla_inodo int64
    Sb_first_free_bit_bloques int64
    Sb_magic_num int64
    ConteoAVD int64

}

//Arbol virtual de directorio
type AVD struct{
    Avd_fecha_creacion [19]byte
    Avd_nomre_directotrio [15]byte
    Avd_ap_array_subdirectoios [6]int64
    Avd_ap_detalle_directorio int64
    Avd_ap_arbol_virtual_directorio int64
    Avd_proper [10]byte
}
//Detalle dde Directorio

type ArregloDD struct{
    Dd_file_nombre [15]byte
    Dd_file_ap_inodo int64
    Dd_file_date_creacion [19]byte
    Dd_file_date_modificacion [19]byte
}
type DD struct{
    Dd_array_files [5]ArregloDD
    Dd_ap_detalle_directorio int64
}


//Cantidad de Inodos
type Inodo struct{
    I_count_inodo int64
    I_size_archivo int64
    I_count_bloques_asignados int64
    I_array_bloques [4]int64
    I_ao_indirecto int64
    I_id_proper int64
}

//Bloque
type Bloque struct{
    Db_data [25]byte
}

//bitacora
type Bitacora struct{
    Log_tipo_operacion [19]byte
    Log_tipo [1]byte
    Log_nombre [15]byte
    Log_Contenido [25]byte
    Log_fecha [19]byte
}

//
func BytesNombreParticion(data [15]byte) string {
    return string(data[:])
}
func ConvertData(data [25]byte) string {
    return string(data[:])
}