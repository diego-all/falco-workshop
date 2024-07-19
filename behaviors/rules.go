package behaviors

import (
	"fmt"
	"os"
)

func ShowMenu() {
	fmt.Println("Seleccione una opción:")
	fmt.Println("1. Create files below dev")
	fmt.Println("2. Read ssh information")
	fmt.Println("3. Launch Ingress Remote File Copy Tools in Container")
	fmt.Println("4. Outbound Connection to C2 Servers")
	fmt.Println("5. Launch Remote File Copy Tools in Container")
	fmt.Println("6. Directory Traversal monitored file read")
	fmt.Println("7. Linux Kernel Module Injection")
	fmt.Println("8. Clear Log Activities")
	fmt.Println("9. Remove Bulk Data from Disk")
	fmt.Println("10. Clear Log Activities")
	fmt.Println("11. Search Private Keys or Passwords")
	fmt.Println("12. Detect crypto miners using the Stratum protocol")
	fmt.Println("13. XXX")
	fmt.Println("14. XXXX")
	fmt.Println("15. XXXXX")

	var choice int
	fmt.Print("Ingrese el número de la opción: ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		CreateFilesBelowDev()
	case 2:
		ReadSSHInformation()
	case 3:
		LaunchIngressRemoteFileCopyToolsInContainer()
	case 4:
		OutboundConnectionToC2Servers()
	case 5:
		LaunchRemoteFileCopyToolsInContainer()
	case 6:
		DirectoryTraversalMonitoredFileRead()
	case 7:
		LinuxKernelModuleInjection()
	case 8:
		ClearLogActivities()
	case 9:
		RemoveBulkDatafromDisk()
	case 10:
		DetectCryptoMinersUsingTheStratumProtocol()
	case 11:
		LaunchSuspiciousNetworkToolInContainer()
	case 12:
		NetcatRemoteCodeExecutionInContainer()
	default:
		fmt.Println("Opción no válida")
	}
}

func CreateFilesBelowDev() {
	fmt.Println("Ejecutando: Create files below dev")

	filePath := "/dev/randomrtk"

	// Crear el archivo
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer file.Close()

	// Escribir en el archivo
	content := "Este es un rootkit de ejemplo en /dev/randomrtk"
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error al escribir en el archivo:", err)
		return
	}

	fmt.Println("Archivo creado y contenido escrito con éxito")

}

func ReadSSHInformation() {
	fmt.Println("Ejecutando: Read ssh information")
	// Lógica para Regla2
}

func LaunchIngressRemoteFileCopyToolsInContainer() {
	fmt.Println("Ejecutando Regla3")
	// Lógica para Regla3
}

func OutboundConnectionToC2Servers() {
	fmt.Println("Ejecutando: Outbound Connection to C2 Servers")
	// Lógica para Regla3
}

func LaunchRemoteFileCopyToolsInContainer() {
	fmt.Println("Ejecutando: Launch Remote File Copy Tools in Container")
	// Lógica para Regla3
}

func DirectoryTraversalMonitoredFileRead() {
	fmt.Println("Ejecutando: Directory Traversal monitored file read")
	// Lógica para Regla3
}

func LinuxKernelModuleInjection() {
	fmt.Println("Ejecutando: Linux Kernel Module Injection")
	// Lógica para Regla3
}

func ClearLogActivities() {
	fmt.Println("Ejecutando: Clear Log Activities")
	// Lógica para Regla3
}

func RemoveBulkDatafromDisk() {
	fmt.Println("Ejecutando: Remove Bulk Data from Disk")
	// Lógica para Regla3
}

func DetectCryptoMinersUsingTheStratumProtocol() {
	fmt.Println("Ejecutando: Detect crypto miners using the Stratum protocol")
	// Lógica para Regla3
}

func LaunchSuspiciousNetworkToolInContainer() {
	fmt.Println("Ejecutando: Remove Bulk Data from Disk")
	// Lógica para Regla3
}

func NetcatRemoteCodeExecutionInContainer() {
	fmt.Println("Ejecutando: Detect crypto miners using the Stratum protocol")
	// Lógica para Regla3
}
