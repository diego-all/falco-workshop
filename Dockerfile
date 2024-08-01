# Imagen base (Ubuntu 20.04)
FROM ubuntu:20.04

# Instalar utilidades (opcional)
RUN apt-get update && apt-get install -y parted

# Crear particiones (ejemplo)
RUN parted /dev/sda mklabel gpt
RUN parted /dev/sda mkpart primary ext4 0% 33%
RUN parted /dev/sda mkpart primary ext4 33% 66%
RUN parted /dev/sda mkpart primary ext4 66% 100%

# Formatear particiones (ejemplo)
RUN mkfs.ext4 /dev/sda1
RUN mkfs.ext4 /dev/sda2
RUN mkfs.ext4 /dev/sda3

# Montar particiones (ejemplo)
RUN mkdir /mnt/sda1
RUN mkdir /mnt/sda2
RUN mkdir /mnt/sda3
RUN mount /dev/sda1 /mnt/sda1
RUN mount /dev/sda2 /mnt/sda2
RUN mount /dev/sda3 /mnt/sda3

# Configurar /etc/fstab para montar particiones al inicio (opcional)
RUN echo '/dev/sda1 /mnt/sda1 ext4 defaults 0 0' >> /etc/fstab
RUN echo '/dev/sda2 /mnt/sda2 ext4 defaults 0 0' >> /etc/fstab
RUN echo '/dev/sda3 /mnt/sda3 ext4 defaults 0 0' >> /etc/fstab



# Crear carpeta .ssh para el usuario con el fin de quitar ruido de las reglas.
# 
