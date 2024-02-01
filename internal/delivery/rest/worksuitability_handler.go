package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Palworld/internal/model"
	"github.com/labstack/echo/v4"
)

func (h *handler) GetWorkSuitabilityList(c echo.Context) error {

	if worksuitData := h.worksuitUsecase.GetWorkSuitabilityList(); len(worksuitData) > 0 {
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

	if worksuitData := h.worksuitUsecase.GetWorkSuitability(id); worksuitData.ID != "" {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data": worksuitData,
		})
	}
	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
		"data": "can't retrieve data",
	})
}

func (h *handler) UpdateWorkSuitability(c echo.Context) error {
	id := c.Param("workID")
	var req model.WorkSuitability
	err := json.NewDecoder(c.Request().Body).Decode(&req)

	if err != nil {
		fmt.Printf("Got error %s\n", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	if worksuitData := h.worksuitUsecase.UpdateWorkSuitability(id, req.Name); worksuitData.ID != "" {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data": worksuitData,
		})
	}

	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
		"data": "can't retrieve data",
	})
}

func (h *handler) DeleteWorkSuitability(c echo.Context) error {
	id := c.Param("workID")

	if worksuitData := h.worksuitUsecase.DeleteWorkSuitability(id); worksuitData.ID != "" {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data": worksuitData,
		})
	}
	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
		"data": "id " + id + " is not found",
	})
}

func (h *handler) CreateWorkSuitability(c echo.Context) error {
	var req model.WorkSuitability
	err := json.NewDecoder(c.Request().Body).Decode(&req)

	if err != nil {
		fmt.Printf("Got error %s\n", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	if worksuitData := h.worksuitUsecase.CreateWorkSuitability(&req); worksuitData != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data": worksuitData,
		})
	}

	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
		"data": "can't retrieve data",
	})
}
