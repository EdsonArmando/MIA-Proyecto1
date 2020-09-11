mkdisk -size->30 -path->"/home/edson/discoPrueba.dsk" -unit->m
#Particones
fdisk -type->P -unit->m -size->15 -fit->BF -path->"/home/edson/discoPrueba.dsk" -name->ParticionX
fdisk -type->P -unit->m -size->2 -fit->BF -path->"/home/edson/discoPrueba.dsk" -name->Particion1
fdisk -type->E -unit->m -size->3 -fit->BF -path->"/home/edson/discoPrueba.dsk" -name->Particion2
fdisk -type->P -unit->m -size->7 -fit->BF -path->"/home/edson/discoPrueba.dsk" -name->Particion3
fdisk -type->l -unit->m -size->1 -fit->BF -path->"/home/edson/discoPrueba.dsk" -name->Particion4
fdisk -type->l -unit->m -size->1 -fit->BF -path->"/home/edson/discoPrueba.dsk" -name->Particion5
fdisk -type->l -unit->m -size->1 -fit->BF -path->"/home/edson/discoPrueba.dsk" -name->Particion6
#exec -path->/home/edson/Escritorio/Proyecto/Proyecto1/entrada.sh
mount -path->/home/edson/discoPrueba.dsk -name->ParticionX
mkfs -id->vda1 -type->full
#login
login -usr->root -pwd->201701029 -id->vda1
mkdir -p -id->vda1 -path->/bin
mkdir -p -id->vda1 -path->/usr
mkdir -p -id->vda1 -path->/hola
mkdir -p -id->vda1 -path->/etc/
mkdir -p -id->vda1 -path->/home/user/docs
mkdir -p -id->vda1 -path->/home/usr1
mkdir -p -id->vda1 -path->/home/usr2
mkdir -p -id->vda1 -path->/home/usr3
mkdir -p -id->vda1 -path->/home/usr4
mkdir -p -id->vda1 -path->/home/usr5
mkdir -p -id->vda1 -path->/home/usr6
mkdir -p -id->vda1 -path->/home/usr7
mkdir -p -id->vda1 -path->/home/usr8
mkFile -p -SIZE->200 -id->vda1 -path->"/home/user/docs/alice.txt" -cont->"Edson Guix"
mkFile -p -SIZE->50 -id->vda1 -path->"/home/user/docs/alice2.txt" -cont->"lffgyhgyb214"
mkFile -p -SIZE->15 -id->vda1 -path->"/home/user/docs/alice3.txt" -cont->"lffgyhgyb214"
mkFile -p -SIZE->15 -id->vda1 -path->"/home/user/docs/alice4.txt" -cont->"lffgyhgyb214"
mkFile -p -SIZE->500 -id->vda1 -path->"/home/user/docs/alice8.txt" -cont->"lffgyhgyb214"
mkFile -p -SIZE->50 -id->vda1 -path->"/home/user/docs/prueba.txt" -cont->"lffgyhgyb214"
mkFile -p -SIZE->50 -id->vda1 -path->"/home/usr5/prueba4.txt" -cont->"lffgyhgyb214"
#fdisk -delete->fast -name->Particion4 -path->"/home/edson/discoPrueba.dsk"
rep -id->vda1 -path->"/home/edson/Escritorio/Proyecto/Proyecto1/Direcorio.dot" -nombre->tree_directorio -ruta->"/home/usr5"
rep -id->vda1 -path->"/home/edson/Escritorio/Proyecto/Proyecto1/TreeFull.dot" -nombre->tree_complete