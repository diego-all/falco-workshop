package behaviors

import (
	"fmt"
	"time"
)

//

func RunVectorB() {
	fmt.Println("Ejecutando la función para vector B")

	RemoveBulkDatafromDisk()

	time.Sleep(3 * time.Second)

	LaunchIngressRemoteFileCopyToolsInContainer()

	time.Sleep(3 * time.Second)

	DirectoryTraversalMonitoredFileRead()

	time.Sleep(3 * time.Second)

	SudoPotentialPrivilegeEscalation()

}
