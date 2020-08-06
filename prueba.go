
// Go program to illustrate how to split a string 
/*
    Universidad de San Carlos de Guatemala
    Developer: Edson Armando Guix Manuel
    Manejo e Implementacion de Archivos
*/
package main 
  
import ( 
    "io/ioutil"
    L "./Estructuras"
) 

// Main function 
func main() { 
    //LeerArchivo de Entrada
    dat, err := ioutil.ReadFile("entrada.sh")
    L.Check(err)
    L.LeerTexto(string(dat))
} 
