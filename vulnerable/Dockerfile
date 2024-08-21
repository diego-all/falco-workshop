FROM ubuntu:18.04
MAINTAINER DiegoAll <dposadallano@hotmail.com>

# Actualiza el sistema y elimina cualquier sudo y nano existentes
RUN apt-get update && \
    apt-get remove -y sudo nano && \
    apt-get install -y wget dpkg

# Copia los paquetes .deb al contenedor
COPY sudo_1.8.21p2-3ubuntu1_amd64.deb /tmp/sudo_vulnerable.deb
COPY nano_2.9.3-2_amd64.deb /tmp/nano_old.deb

# Instala el paquete .deb de sudo
RUN dpkg -i /tmp/sudo_vulnerable.deb || apt-get install -f -y

# Instala el paquete .deb de nano
RUN dpkg -i /tmp/nano_old.deb || apt-get install -f -y

# Crea el usuario falcox con contraseña falcox
RUN useradd -m -s /bin/bash falcox && \
    echo "falcox:falcox" | chpasswd

# Agrega el usuario falcox al grupo sudo y configura NOPASSWD
RUN usermod -aG sudo falcox && \
    echo "falcox ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/falcox

# Verifica que la versión instalada de sudo sea la correcta
RUN sudo --version

# Limpia los archivos temporales
RUN rm /tmp/sudo_vulnerable.deb /tmp/nano_old.deb && apt-get clean

# Establece el comando por defecto
CMD ["/bin/bash"]
