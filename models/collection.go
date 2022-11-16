package models

type Collection struct {
	ID          int                  `json:"id" gorm:"primary_key:auto_increment"`
	User        UsersProfileResponse `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID      int                  `json:"user_id" gorm:"int"`
	LiteraturID int                  `json:"literatur_id" gorm:"int"`
	Literatur   LiteraturResponse    `json:"literatur"`
}

type CollectionResponse struct {
	ID          int `json:"id"`
	LiteraturID int
	Literatur   LiteraturResponse    `json:"literatur"`
	User        UsersProfileResponse `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID      int                  `json:"user_id" gorm:"int"`
}

func (CollectionResponse) TableName() string {
	return "collections"
}
