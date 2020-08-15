package main
import (
	"log"
	"os"
	"encoding/binary"
	"bytes"
	"unsafe"
	"fmt"
)
type mbr struct {
	Numero uint8
	Numero2 int64
	Caracter byte
	Cadena [20]byte
}
type mbr2 struct {
	Numero uint8
	Caracter byte
	Cadena [20]byte
}
func Dos() {
	/*writeFile()
	fmt.Println("Reading File: ")
	readFile()*/
	f, err := os.OpenFile("/home/edson/Escritorio/Proyecto/Proyecto1/dico1.disk",os.O_WRONLY,0755)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	if err != nil {
		log.Fatalln(err)
	}
	//Inicalizando mbr
	disco := mbr{Numero: 5, Numero2: 150}
	disco.Caracter = 'a'
	cadenita := "Hola Amigos"
	copy(disco.Cadena[:], cadenita)
	err = binary.Write(f, binary.BigEndian, disco)
	f.Seek(int64(unsafe.Sizeof(disco)),0)
	disco = mbr{Numero: 10, Numero2: 200}
	disco.Caracter = 'E'
	//Inicalizando 2mbr
	cadenita = "Hola Edson"
	copy(disco.Cadena[:], cadenita)
	err = binary.Write(f, binary.BigEndian, disco)
	//Guardando otro struct mbr2
	/*f.Seek(0,0)
	disco2 := mbr2{Numero: 10}
	disco2.Caracter = 'E'
	// Igualar cadenas a array de bytes (array de chars)
	cadenita = "Hola Edson2"
	copy(disco2.Cadena[:], cadenita)
	err = binary.Write(f, binary.BigEndian, disco2)**/
	ReadFile2()
}
func ReadFile2(){
	f, err := os.OpenFile("/home/edson/Escritorio/Proyecto/Proyecto1/dico1.disk",os.O_RDONLY,0755)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	disco := mbr{}
	f.Seek(0, 0) 
	err = binary.Read(f, binary.BigEndian, &disco)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(disco)	
	fmt.Printf("Caracter: %c\nCadena: %s\n", disco.Caracter,  disco.Cadena)
}
func escribirBytes(file *os.File, bytes []byte) {
	_, err := file.Write(bytes)

	if err != nil {
		log.Fatal(err)
	}
}
func readFile() {
	//Abrimos/creamos un archivo.
	file, err := os.OpenFile("/home/edson/Escritorio/Proyecto/Proyecto1/dico1.disk",os.O_RDONLY,0755)
	defer file.Close() 
	if err != nil { //validar que no sea nulo.
		log.Fatal(err)
	}

	//Declaramos variable de tipo mbr
	m := mbr{}
	//Obtenemos el tamanio del mbr
	var size int = int(unsafe.Sizeof(m))

	//Lee la cantidad de <size> bytes del archivo
	data := leerBytes(file, size)
	//Convierte la data en un buffer,necesario para
	//decodificar binario
	buffer := bytes.NewBuffer(data)
	
	//Decodificamos y guardamos en la variable m
	err = binary.Read(buffer, binary.BigEndian, &m)
	if err != nil {
		log.Fatal("binary.Read failed", err)
	}
	
	//Se imprimen los valores guardados en el struct
	fmt.Println(m)
	fmt.Printf("Caracter: %c\nCadena: %s\n", m.Caracter,  m.Cadena)
}
func writeFile() {
	file, err := os.OpenFile("/home/edson/Escritorio/Proyecto/Proyecto1/dico1.disk",os.O_WRONLY,0755)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	var otro int8 = 0

	s := &otro
	//Escribimos un 0 en el inicio del archivo.
	var binario bytes.Buffer
	binary.Write(&binario, binary.BigEndian, s)
	escribirBytes(file, binario.Bytes())
	//Nos posicionamos en el byte 1023 (primera posicion es 0)	
	file.Seek(1023, 0) // segundo parametro: 0, 1, 2.     0 -> Inicio, 1-> desde donde esta el puntero, 2 -> Del fin para atras
	
	//Escribimos un 0 al final del archivo.
	var binario2 bytes.Buffer
	binary.Write(&binario2, binary.BigEndian, s)
	escribirBytes(file, binario2.Bytes())

//----------------------------------------------------------------------- //
	//Escribimos nuestro struct en el inicio del archivo

	file.Seek(0, 0) // nos posicionamos en el inicio del archivo.
	
	//Asignamos valores a los atributos del struct.
	disco := mbr{Numero: 5, Numero2: 150}
	disco.Caracter = 'a'

	// Igualar cadenas a array de bytes (array de chars)
	cadenita := "Hola Amigos"
	copy(disco.Cadena[:], cadenita)

	s1 := &disco
	
	//Escribimos struct.
	var binario3 bytes.Buffer
	binary.Write(&binario3, binary.BigEndian, s1)
	escribirBytes(file, binario3.Bytes())
	//Asignamos valores a los atributos del struct.
	disco = mbr{Numero: 10, Numero2: 200}
	disco.Caracter = 'E'

	// Igualar cadenas a array de bytes (array de chars)
	cadenita = "Hola Edson"
	copy(disco.Cadena[:], cadenita)

	s1 = &disco
	
	//Escribimos struct.
	var binario4 bytes.Buffer
	binary.Write(&binario4, binary.BigEndian, s1)
	escribirBytes(file, binario4.Bytes())
}
func leerBytes(file *os.File, number int) []byte {
	file.Seek(int64(number)-10,0)
	bytes := make([]byte, number) //array de bytes

	_, err := file.Read(bytes) // Leido -> bytes
	if err != nil {
		log.Fatal(err)
	}

	return bytes
}