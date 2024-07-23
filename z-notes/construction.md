
ClearLogsActivities (**O_SYNC**)

Adición de syscall.O_SYNC:
La bandera syscall.O_SYNC se asegura de que los cambios en el archivo se sincronicen inmediatamente con el almacenamiento, lo que puede ayudar a garantizar que el evento se registre de manera efectiva.
Combina os.O_WRONLY y os.O_TRUNC:
Estas banderas garantizan que el archivo se abra en modo de solo escritura y que se trunque a una longitud de cero, eliminando su contenido.


