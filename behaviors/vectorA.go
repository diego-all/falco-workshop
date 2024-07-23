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

	// Llama a otras funciones necesarias aquí

	if err := LinuxKernelModuleInjection(); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Módulo del kernel inyectado exitosamente.")
	}

	CreateFilesBelowDev()

	time.Sleep(3 * time.Second)

	ReadSSHInformation()

	SearchPrivateKeysOrPasswords() // Ajustar tarda bastante tiempo

	time.Sleep(3 * time.Second)

	err := ClearLogActivities()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("El contenido del archivo fue eliminado exitosamente.")
	}

	// Esperar 3 segundos
	time.Sleep(3 * time.Second)

	LaunchSuspiciousNetworkToolInContainer()

	//SubFunctionA1()

	//SubFunctionA2()
}

func SubFunctionA1() {
	fmt.Println("SubFunctionA1 ejecutada")
}

func SubFunctionA2() {
	fmt.Println("SubFunctionA2 ejecutada")
}
