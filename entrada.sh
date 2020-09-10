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
mkFile -p -SIZE->15 -id->vda1 -path->"/home/user/docs/alice.txt" -cont->"lffgyhgyb214"
mkFile -p -SIZE->15 -id->vda1 -path->"/home/usr3/otra5.txt" -cont->"lffgyhgyb214"
mkFile -p -SIZE->15 -id->vda1 -path->"/siete/otra.txt" -cont->"lfghyhfytfgyhgyb214"
mkFile -p -SIZE->15 -id->vda1 -path->"/siete/otra2.txt" -cont->"gfyghfytfgyhgyb214"
mkFile -p -SIZE->15 -id->vda1 -path->"/siete/otra3.txt" -cont->"gfyghfytfgyhgyb214"
mkFile -p -SIZE->15 -id->vda1 -path->"/siete/Ferana.txt" -cont->"gfyghfytfgyhgyb214"
mkFile -p -SIZE->15 -id->vda1 -path->"/siete/EDson.txt" -cont->"gfyghfytfgyhgyb214"
mkFile -p -SIZE->15 -id->vda1 -path->"/siete/Luci.txt" -cont->"gfyghfytfgyhgyb214"
mkFile -p -SIZE->15 -id->vda1 -path->"/siete/dos.txt" -cont->"gfyghfytfgyhddasdsdasdsadsdkaksfdsfdsfsdfsdfsdfsdfsdfsdfdlkadlsmkdsagyb214hfytfgyhgyb214b214hfytfgyhgyb214"
mkFile -p -SIZE->15 -id->vda1 -path->"/siete/EDson.txt" -cont->"gfyghfytfgyhgyb214"
mkFile -p -SIZE->15 -id->vda1 -path->"/siete/Luci.txt" -cont->"gfyghfytfgyhgyb214"
mkFile -p -SIZE->15 -id->vda1 -path->"/home/usr8/Luci.txt" -cont->"gfyghgyhgyb214"
reporte -id->vda1