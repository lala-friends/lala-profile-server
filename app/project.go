package main

type Project struct {
	//PersonId         int    `json:"personId"`
	ProjectName      string `json:"projectName"`
	Period           string `json:"period"`
	PersonalRole     string `json:"personalRole"`
	MainOperator     string `json:"mainOperator"`
	ProjectSummary   string `json:"projectSummary"`
	Responsibilities string `json:"responsibilities"`
	UsedTechnology   string `json:"usedTechnology"`
	PrimaryRole      string `json:"primaryRole"`
	ProjectResult    string `json:"projectResult"`
	LinkedSite       string `json:"linkedSite"`
}
