package models

type Literatur struct {
	ID                 int                  `json:"id" gorm:"primary_key:auto_increment"`
	Title              string               `json:"title" form:"title" gorm:"type: varchar(255)"`
	UserID             int                  `json:"user_id"`
	User               UsersProfileResponse `json:"user"`
	PublicationDate    string               `json:"publicationdate" form:"publication_date" gorm:"type: varchar(255)"`
	Pages              int                  `json:"pages" form:"pages" gorm:"type: varchar(255)"`
	ISBN               string               `json:"isbn" form:"isbn" gorm:"type: varchar(255)"`
	Author             string               `json:"author" form:"author" gorm:"type: varchar(255)"`
	Attache            string               `json:"attache" form:"attache" gorm:"type: varchar(255)"`
	Cover              string               `json:"cover" gorm:"type: varchar(255)"`
	Statusverification string               `json:"statusverification" form:"statusverification" gorm:"type: varchar(255)"`
}

type LiteraturResponse struct {
	ID                 int                  `json:"id"`
	Title              string               `json:"title"`
	UserID             int                  `json:"user_id"`
	User               UsersProfileResponse `json:"user"`
	PublicationDate    string               `json:"publication_date"`
	Pages              int                  `json:"pages"`
	ISBN               string               `json:"isbn"`
	Author             string               `json:"author"`
	Attache            string               `json:"attache"`
	Cover              string               `json:"cover"`
	Statusverification string               `json:"statusverification"`
}

type LiteraturUserResponse struct {
	ID              int                  `json:"id"`
	Title           string               `json:"title"`
	PublicationDate string               `json:"publication_date"`
	Pages           int                  `json:"pages"`
	ISBN            string               `json:"isbn"`
	Author          string               `json:"author"`
	Attache         string               `json:"attache"`
	UserID          int                  `json:"-"`
	User            UsersProfileResponse `json:"user"`
}

func (LiteraturResponse) TableName() string {
	return "literaturs"
}

func (LiteraturUserResponse) TableName() string {
	return "literaturs"
}
