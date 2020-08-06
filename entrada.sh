#Crear Discos
mkdisk -size=3 -unit=m -path=/home/edson/archivos/DiscoPrueba.disk
#Eliminar Disco 
rmdisk -path=/home/edson/MIA/Luciana/DiscoSeis.disk
#Particones
fdisk -type=P -unit=K -size=1 -fit=BF -path=/home/edson/archivos/DiscoPrueba.disk