package entity

type GetReq struct{
	ID uint `json:"id"`
	Entity Event `json:"entity"`
}