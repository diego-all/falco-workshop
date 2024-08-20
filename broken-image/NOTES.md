# NOTES.md

    # Utiliza Ubuntu 18.04 como imagen base
    FROM ubuntu:18.04

    # Actualiza el sistema y elimina cualquier sudo existente
    RUN apt-get update && \
        apt-get remove -y sudo && \
        apt-get install -y wget dpkg

    # Copia el paquete .deb con la versión vulnerable de sudo al contenedor
    COPY sudo_1.8.21p2-3ubuntu1.6_amd64.deb /tmp/sudo_vulnerable.deb

    # Instala el paquete .deb de sudo
    RUN dpkg -i /tmp/sudo_vulnerable.deb || apt-get install -f -y

    # Crea un usuario no root
    RUN useradd -ms /bin/bash nonrootuser

    # Establece el usuario por defecto como nonrootuser
    USER nonrootuser

    # Establece el directorio de trabajo
    WORKDIR /home/nonrootuser

    # Verifica que la versión instalada de sudo sea la correcta
    RUN sudo --version

    # Limpia los archivos temporales
    RUN rm /tmp/sudo_vulnerable.deb && apt-get clean

    # Establece el comando por defecto
    CMD ["/bin/bash"]



    docker build -t ubuntu-vulnerable-sudo-nonroot .

    docker run -it -d --name baron-samedit -h broken-sudo ubuntu-vulnerable-sudo-nonroot:latest

    docker exec -it baron-samedit /bin/bash




    opcion2

    docker exec -it baron-samedit /bin/bash -c "useradd -ms /bin/bash pepe"

    docker exec -it -u pepe baron-samedit /bin/bash


    usermod -aG sudo nonrootuser

    id pepe


Configurar sudo para no pedir contraseña

    nonrootuser ALL=(ALL) NOPASSWD: ALL


Agregue a pepe a mano




