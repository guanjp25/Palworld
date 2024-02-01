package main

import (
	"fmt"

	"github.com/Palworld/internal/delivery/rest"
	"github.com/Palworld/internal/model"
	"github.com/Palworld/internal/model/constant"
	wRepo "github.com/Palworld/internal/repository/worksuitability"
	wUsecase "github.com/Palworld/internal/usecase/work"
	"github.com/labstack/echo/v4"
)

var ws = []model.WorkSuitability{
	{ID: "1", Name: "Farming", Type: constant.TypeGrass},
	{ID: "2", Name: "Lumbering", Type: constant.TypeNeutral},
	{ID: "3", Name: "Kindling", Type: constant.TypeFire},
}

func main() {
	e := echo.New()

	worksRepo := wRepo.GetRepository(ws)
	worksUsecase := wUsecase.GetUsecase(worksRepo)

	h := rest.NewHandler(worksUsecase)

	rest.LoadRoutes(e, h)

	fmt.Println("Hello, World!")
	e.Logger.Fatal(e.Start(":14045"))
}
