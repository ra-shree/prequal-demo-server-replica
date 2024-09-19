package middleware

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
)

var GlobalProbe Probe = *NewProbe("Server 1")

type (
	Probe struct {
		ServerName       string `json:"serverName"`
		RequestsInFlight uint64 `json:"requestInFlight"`
		Latency          uint64 `json:"latency"`
		mutex            sync.RWMutex
	}
)

func NewProbe(serverName string) *Probe {
	return &Probe{
		ServerName: serverName,
	}
}

func Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		startTime := time.Now()

		GlobalProbe.mutex.Lock()
		GlobalProbe.RequestsInFlight++
		GlobalProbe.mutex.Unlock()

		if err := next(c); err != nil {
			c.Error(err)
		}
		duration := uint64(time.Since(startTime).Milliseconds())
		GlobalProbe.mutex.Lock()
		defer GlobalProbe.mutex.Unlock()
		GlobalProbe.RequestsInFlight--
		GlobalProbe.Latency = (GlobalProbe.Latency + duration) / 2
		return nil
	}
}

func HandleGetPing(c echo.Context) error {
	fmt.Print("Was here once")
	GlobalProbe.mutex.RLock()
	defer GlobalProbe.mutex.RUnlock()
	return c.JSON(http.StatusOK, &GlobalProbe)
}
