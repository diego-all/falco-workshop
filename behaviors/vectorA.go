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

	CreateFilesBelowDev()

	time.Sleep(3 * time.Second)

	SearchPrivateKeysOrPasswords()

	time.Sleep(3 * time.Second)

	ClearLogActivities()

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
