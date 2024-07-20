package behaviors

import "fmt"

// Infection mode guarantees that 2 rules are triggered initially

// Download malware (rootkit)? persitence
// busca keys , envia keys

// Borra evidencia Validar que el script deje esta ejecucion en history

func RunVectorA() {
	fmt.Println("Ejecutando la función para vector A")

	// Llama a otras funciones necesarias aquí

	LaunchSuspiciousNetworkToolInContainer()

	fmt.Println("Done Launch external")

	// CreateFilesBelowDev()

	// SearchPrivateKeysOrPasswords()

	//ClearLogActivities()

	//SubFunctionA1()

	//SubFunctionA2()
}

func SubFunctionA1() {
	fmt.Println("SubFunctionA1 ejecutada")
}

func SubFunctionA2() {
	fmt.Println("SubFunctionA2 ejecutada")
}
