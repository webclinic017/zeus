package v1_load_simulator

import (
	"crypto/rand"
	"fmt"
	rand2 "math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

const (

	// RouteResponseSizeHeader Will default to KiB units if the units header is not set
	RouteResponseSizeHeader = "X-Sim-Response-Size"

	// RouteResponseSizeUnitsHeader only supporting KiB and MiB for now
	RouteResponseSizeUnitsHeader = "X-Sim-Response-Size-Units"

	// RouteResponseFormat only supporting bytes, string, and json for now. default is string
	RouteResponseFormat = "X-Sim-Response-Format"

	RouteResponseSuccessStatusCode = "X-Sim-Response-Success-Status-Code"
	RouteResponseFailureStatusCode = "X-Sim-Response-Failure-Status-Code"

	// RouteResponseFailurePercentage only supporting [0-100] in float for now
	// it will check the random number generated against this value and return the failure status code
	// if it is less than or equal to the failure percentage
	RouteResponseFailurePercentage = "X-Sim-Failure-Percentage"

	RouteResponseLatency = "X-Sim-Latency-Milliseconds"
)

func Routes(e *echo.Echo) *echo.Echo {
	// Routes
	e.GET("/health", Health)
	e.GET("/healthz", Health)

	e.PUT("/v1/load/bias", BiasLoadResponse)

	e.GET("/v1/load/simulate", SimulatedLoadResponse)
	e.POST("/v1/load/simulate", SimulatedLoadResponse)
	e.PUT("/v1/load/simulate", SimulatedLoadResponse)
	e.DELETE("/v1/load/simulate", SimulatedLoadResponse)

	return e
}

type Response struct {
	Message string `json:"message"`
}

func Health(c echo.Context) error {
	return c.String(http.StatusOK, "Healthy")
}

var FailureOffset = 0.0

func BiasLoadResponse(c echo.Context) error {
	failureRateStr := c.Request().Header.Get(RouteResponseFailurePercentage)
	if failureRateStr != "" {
		var err error
		failureRate, err := strconv.ParseFloat(failureRateStr, 64)
		if err != nil {
			log.Err(err).Msgf("SimulatedLoadResponse: strconv.ParseFloat")
			return c.JSON(http.StatusInternalServerError, nil)
		}
		FailureOffset = failureRate
	}
	return c.JSON(http.StatusOK, Response{
		Message: "OK",
	})
}

func SimulatedLoadResponse(c echo.Context) error {
	respSize := c.Request().Header.Get(RouteResponseSizeHeader)
	respSizeUnit := c.Request().Header.Get(RouteResponseSizeUnitsHeader)
	var respSizeNum int
	if respSize == "" {
		respSizeNum = 0
	} else {
		sz, err := strconv.Atoi(respSize)
		if err != nil {
			log.Err(err).Msgf("SimulatedLoadResponse: strconv.Atoi")
			err = nil
			respSizeNum = 0
		} else {
			respSizeNum = sz
		}
	}

	unitBytes := 1024
	switch respSizeUnit {
	case "B":
		unitBytes = 1
	case "KiB":
		unitBytes = 1024
	case "MiB":
		unitBytes = 1024 * 1024
	default:
		unitBytes = 1024
	}
	respUnitSizeTotal := respSizeNum * unitBytes
	latencyDelayMs := 0
	respLatency := c.Request().Header.Get(RouteResponseLatency)
	switch respLatency {
	case "":
		latencyDelayMs = 0
	default:
		var err error
		latencyDelayMs, err = strconv.Atoi(respLatency)
		if err != nil {
			log.Err(err).Msgf("SimulatedLoadResponse: strconv.Atoi")
			latencyDelayMs = 0
			err = nil
		}
	}
	if latencyDelayMs > 0 {
		time.Sleep(time.Duration(latencyDelayMs))
	}

	respStatusCode := http.StatusOK
	respStatusCodeStr := c.Request().Header.Get(RouteResponseSuccessStatusCode)
	if respStatusCodeStr != "" {
		var err error
		respStatusCode, err = strconv.Atoi(respStatusCodeStr)
		if err != nil {
			log.Err(err).Msgf("SimulatedLoadResponse: strconv.Atoi")
			respStatusCode = http.StatusOK
			err = nil
		}
	}
	failureStatusCode := http.StatusInternalServerError
	failureStatusCodeStr := c.Request().Header.Get(RouteResponseFailureStatusCode)
	if failureStatusCodeStr != "" {
		var err error
		failureStatusCode, err = strconv.Atoi(failureStatusCodeStr)
		if err != nil {
			log.Err(err).Msgf("SimulatedLoadResponse: strconv.Atoi")
			failureStatusCode = http.StatusInternalServerError
			err = nil
		}
	}
	failureRate := FailureOffset
	failureRateStr := c.Request().Header.Get(RouteResponseFailurePercentage)
	if failureRateStr != "" {
		var err error
		failureRate, err = strconv.ParseFloat(failureRateStr, 64)
		if err != nil {
			log.Err(err).Msgf("SimulatedLoadResponse: strconv.ParseFloat")
			failureRate = 0.0
			err = nil
		}
	}
	data := make([]byte, respUnitSizeTotal)
	if _, err := rand.Read(data); err != nil {
		log.Err(err).Msgf("SimulatedLoadResponse: rand.Read")
		return c.JSON(failureStatusCode, nil)
	}
	if failureRate > 0.0 {
		r := rand2.Float64() * 100.0
		if r <= failureRate {
			return c.JSON(failureStatusCode, Response{
				Message: fmt.Sprintf("Failure rate: %f: triggered", failureRate),
			})
		}
	}
	respFormat := c.Request().Header.Get(RouteResponseFormat)
	switch respFormat {
	case "bytes":
		return c.Blob(respStatusCode, "application/octet-stream", data)
	case "json":
		type RandomJSON struct {
			Data []byte `json:"data"`
		}
		randomObject := RandomJSON{
			Data: data,
		}
		return c.JSON(respStatusCode, randomObject)
	case "string":
		return c.String(respStatusCode, string(data))
	default:
		type RandomJSON struct {
			Data []byte `json:"data"`
		}
		data = make([]byte, 1*1024)
		if _, err := rand.Read(data); err != nil {
			log.Err(err).Msgf("SimulatedLoadResponse: rand.Read")
			return c.JSON(failureStatusCode, nil)
		}
		randomObject := RandomJSON{
			Data: data,
		}
		return c.JSON(respStatusCode, randomObject)
	}
}
