package work

import (
	"github.com/Palworld/internal/model"
	"github.com/Palworld/internal/repository/worksuitability"
)

type worksuitUsecase struct {
	worksuitRepo worksuitability.Repository
}

func GetUsecase(worksuitRepo worksuitability.Repository) Usecase {
	return &worksuitUsecase{
		worksuitRepo: worksuitRepo,
	}
}

func (ws *worksuitUsecase) GetWorkSuitabilityList() []model.WorkSuitability {
	return ws.worksuitRepo.GetWorkSuitabilityList()
}

func (ws *worksuitUsecase) GetWorkSuitability(workID string) model.WorkSuitability {
	return ws.worksuitRepo.GetWorkSuitability(workID)
}

func (ws *worksuitUsecase) UpdateWorkSuitability(workID string, name string) model.WorkSuitability {
	return ws.worksuitRepo.UpdateWorkSuitability(workID, name)
}

func (ws *worksuitUsecase) DeleteWorkSuitability(workID string) model.WorkSuitability {
	return ws.worksuitRepo.DeleteWorkSuitability(workID)
}

func (ws *worksuitUsecase) CreateWorkSuitability(works *model.WorkSuitability) []model.WorkSuitability {
	return ws.worksuitRepo.CreateWorkSuitability(works)
}
