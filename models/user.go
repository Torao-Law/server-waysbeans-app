package models

// User model struct
type User struct {
	ID        int             `json:"id" gorm:"primary_key:auto_increment"`
	Name      string          `json:"name" gorm:"type: varchar(255)"`
	Email     string          `json:"email" gorm:"type: varchar(255)"`
	Password  string          `json:"password" gorm:"type: varchar(255)"`
	Status    string          `json:"status"`
	ProfileID int             `json:"-"`
	Profile   ProfileResponse `json:"profile"`
}

type UsersProfileResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (UsersProfileResponse) TableName() string {
	return "users"
}
