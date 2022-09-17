package server

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	remoAPICallCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "remo_api_call",
			Help: "Count of API Call for Nature Remo API",
		},
		[]string{"server"},
	)
	// sensorMetricLabels = []string{"sensor_name", "ip", "mac", "kind"}

	// sensorTemperatureGauge = promauto.NewGaugeVec(
	// 	prometheus.GaugeOpts{
	// 		Name: "sensors_temperature",
	// 		Help: "Temperature of a sensor in Celcius temperature.",
	// 	},
	// 	sensorMetricLabels,
	// )

	// sensorHumidityGauge = promauto.NewGaugeVec(
	// 	prometheus.GaugeOpts{
	// 		Name: "sensors_humidity",
	// 		Help: "Humidity of a sensor in %RH.",
	// 	},
	// 	sensorMetricLabels,
	// )

	// sensorHumidityGauge = promauto.NewGaugeVec(
	// 	prometheus.GaugeOpts{
	// 		Name: "sensors_humidity",
	// 		Help: "Humidity of a sensor in %RH.",
	// 	},
	// 	sensorMetricLabels,
	// )
)
