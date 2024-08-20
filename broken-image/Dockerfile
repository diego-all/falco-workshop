# DiegoAll FalcoxLab Docker Image
# 2022
#
# Build the image with:
#   $ docker build -t diegoall1990/falco-workshop .
#
# Start a container for the first time from a built image with:
#   $ docker run -it --name <name> -h <hostname> diegoall1990/falco-workshop:tagname
#   $ docker run -it --name falco-workshop -h training diegoall1990/falco-workshop:tagname

FROM ubuntu:18.04
MAINTAINER DiegoAll <dposadallano@hotmail.com>

# Actualiza el sistema y elimina cualquier sudo existente
RUN apt-get update && \
    apt-get remove -y sudo && \
    apt-get install -y wget dpkg

# Copia el paquete .deb con la versión vulnerable de sudo al contenedor
COPY sudo_1.8.21p2-3ubuntu1.6_amd64.deb /tmp/sudo_vulnerable.deb

# Instala el paquete .deb de sudo
RUN dpkg -i /tmp/sudo_vulnerable.deb || apt-get install -f -y

# Verifica que la versión instalada de sudo sea la correcta
RUN sudo --version

# Limpia los archivos temporales
RUN rm /tmp/sudo_vulnerable.deb && apt-get clean

# Establece el comando por defecto
CMD ["/bin/bash"]


