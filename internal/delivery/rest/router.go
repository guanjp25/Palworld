package rest

import "github.com/labstack/echo/v4"

func LoadRoutes(e *echo.Echo, handler *handler) {
	e.GET("/work-suitability-list", handler.GetWorkSuitabilityList)
	e.GET("/work-suitability-list/:workID", handler.GetWorkSuitability)
	e.POST("/work-suitability-update/:workID", handler.UpdateWorkSuitability)
	e.POST("/work-suitability-create", handler.CreateWorkSuitability)
	e.DELETE("/work-suitability-delete/:workID", handler.DeleteWorkSuitability)
}
