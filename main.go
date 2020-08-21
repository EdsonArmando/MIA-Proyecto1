
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
    "container/list"
    "os"
    L "./Estructuras"
) 

// Main function 
func main() { 
    ListaDiscos := list.New()
    LlenarListaDisco(ListaDiscos)
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
    	L.LeerTexto(comando,ListaDiscos)
    }
} 
func LlenarListaDisco(ListaDiscos *list.List){
    IdDisco := [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", 
    "j", "k", "l", "m", "n", "o", "p", "q", 
    "r", "s", "t", "u", "v", "w", "x", "y", "z"} 
    for i:=0;i<26;i++{
        disco := L.DISCO{}
        copy(disco.Estado[:],"0")
        copy(disco.Id[:],IdDisco[i])
        for j:=0;j<len(disco.Particiones);j++{
            mount :=L.MOUNT{}
            mount.NombreParticion = ""
            mount.Id = int64(j+1)
            copy(mount.Estado[:],"0")       
            disco.Particiones[j] = mount
        }
        ListaDiscos.PushBack(disco)
    }
}   