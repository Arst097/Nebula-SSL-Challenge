package main

import (
	"challenge-ssl-labs/analysis" //Analisis de resultados
	"challenge-ssl-labs/api"      //Obtencion de datos
	"challenge-ssl-labs/models"   //Forma de estructura de los datos
	"encoding/json"
	"fmt"
	"net"     //Funciones de red
	"net/url" //Parsear URLs
	"os"
	"regexp"  //Para validar patrones como formatos de host
	"strings" //Manipulación de cadenas
	"time"    //Manejo de tiempos y duraciones
)

func main() {

	//Flujo del programa
	for {
		//Dominio a consultar dado por CLI
		var h string
		fmt.Print("Host por analizar: ")
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
			case 3:
				readFromCache()
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

// Proceso de escaneo
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
			saveToCache(result, 1*time.Minute)
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
	fmt.Println("3. Revisar cache")
	fmt.Println("0. Salir")
	fmt.Println("Seleccione una opción: ")
	fmt.Scanln(&opt)
	return opt
}

// Validación del input
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

	validHost := regexp.MustCompile(`^([a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}$`)

	if !validHost.MatchString(host) {
		return "", fmt.Errorf(" × Host inválido!")
	}

	return host, nil
}

// Guardado en cache de escaneo
func saveToCache(resp *models.GeneralResp, ttl time.Duration) error {

	entry := models.CacheEntry{
		Data:       *resp,
		ExpiriesAt: time.Now().Add(ttl),
	}

	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	name := resp.Host

	return os.WriteFile("Cache/"+name, data, 0644)
}

// Lectura de cache
func readFromCache() error {
	var h string
	fmt.Println("Host del cache: ")
	fmt.Scanln(&h)

	hostCache, err := verifyHost(h)

	data, err := os.ReadFile("Cache/" + hostCache)
	if err != nil {
		return err
	}

	var entry models.CacheEntry

	if err := json.Unmarshal(data, &entry); err != nil {
		return err
	}

	if time.Now().After(entry.ExpiriesAt) {
		os.Remove("Cache/" + entry.Data.Host)
		return fmt.Errorf("Cache expirado")
	}

	analysis.Summary(&entry.Data)
	analysis.AllEndpoints(&entry.Data)

	return nil
}
