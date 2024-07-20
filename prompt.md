Actualmente estoy descargando un archivo vuelta.txt que contiene un script en python que  se ejecuta a si mismo a traves de un pipe, este internamente ejecuta desde un repositorio de github un script de golang.

curl -sL http://34.27.180.215/sitio/vuelta.txt | python3

Aca esta el script en python:

import subprocess

def ejecutar_script_golang():
    try:
        # Ejecutar el comando de Golang
        resultado = subprocess.run(
            ['go', 'run', 'github.com/diego-all/falco-workshop@latest', 'init', '--vector', 'A'],
            capture_output=True,
            text=True,
            check=True
        )
        
        # Imprimir la salida del comando de Golang
        print("Salida estándar:")
        print(resultado.stdout)
        print("Salida de error:")
        print(resultado.stderr)
    
    except subprocess.CalledProcessError as e:
        print(f"Error al ejecutar el comando de Golang: {e}")
        print("Salida de error:")
        print(e.stderr)
    except Exception as e:
        print(f"Error inesperado: {e}")

if __name__ == "__main__":
    ejecutar_script_golang()

Quiero saber por que cuando ejecuto el script descargandolo con curl de un repositorio remoto se comporta diferente a ejecutarlo de forma local.

La salida de ejecutar el script de forma local es:

go run main.go init --vector A

root@pho3nix:/home/diegoall/FALCO/falco-workshop# go run main.go init --vector A
Ejecutando la función para vector A
Ejecutando: LaunchSuspiciousNetworkToolInContainer
nmap ya está instalado.
Resultado del comando 'nmap --version':
Versión de Nmap 7.80 ( https://nmap.org )
Plataforma: x86_64-pc-linux-gnu
Compilado con: liblua-5.3.6 openssl-3.0.2 nmap-libssh2-1.8.2 libz-1.2.11 libpcre-8.39 libpcap-1.10.1 nmap-libdnet-1.12 ipv6
Compilado conout:
Available nsock engines: epoll poll select

Pero cuando lo ejecuto con curl o directamente con me genera la siguiente salida y no logra ejecutar el nmap.
go run github.com/diego-all/falco-workshop@latest init --vector A

root@draios:~/sitio# go run github.com/diego-all/falco-workshop@latest init --vector A
Ejecutando la función para vector A
SubFunctionA1 ejecutada
SubFunctionA2 ejecutada
root@draios:~/sitio# curl -sL http://34.27.180.215/sitio/vuelta.txt | python3
Salida estándar:
Ejecutando la función para vector A
SubFunctionA1 ejecutada
SubFunctionA2 ejecutada

Salida de error:


Podrias darme la respuesta en español