package entity

type Topic struct {
	ID          uint   `bson:"id"`
	TextExplain string `bson:"text_explain"`
	Desc        string `bson:"desc"`
}
