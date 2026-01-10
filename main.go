package main

import (
	"challenge-ssl-labs/analysis" //Analisis de resultados
	"challenge-ssl-labs/api"      //Obtencion de datos
	"challenge-ssl-labs/models"   //Forma de estructura de los datos
	"fmt"
	"net"     //Funciones de red
	"net/url" //Parsear URLs
	"regexp"  //Para validar patrones como formatos de host
	"strings" //Manipulación de cadenas
	"time"    //Manejo de tiempòs y duraciones
)

func main() {

	//Flujo del programa
	for {
		//Dominio a consultar dado por CLI
		var h string
		fmt.Print("Host por analizar (sin http/https): ")
		fmt.Scanln(&h)

		//Verifica si el host es válido
		host, err := verifyHost(h)
		if err != nil {
			fmt.Println(" × Error:", err)
			continue //Vuelve a pedir el host
		}

		//Verifica si el host existe
		ips, err := net.LookupHost(host)
		if err != nil || len(ips) == 0 {
			fmt.Println(" × Host no existe o no se pudo resolver")
			continue
		}

		//Analisis del host
		fmt.Println("Analizando host:", host)

		result, err := analizeHost(host)
		if err != nil {
			fmt.Println(" × Error:", err)
			continue //Vuelve a pedir el host
		}

		//Resumen informativo
		analysis.Summary(result)

		//Menu
		for {
			opt := menu()

			switch opt {
			case 1:
				analysis.DetailsEndpoint(result)
			case 2:
				goto AGAIN
			case 0:
				fmt.Println("Saliendo...")
				return
			default:
				fmt.Println(" × Opcion invalida")
			}
		}

	AGAIN:
	}

}

func analizeHost(host string) (*models.GeneralResp, error) {
	for {
		var err error

		//Se llama al package de la API y se obtienen los resultados
		result, err := api.SSLReport(host)

		if err != nil {
			fmt.Println(" × Error:", err)
			return nil, err
		}

		//Visualización según estado
		if result.Status == "READY" {
			return result, nil
		}

		fmt.Println("Escaneo en progreso...")

		//Ver progreso de escaneo por cada endpoint
		for _, ep := range result.Endpoints {
			if ep.Progress >= 0 {
				fmt.Printf(" - %s: %d%%\n", ep.IPAddress, ep.Progress)
			}
		}

		time.Sleep(10 * time.Second)
	}
}

func menu() int {
	var opt int
	fmt.Println("\n▼ ▼ ------------ Menú ------------ ▼ ▼")
	fmt.Println("1. Ver detalles de un endpoint")
	fmt.Println("2. Analizar otro host")
	fmt.Println("0. Salir")
	fmt.Println("Seleccione una opción: ")
	fmt.Scanln(&opt)
	return opt
}

func verifyHost(input string) (string, error) {
	input = strings.TrimSpace(input)

	//Esquema para usar url.Parse bien
	if !strings.HasPrefix(input, "http://") && !strings.HasPrefix(input, "https://") {
		input = "https://" + input
	}

	//Parseo de la URL
	url, err := url.Parse(input)
	if err != nil {
		return "", err
	}
	//Seleccionamos la parte que nos interesa
	host := url.Hostname()
	//
	validHost := regexp.MustCompile(`^([a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}$`)

	if !validHost.MatchString(host) {
		return "", fmt.Errorf(" × Host inválido!")
	}

	return host, nil
}
