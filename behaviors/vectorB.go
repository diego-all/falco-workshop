package behaviors

import (
	"fmt"
	"time"
)

// Puede ser algo externo con LaunchIngressRemoteFileCopyToolsInContainer()  SERIA REPETIDO

func RunVectorB() {
	fmt.Println("Ejecutando la función para vector B")

	// DirectoryTraversalMonitoredFileRead() //funciona
	//time.Sleep(3 * time.Second)

	MountLaunchedInPrivilegedContainer()
	time.Sleep(3 * time.Second)

	// 1. Presunto escape de contenedor
	//DetecteReleaseAgentFileContainerEscapes() // requiere privilegios
	//time.Sleep(3 * time.Second)

	// 2. SearchPrivateKeysOrPasswords() // Ajustar tarda bastante tiempo

	// Descarga de archivo con curl -o (No parece ser necesaria)
	//LaunchIngressRemoteFileCopyToolsInContainer()
	// time.Sleep(3 * time.Second)

	PolkitLocalPrivilegeEscalationVulnerability_CVE_2021_4034()
	time.Sleep(3 * time.Second)

	// CVE-2021-3156 (Baron samedit) - The vulnerability does not exist, therefore, it is a failed scenario.
	SudoPotentialPrivilegeEscalation()
	time.Sleep(3 * time.Second)

	// DirectoryTraversalMonitoredFileRead()
	//time.Sleep(3 * time.Second)
	// POdria ser otro pod con un webserver, que obtiene la ip tras un escaneo
	// La request es de este tipo: curl -X POST http://localhost:8080/execute -d 'cat ../../../../etc/passwd' -H 'Content-Type: text/plain'
	// Deberia armar esta request o apoyarme con alguna tool para llegar a esta request.
	// Dirbuster ??

	// va
	// RemoveBulkDatafromDisk()
	//time.Sleep(3 * time.Second)

	// RESTANTES
	// Launch Remote File Copy Tools in Container
	// Search Private Keys or Passwords
	// Detect crypto miners using the Stratum protocol
	// Packet socket created in container NO ESTA
	// Netcat Remote Code Execution in Container NO ESTA

}
