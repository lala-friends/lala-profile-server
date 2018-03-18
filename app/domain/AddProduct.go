package domain

type AddProduct struct {
	name 	  string `json:"name"`
	introduce string `json:"introduce"`
	techs 	[]string `json:"techs"`
	repColor  string `json:"repColor"`
	imageUrl  string `json:"imageUrl"`
	details []Detail `json:"details"`
} 

type Detail struct {
	title 		string `json:"title"`
	description string `json:"description"`
	imageUrl 	string `json:"imageUrl"`
	productId 	int    `json:"productId"`
}