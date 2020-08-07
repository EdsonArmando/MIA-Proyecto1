
// Go program to illustrate how to split a string 
/*
    Universidad de San Carlos de Guatemala
    Developer: Edson Armando Guix Manuel
    Manejo e Implementacion de Archivos
*/
package main 
  
import ( 
    "bufio"
    "fmt"
    "os"
    L "./Estructuras"
) 

// Main function 
func main() { 
    //LeerArchivo de Entrada
    var comando string = ""
    scanner := bufio.NewScanner(os.Stdin)
    for comando != "exit"{
    	fmt.Println("--------------------------------------------------------------------------------")
    	fmt.Println("                            Ingrese un comando                         ")
    	fmt.Println("--------------------------------------------------------------------------------")
    	fmt.Println(">>")
    	scanner.Scan()
    	comando = scanner.Text()
    	L.LeerTexto(comando)
    }
} 
