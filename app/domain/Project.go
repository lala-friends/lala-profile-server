package domain

import (
	"time"
)

type Project struct {
	//PersonId         int    `json:"personId"`
	ProjectName  string    `json:"projectName"`
	PeriodFrom   time.Time `json:"periodFrom"`
	PeriodTo     time.Time `json:"periodTo"`
	Introduce    string    `json:"introduce"`
	Description  string    `json:"description"`
	Techs        []string  `json:"techs"`
	PersonalRole string    `json:"personalRole"`
	Link         string    `json:"link"`
}
