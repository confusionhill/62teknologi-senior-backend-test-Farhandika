package health

import (
	"github.com/labstack/echo/v4"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"net/http"
	"project-template/model/dto"
	"time"
)

type checkHealthDTOResponse struct {
	Cpu  cpuDTOResponse           `json:"cpu"`
	Vmem virtualMemoryDTOResponse `json:"virtual_memory"`
}

type cpuDTOResponse struct {
	Used  int       `json:"used"`
	Usage []float64 `json:"usage"`
}

type virtualMemoryDTOResponse struct {
	Total     uint64 `json:"total"`
	Available uint64 `json:"available"`
	Used      uint64 `json:"used"`
}

func CheckHealth(c echo.Context) error {
	v, err := mem.VirtualMemory()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Cannot Check Health")
	}
	cpuPercent, err := cpu.Percent(time.Second, true)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Cannot Check Health")
	}

	numCPU, err := cpu.Counts(true)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Cannot Check Health")
	}

	return c.JSON(200, dto.ResponseWrapper{
		Status: 200,
		Msg:    "HTTP Service Health",
		Data: checkHealthDTOResponse{
			Cpu: cpuDTOResponse{
				Used:  numCPU,
				Usage: cpuPercent,
			},
			Vmem: virtualMemoryDTOResponse{
				Total:     v.Total,
				Available: v.Available,
				Used:      v.Used,
			},
		},
	})
}
