#Crear Discos
mkdisk -size->12 -path->/home/edson/Escritorio/Proyecto/Proyecto1/Vivo/dico1.disk -unit->MB
#Crear Discos
mkdisk -size->12 -path->/home/edson/Escritorio/Proyecto/Proyecto1/Vivo/dico2.disk -unit->MB
#Crear Discos
mkdisk -size->12 -path->/home/edson/Escritorio/Proyecto/Proyecto1/Vivo2/dico1.disk -unit->MB
#Crear Discos
mkdisk -size->12 -path->/home/edson/Escritorio/Proyecto/Proyecto1/Vivo2/dico2.disk -unit->MB
#Eliminar Disco 
rmdisk -path->/home/edson/Escritorio/Proyecto/Proyecto1/Vivo2/dico2.disk
#Eliminar Disco 
rmdisk -path->/home/edson/Escritorio/Proyecto/Proyecto1/Vivo2/dico1.disk
#Particones
fdisk -type->P -unit->K -size->1 -fit->BF -path->/home/edson/archivos/DiscoPrueba.disk
Fdisk -sizE->72 -path->/home/Disco1.dsk -name->Particion1
#exec -path->/home/edson/Escritorio/Proyecto/Proyecto1/entrada.sh