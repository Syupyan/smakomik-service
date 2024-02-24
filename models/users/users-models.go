package users

type Users struct {
	IdUsers     int64  `gorm:"primaryKey" json:"id_users" form:"id_users"`
	Username string `gorm:"varchar(150)" json:"username" form:"username"`
	Password string `gorm:"varchar(50)" json:"password" form:"password"`
	AccessU  string `gorm:"varchar(50)" json:"access_u" form:"access_u"`
}
