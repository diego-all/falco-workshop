package behaviors

import (
	"fmt"
	"time"
)

// Puede ser algo externo con LaunchIngressRemoteFileCopyToolsInContainer()  SERIA REPETIDO

func RunVectorB() {
	fmt.Println("Ejecutando la función para vector B")

	LaunchIngressRemoteFileCopyToolsInContainer()

	time.Sleep(3 * time.Second)

	DirectoryTraversalMonitoredFileRead()

	time.Sleep(3 * time.Second)

	SudoPotentialPrivilegeEscalation()

	RemoveBulkDatafromDisk()

	time.Sleep(3 * time.Second)

}
