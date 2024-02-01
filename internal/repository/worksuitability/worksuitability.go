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

func (ws *workSuitabilityRepo) UpdateWorkSuitability(workID string, name string) model.WorkSuitability {

	for i := 0; i < len(ws.ws); i++ {
		if workID == ws.ws[i].ID {
			ws.ws[i].Name = name
			return ws.ws[i]
		}
	}

	return model.WorkSuitability{
		ID:   "",
		Name: "",
		Type: "",
	}
}

func (ws *workSuitabilityRepo) DeleteWorkSuitability(workID string) model.WorkSuitability {
	for i := 0; i < len(ws.ws); i++ {
		if workID == ws.ws[i].ID {
			data := ws.ws[i]
			ws.ws = append(ws.ws[:i], ws.ws[i+1:]...)
			return data
		}
	}
	return model.WorkSuitability{
		ID:   "",
		Name: "",
		Type: "",
	}
}

func (ws *workSuitabilityRepo) CreateWorkSuitability(works *model.WorkSuitability) []model.WorkSuitability {
	if works.ID != "" {
		ws.ws = append(ws.ws, *works)
		return ws.ws
	}

	return ws.ws
}
