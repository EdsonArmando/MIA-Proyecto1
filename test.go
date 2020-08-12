package main

/*import (
	"encoding/binary"
	"log"
	"os"
	"fmt"
)

type Integers struct {
	I1 uint16
	I2 int32
	I3 int64
}

func main() {
	f, err := os.OpenFile("/home/edson/Escritorio/Proyecto/Proyecto1/dico1.disk",os.O_WRONLY,0755)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	i := Integers{I1: 16, I2: 180, I3: 164}
	err = binary.Write(f, binary.LittleEndian, i)
	if err != nil {
		log.Fatalln(err)
	}
	ReadFile()
}

func ReadFile(){
	f, err := os.Open("/home/edson/Escritorio/Proyecto/Proyecto1/dico1.disk")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	i := Integers{}
	err = binary.Read(f, binary.LittleEndian, &i)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(i)
}*/