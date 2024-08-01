package behaviors

import (
	"fmt"
	"time"
)

// Infection mode guarantees that 2 rules are triggered initially

// Download malware (rootkit)? persitence
// busca keys , envia keys

// Borra evidencia Validar que el script deje esta ejecucion en history

func RunVectorA() {
	fmt.Println("Ejecutando la función para vector A")

	// Insmod carga el modulo en la memoria del nucleo del sistema.
	if err := LinuxKernelModuleInjection(); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Módulo del kernel inyectado exitosamente.")
	}
	time.Sleep(3 * time.Second)

	// Delete evidence in logs
	err := ClearLogActivities()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("El contenido del archivo fue eliminado exitosamente.")
	}

	LaunchSuspiciousNetworkToolInContainer()
	time.Sleep(3 * time.Second)

	// Trasladar a la imagen la creacion de la carpeta de llaves para que solo quede la lectura en logs
	ReadSSHInformation()
	time.Sleep(3 * time.Second)

	//SearchPrivateKeysOrPasswords() // Ajustar tarda bastante tiempo
	time.Sleep(3 * time.Second)

	RemoveBulkDatafromDisk()

	time.Sleep(3 * time.Second)

	LaunchIngressRemoteFileCopyToolsInContainer()

	time.Sleep(3 * time.Second)

	DirectoryTraversalMonitoredFileRead()

	time.Sleep(3 * time.Second)

	SudoPotentialPrivilegeEscalation()

	//SubFunctionA1()

	//SubFunctionA2()
}

func SubFunctionA1() {
	fmt.Println("SubFunctionA1 ejecutada")
}

func SubFunctionA2() {
	fmt.Println("SubFunctionA2 ejecutada")
}
