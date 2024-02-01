package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) GetWorkSuitabilityList(c echo.Context) error {

	if worksuitData := h.worksuitUsecase.GetWorkSuitabilityList(); worksuitData != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data": worksuitData,
		})
	}

	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
		"data": "can't retrieve data",
	})
}

func (h *handler) GetWorkSuitability(c echo.Context) error {

	id := c.Param("workID")

	worksuitData := h.worksuitUsecase.GetWorkSuitability(id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": worksuitData,
	})
}
