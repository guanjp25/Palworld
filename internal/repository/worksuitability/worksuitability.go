package worksuitability

import (
	"github.com/Palworld/internal/model"
)

type workSuitabilityRepo struct {
	ws []model.WorkSuitability
}

func GetRepository(ws []model.WorkSuitability) Repository {
	return &workSuitabilityRepo{
		ws: ws,
	}
}

func (ws *workSuitabilityRepo) GetWorkSuitability(workID string) model.WorkSuitability {
	var data model.WorkSuitability

	for i := 0; i < len(ws.ws); i++ {
		if workID == ws.ws[i].ID {
			data = ws.ws[i]
			return data
		}
	}

	return data
}

func (ws *workSuitabilityRepo) GetWorkSuitabilityList() []model.WorkSuitability {

	if list := ws.ws; list != nil {
		return list
	}

	return nil
}
