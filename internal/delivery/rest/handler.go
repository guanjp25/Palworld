package rest

import "github.com/Palworld/internal/usecase/work"

type handler struct {
	worksuitUsecase work.Usecase
}

func NewHandler(worksuitUsecase work.Usecase) *handler {
	return &handler{
		worksuitUsecase: worksuitUsecase,
	}
}
