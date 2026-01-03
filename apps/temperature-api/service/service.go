package service

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"strconv"
	"temperature-api/model"
	"time"
)

type SensorType string

const (
	Temperature SensorType = "temperature"

	TemperatureMin = 15
	TemperatureMax = 28
)

type TemperatureInterface interface {
	GetTemperature(sensorID string, location string) (model.TemperatureResponse, error)
}

type TemperatureService struct {
}

func NewTemperatureService() *TemperatureService {
	return &TemperatureService{}
}

// GetTemperature Получение значения температуры
func (s *TemperatureService) GetTemperature(sensorID string, location string) (model.TemperatureResponse, error) {
	if sensorID == "" && location == "" {
		return model.TemperatureResponse{}, errors.New("sensor id or location is empty")
	}

	if location == "" {
		switch sensorID {
		case "1":
			location = "Living Room"
		case "2":
			location = "Bedroom"
		case "3":
			location = "Kitchen"
		default:
			location = "Unknown"
		}
	} else if sensorID == "" {
		switch location {
		case "Living Room":
			sensorID = "1"
		case "Bedroom":
			sensorID = "2"
		case "Kitchen":
			sensorID = "3"
		default:
			sensorID = "0"
		}
	}

	val := TemperatureMin + rand.Float64()*(TemperatureMax-TemperatureMin)
	result, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", val), 64)

	return model.TemperatureResponse{
		Value:       result,
		Unit:        "°C",
		Timestamp:   time.Now(),
		Location:    location,
		Status:      "active",
		SensorID:    sensorID,
		SensorType:  string(Temperature),
		Description: "Temperature in location",
	}, nil
}
