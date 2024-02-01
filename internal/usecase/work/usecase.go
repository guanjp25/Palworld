package work

import "github.com/Palworld/internal/model"

type Usecase interface {
	GetWorkSuitabilityList() []model.WorkSuitability
	GetWorkSuitability(workID string) model.WorkSuitability
	UpdateWorkSuitability(workID string, name string) model.WorkSuitability
	DeleteWorkSuitability(workID string) model.WorkSuitability
	CreateWorkSuitability(works *model.WorkSuitability) []model.WorkSuitability
}
