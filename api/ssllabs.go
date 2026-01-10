package api

import (
	"challenge-ssl-labs/models"
	"encoding/json"
	"io"       //Lectura de datos por streams
	"net/http" //Uso de requests HTTP
)

// Aqui se recibe el host y se devuelve el JSON formateado al usar models
func SSLReport(host string) (*models.GeneralResp, error) {
	//Conjunción para crear la URL completa
	url := "https://api.ssllabs.com/api/v2/analyze?host=" + host

	//Consulta ---------------------------------------------------------
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	//Cierra la conexion al salir del main
	defer resp.Body.Close()

	//Lectura de lo obtenido por la API  ---------------------------------------------------------
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result models.GeneralResp

	//Reescritura de JSON a estructura Go
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
