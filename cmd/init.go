package cmd

import (
	"fmt"
	"log"

	"github.com/diego-all/falco-workshop/behaviors"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Inicializa la aplicación",
	Long:  `Inicializa la aplicación con las configuraciones proporcionadas.`,
	Run: func(cmd *cobra.Command, args []string) {
		vector, err := cmd.Flags().GetString("vector")
		if err != nil {
			log.Fatalf("Error al obtener el vector: %v", err)
		}

		switch vector {
		case "A":
			behaviors.RunVectorA()
		case "B":
			behaviors.RunVectorB()
		default:
			fmt.Println("Por favor, especifica un vector válido: A o B")
		}
	},
}

func init() {
	initCmd.Flags().String("vector", "", "Vector type (A or B)")
}
