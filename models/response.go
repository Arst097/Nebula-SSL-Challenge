package models

import "time"

//Estructuras para utilizar informacion del JSON obtenido por la API
type GeneralResp struct {
	Status    string     `json:"status"`
	Host      string     `json:"host"`
	Port      int        `json:"port"`
	Protocol  string     `json:"protocol"`
	StartTime int64      `json:"startTime"`
	TestTime  int64      `json:"testTime"`
	Endpoints []Endpoint `json:"endpoints"`
}

type Endpoint struct {
	IPAddress     string `json:"ipAddress"`
	ServerName    string `json:"serverName"`
	Grade         string `json:"grade"`
	HasWarnings   bool   `json:"hasWarnings"`
	IsExceptional bool   `json:"isExceptional"`
	Progress      int    `json:"progress"`
	Duration      int    `json:"duration"`
}

type CacheEntry struct {
	Data       GeneralResp
	ExpiriesAt time.Time
}
