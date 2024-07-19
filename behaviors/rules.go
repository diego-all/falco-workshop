package behaviors

import "fmt"

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
		Regla1()
	case 2:
		Regla2()
	case 3:
		Regla3()
	default:
		fmt.Println("Opción no válida")
	}
}

func Regla1() {
	fmt.Println("Ejecutando Regla1")
	// Lógica para Regla1
}

func Regla2() {
	fmt.Println("Ejecutando Regla2")
	// Lógica para Regla2
}

func Regla3() {
	fmt.Println("Ejecutando Regla3")
	// Lógica para Regla3
}
