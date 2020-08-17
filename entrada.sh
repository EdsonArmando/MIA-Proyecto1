#Crear Discos
mkdisk -size->21 -path->/home/edson/Escritorio/Proyecto/Proyecto1/Vivo/dico1.dsk -unit->MB
#Crear Discos
mkdisk -size->10 -path->/home/edson/Escritorio/Proyecto/Proyecto1/Vivo/disco2.dsk -unit->MB
#Particones
fdisk -type->P -unit->m -size->8 -fit->BF -path->/home/edson/Escritorio/Proyecto/Proyecto1/Vivo/dico1.dsk -name->Particion2
#Particiones
pause
fdisk -type->P -unit->m -size->10 -fit->BF -path->/home/edson/Escritorio/Proyecto/Proyecto1/Vivo/dico1.dsk -name->Particion3
pause
fdisk -type->P -unit->m -size->1 -fit->BF -path->/home/edson/Escritorio/Proyecto/Proyecto1/Vivo/dico1.dsk -name->Particion4
pause
fdisk -type->P -unit->m -size->1 -fit->BF -path->/home/edson/Escritorio/Proyecto/Proyecto1/Vivo/dico1.dsk -name->Particion5
pause
Fdisk -type->P -sizE->72 -path->/home/edson/Escritorio/Proyecto/Proyecto1/Vivo/disco2.dsk -name->Particion1
pause
Fdisk -type->E -unit->k -sizE->10 -path->/home/edson/Escritorio/Proyecto/Proyecto1/Vivo/disco2.dsk -name->Particion7
pause
Fdisk -type->l -unit->k -sizE->2 -path->/home/edson/Escritorio/Proyecto/Proyecto1/Vivo/disco2.dsk -name->Particion9
pause
Fdisk -type->l -unit->M -sizE->1 -path->/home/edson/Escritorio/Proyecto/Proyecto1/Vivo/disco2.dsk -name->Particion10
pause
Fdisk -type->l -unit->k -sizE->1 -path->/home/edson/Escritorio/Proyecto/Proyecto1/Vivo/disco2.dsk -name->Particion11L
#exec -path->/home/edson/Escritorio/Proyecto/Proyecto1/entrada.sh