mkdisk -size->30 -path->/home/edson/discoPrueba.dsk -unit->MB
#Particones
fdisk -type->P -unit->m -size->15 -fit->BF -path->/home/edson/discoPrueba.dsk -name->ParticionX
#exec -path->/home/edson/Escritorio/Proyecto/Proyecto1/entrada.sh
mount -path->/home/edson/discoPrueba.dsk -name->ParticionX
mkfs -id->vda1 -type->full
#login
login -usr->root -pwd->201701029 -id->vda1
mkdir -p -id->vda1 -path->/bin
mkdir -p -id->vda1 -path->/usr
mkdir -p -id->vda1 -path->/hola
mkdir -p -id->vda1 -path->/etc
mkdir -p -id->vda1 -path->/home
mkdir -p -id->vda1 -path->/home/user/docs
mkdir -p -id->vda1 -path->/home/user1
mkdir -p -id->vda1 -path->/home/user2
mkdir -p -id->vda1 -path->/home/user3
mkdir -p -id->vda1 -path->/home/user4
mkdir -p -id->vda1 -path->/home/user5
mkdir -p -id->vda1 -path->/home/user6
mkdir -p -id->vda1 -path->/home/user7
mkdir -p -id->vda1 -path->/home/user8
mkdir -p -id->vda1 -path->/home/user9
mkdir -p -id->vda1 -path->/home/user10
mkdir -p -id->vda1 -path->/home/user11
mkdir -p -id->vda1 -path->/home/user12
mkdir -p -id->vda1 -path->/home/user13
mkdir -p -id->vda1 -path->/home/user14
mkdir -p -id->vda1 -path->/home/user15