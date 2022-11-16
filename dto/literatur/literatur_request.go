package literaturdto

type CreateLiteratureRequest struct {
	UserID             int    `json:"user_id" form:"user_id"`
	Title              string `json:"title" form:"title" gorm:"type: varchar(255)"`
	PublicationDate    string `json:"publication_date" form:"publication_date" gorm:"type:varchar(255)"`
	Pages              int    `json:"pages" form:"pages" gorm:"type: varchar(255)" `
	ISBN               string `json:"isbn" form:"isbn" gorm:"type:varchar(255)"`
	Author             string `json:"author" form:"author" gorm:"type: varchar(255)"`
	Attache            string `json:"attache" form:"attache" gorm:"type: varchar(255)"`
	Cover              string `json:"cover" form:"cover" gorm:"type:varchar(255)"`
	Statusverification string `json:"statusverification" form:"statusverification" gorm:"type:varchar(255)"`
}

type UpdateLiteraturRequest struct {
	Statusverification string `json:"statusverification"`
}
