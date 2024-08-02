package behaviors

import (
	"fmt"
	"time"
)

// Puede ser algo externo con LaunchIngressRemoteFileCopyToolsInContainer()  SERIA REPETIDO

func RunVectorB() {
	fmt.Println("Ejecutando la función para vector B")

	// Descarga de archivo con curl -o
	LaunchIngressRemoteFileCopyToolsInContainer()

	time.Sleep(3 * time.Second)

	// POdria ser otro pod con un webserver, que obtiene la ip tras un escaneo
	// La request es de este tipo: curl -X POST http://localhost:8080/execute -d 'cat ../../../../etc/passwd' -H 'Content-Type: text/plain'
	// Deberia armar esta request o apoyarme con alguna tool para llegar a esta request.
	// Dirbuster ??
	DirectoryTraversalMonitoredFileRead()

	time.Sleep(3 * time.Second)

	// Va, es fallido por que no existe la vulnerabilidad
	SudoPotentialPrivilegeEscalation()

	// va
	RemoveBulkDatafromDisk()

	time.Sleep(3 * time.Second)

	// RESTANTES
	// Launch Remote File Copy Tools in Container
	// Search Private Keys or Passwords
	// Detect crypto miners using the Stratum protocol
	// Packet socket created in container NO ESTA
	// Netcat Remote Code Execution in Container NO ESTA

}
