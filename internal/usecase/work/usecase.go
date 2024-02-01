package work

import "github.com/Palworld/internal/model"

type Usecase interface {
	GetWorkSuitabilityList() []model.WorkSuitability
	GetWorkSuitability(workID string) model.WorkSuitability
}
