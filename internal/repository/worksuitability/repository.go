package worksuitability

import "github.com/Palworld/internal/model"

type Repository interface {
	GetWorkSuitability(workID string) model.WorkSuitability
	GetWorkSuitabilityList() []model.WorkSuitability
	UpdateWorkSuitability(workID string, name string) model.WorkSuitability
	DeleteWorkSuitability(workID string) model.WorkSuitability
	CreateWorkSuitability(works *model.WorkSuitability) []model.WorkSuitability
}
