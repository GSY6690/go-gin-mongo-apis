package models

type Address struct {
	City string `json:"city" bson:"city"`
	ID   int    `json:"ID" bson:"ID"`
}

type User struct {
	Name    string  `json:"name" bson:"name"`
	Age     int     `json:"age" bson:"age"`
	Address Address `json:"address" bson:"address"`
}
