package collectiondto

type CollectionResponse struct {
	ID          int `json:"id" gorm:"primary_key:auto_increment"`
	LiteraturID int `json:"literatur_id" gorm:"int"`
	UserID      int `json:"user_id" gorm:"int"`
}
