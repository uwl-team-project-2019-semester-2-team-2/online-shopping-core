package search

import "github.com/uwl-team-project-2019-semester-2-team-2/online-shopping-core/model"

type Marshaller struct {
	Count int					`json:"count"`
	SearchProducts []model.Search		`json:"products"`
}
