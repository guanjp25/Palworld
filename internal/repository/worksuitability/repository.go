package worksuitability

import "github.com/Palworld/internal/model"

type Repository interface {
	GetWorkSuitability(workID string) model.WorkSuitability
	GetWorkSuitabilityList() []model.WorkSuitability
}
