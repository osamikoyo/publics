package entity

type UserIndex struct {
	UserID uint    `bson:"user_id"`
	Topics []Topic `bson:"topics"`
}
