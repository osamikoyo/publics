package entity

type UpdateReq struct {
	ID     uint  `json:"id"`
	Entity Event `json:"entity"`
}
