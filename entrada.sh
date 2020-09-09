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
mkFile -p -SIZE->15 -id->vda1 -path->"/siete/otra.txt" -cont->"lkjhhgulllllgfyghyghfytsaasdasasffsdgdgfdgdfgdgdfgfdfggyb214lkjhhggfyghfytfgyhgyb214lkjhhgugfyghyghfytfggyb214lkjhhggfyghfytfgyhgyb214"
mkFile -p -SIZE->15 -id->vda1 -path->"/siete/otra2.txt" -cont->"gfyghfytfgyhgyb214"
mkFile -p -SIZE->15 -id->vda1 -path->"/siete/otra3.txt" -cont->"gfyghfytfgyhgyb214"
mkFile -p -SIZE->15 -id->vda1 -path->"/siete/otra4.txt" -cont->"gfyghfytfgyhgyb214"
mkFile -p -SIZE->15 -id->vda1 -path->"/siete/otra5.txt" -cont->"gfyghfytfgyhgyb214"
mkFile -p -SIZE->15 -id->vda1 -path->"/siete/otra6.txt" -cont->"gfyghfytfgyhgyb214"
mkFile -p -SIZE->15 -id->vda1 -path->"/siete/ferna.txt" -cont->"gfyghfytfgyhgyb214"
mkFile -p -SIZE->15 -id->vda1 -path->"/siete/dulc.txt" -cont->"gfyghfytfgyhgyb214"
mkFile -p -SIZE->15 -id->vda1 -path->"/siete/otra6.txt" -cont->"gfyghfytfgyhgyb214"
mkFile -p -SIZE->15 -id->vda1 -path->"/siete/ferna.txt" -cont->"gfyghfytfgyhgyb214"
mkFile -p -SIZE->15 -id->vda1 -path->"/siete/dulc.txt" -cont->"gfyghfytfgyhgyb214"
reporte -id->vda1