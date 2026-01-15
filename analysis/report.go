package analysis

import (
	"challenge-ssl-labs/models"
	"fmt"
	"time"
)

func Summary(result *models.GeneralResp) {
	//Cambio de formato UNIX a fecha y hora legible
	start := time.UnixMilli(result.StartTime)
	test := time.UnixMilli(result.TestTime)

	fmt.Println("\n✓ ---------- Resumen de Reporte SSL ---------- ✓")
	fmt.Println("Host:", result.Host)
	fmt.Println("Puerto:", result.Port)
	fmt.Println("Estado:", result.Status)
	fmt.Println("Escaneo iniciado:", start.Format("2006-01-02 15:04:05"))
	fmt.Println("Escaneo terminado:", test.Format("2006-01-02 15:04:05"))
	fmt.Println("\nEndpoints Totales:", len(result.Endpoints))

	//Lista de endpoints encontrados
	for i, ep := range result.Endpoints {
		fmt.Printf("[%d] %s | Grado: %s | Progreso: %d%%\n", i+1, ep.IPAddress, ep.Grade, ep.Progress)
	}

	fmt.Println("------------------------------------------------")
}

func DetailsEndpoint(result *models.GeneralResp) {
	var epIndex int

	fmt.Print(" ➡︎ Endpoint a revisar: ")
	fmt.Scanln(&epIndex)

	fmt.Println("------------------------------------------------")
	if epIndex < 1 || epIndex > len(result.Endpoints) {
		fmt.Println(" × Endpoint inválido")
		return
	}

	ep := result.Endpoints[epIndex-1]
	//Cambio de duracion de milisegundos a minutos/segundos
	d := time.Duration(ep.Duration) * time.Millisecond

	fmt.Printf("\n● -------- Endpoint %d -------- ●\n", epIndex)
	fmt.Println("IP:", ep.IPAddress)
	fmt.Println("Nombre del Servidor: ", ep.ServerName)
	fmt.Println("Grado:", ep.Grade)
	fmt.Println("Advertencias detectadas:", ep.HasWarnings)
	fmt.Println("Config. Especial: ", ep.IsExceptional)
	fmt.Printf("Progreso: %d%%\n", ep.Progress)
	fmt.Println("Duración: ", d)

	fmt.Println("------------------------------------------------")
}

func AllEndpoints(result *models.GeneralResp) {

	for i, ep := range result.Endpoints {
		//Cambio de duracion de milisegundos a minutos/segundos
		d := time.Duration(ep.Duration) * time.Millisecond

		fmt.Printf("\n● -------- Endpoint %d -------- ●\n", i+1)
		fmt.Println("IP:", ep.IPAddress)
		fmt.Println("Nombre del Servidor: ", ep.ServerName)
		fmt.Println("Grado:", ep.Grade)
		fmt.Println("Advertencias detectadas:", ep.HasWarnings)
		fmt.Println("Config. Especial: ", ep.IsExceptional)
		fmt.Printf("Progreso: %d%%\n", ep.Progress)
		fmt.Println("Duración: ", d)

	}

}
