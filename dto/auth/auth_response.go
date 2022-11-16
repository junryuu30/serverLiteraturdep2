package authdto

type RegisterResponse struct {
	FullName string `gorm:"type: varchar(255)" json:"name"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Phone    string `gorm:"type: varchar(255)" json:"phone"`
	Address  string `gorm:"type: varchar(255)" json:"address"`
	Token    string `gorm:"type: varchar(255)" json:"token"`
}

type LoginResponse struct {
	ID       int    `gorm:"type: int" json:"id"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Password string `gorm:"type: varchar(255)" json:"password"`
	Token    string `gorm:"type: varchar(255)" json:"token"`
	Status   string `json:"status"`
	Gender   string `json:"gender" gorm:"type: varchar(255)"`
	Phone    string `json:"phone" gorm:"type: varchar(255)"`
	Address  string `json:"address" gorm:"type: varchar(255)"`
	Image    string `json:"image" gorm:"type:varchar(255)"`
}

// type CheckAuthResponse struct {
// 	Id       int    `gorm:"type: int" json:"id"`
// 	FullName string `gorm:"type: varchar(255)" json:"fullname"`
// 	Email    string `gorm:"type: varchar(255)" json:"email"`
// }

type CheckAuthResponse struct {
	Id       int    `gorm:"type: int" json:"id"`
	FullName string `gorm:"type: varchar(255)" json:"fullname"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Gender   string `json:"gender" gorm:"type: varchar(255)"`
	Phone    string `json:"phone" gorm:"type: varchar(255)"`
	Address  string `json:"address" gorm:"type: varchar(255)"`
	Image    string `json:"image" gorm:"type:varchar(255)"`
	Status   string `json:"status" gorm:"type:varchar(255)"`
}
