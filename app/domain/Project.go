package domain

import "database/sql"

type Project struct {
	//PersonId         int    `json:"personId"`
	ProjectName      string         `json:"projectName"`
	Period           sql.NullString `json:"period"`
	PersonalRole     sql.NullString `json:"personalRole"`
	MainOperator     sql.NullString `json:"mainOperator"`
	ProjectSummary   sql.NullString `json:"projectSummary"`
	Responsibilities sql.NullString `json:"responsibilities"`
	UsedTechnology   sql.NullString `json:"usedTechnology"`
	PrimaryRole      sql.NullString `json:"primaryRole"`
	ProjectResult    sql.NullString `json:"projectResult"`
	LinkedSite       sql.NullString `json:"linkedSite"`
}
