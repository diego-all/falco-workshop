# DiegoAll FalcoxLab Docker Image
# 2022
#
# Build the image with:
#   $ docker build -t diegoall1990/falco-workshop .
#
# Start a container for the first time from a built image with:
#   $ docker run -it --name <name> -h <hostname> diegoall1990/falco-workshop:tagname
#   $ docker run -it --name falco-workshop -h training diegoall1990/falco-workshop:tagname

FROM ubuntu
MAINTAINER DiegoAll <dposadallano@hotmail.com>

# Let the container know there will be no TTY
ENV DEBIAN_FRONTEND=noninteractive
ENV TZ=America/Bogota DEBIAN_FRONTEND=noninteractive
RUN apt-get update && apt install tzdata -y && apt install net-tools nano file -y
RUN yes| unminimize
RUN apt-get update -y && apt-get upgrade -y && apt-get install ubuntu-minimal -y
RUN apt-get install -y wget tar sudo adduser netstat-nat net-tools curl
RUN apt-get install -y nmap pkexec golang-go python3
RUN touch /var/log/kern.log
RUN useradd -m -s /bin/bash falcox
# Security risk (Share)
RUN echo "root:toor" | sudo chpasswd
# Security risk NOPASSWD: ALL (Privilege escalation)
RUN usermod -aG sudo falcox && echo "ubuntu ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/falcox  # Revisar este ubuntu para docker sencillo
# Remember works in k8s but in docker: falcox is not in the sudoers file.
RUN chmod 044 /etc/sudoers.d/falcox

# apt install parted -y #(remove bulk data)
USER falcox:falcox

RUN mkdir /home/falcox/.ssh
RUN touch /home/falcox/.ssh/id_rsa
WORKDIR /home/falcox
CMD ["/bin/bash"]


# # Instalar utilidades (opcional)
# RUN apt-get update && apt-get install -y parted

# # Crear particiones (ejemplo)
# RUN parted /dev/sda mklabel gpt
# RUN parted /dev/sda mkpart primary ext4 0% 33%
# RUN parted /dev/sda mkpart primary ext4 33% 66%
# RUN parted /dev/sda mkpart primary ext4 66% 100%

# # Formatear particiones (ejemplo)
# RUN mkfs.ext4 /dev/sda1
# RUN mkfs.ext4 /dev/sda2
# RUN mkfs.ext4 /dev/sda3

# # Montar particiones (ejemplo)
# RUN mkdir /mnt/sda1
# RUN mkdir /mnt/sda2
# RUN mkdir /mnt/sda3
# RUN mount /dev/sda1 /mnt/sda1
# RUN mount /dev/sda2 /mnt/sda2
# RUN mount /dev/sda3 /mnt/sda3

# # Configurar /etc/fstab para montar particiones al inicio (opcional)
# RUN echo '/dev/sda1 /mnt/sda1 ext4 defaults 0 0' >> /etc/fstab
# RUN echo '/dev/sda2 /mnt/sda2 ext4 defaults 0 0' >> /etc/fstab
# RUN echo '/dev/sda3 /mnt/sda3 ext4 defaults 0 0' >> /etc/fstab
