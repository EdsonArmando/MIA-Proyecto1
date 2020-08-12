#Crear Discos
mkdisk -size->12 -path->/home/edson/Escritorio/Proyecto/Proyecto1/Vivo/dico1.disk -unit->MB
#Particones
fdisk -type->P -unit->K -size->1 -fit->BF -path->/home/edson/archivos/DiscoPrueba.disk
pause
Fdisk -sizE->72 -path->/home/Disco1.dsk -name->Particion1
#exec -path->/home/edson/Escritorio/Proyecto/Proyecto1/entrada.sh