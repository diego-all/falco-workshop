# Knowledge base Falco events



## Privileges


pkexec es una herramienta de autenticación que se utiliza en sistemas Linux para elevar los privilegios de un usuario regular a los de root. Esto es necesario cuando una aplicación necesita realizar acciones que requieren permisos de administrador, como instalar software, modificar archivos del sistema, etc.
¿Por qué los contenedores no suelen necesitar pkexec?

Aislamiento de procesos:

Los contenedores aíslan los procesos en un entorno separado, lo que significa que los procesos que se ejecutan dentro de un contenedor no tienen acceso directo al sistema host.
Esto limita la necesidad de elevar los privilegios, ya que los cambios realizados dentro del contenedor no afectan directamente al sistema host.
Usuario root dentro del contenedor:

Por defecto, los procesos dentro de un contenedor se ejecutan con los privilegios del usuario root. Esto permite que los procesos realicen todas las acciones necesarias dentro del contenedor sin la necesidad de elevar los privilegios.
Sin embargo, es una práctica recomendada ejecutar los procesos dentro del contenedor como usuarios no root para mayor seguridad.







