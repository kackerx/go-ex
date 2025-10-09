package po

type User struct {
	BaseModel

	Mobile   string `gorm:"index:idx_mobile;unique;type:varchar(32);not null;default:''"`
	Password string `gorm:"type:varchar(128);not null;default:''"`
	Nickname string `gorm:"type:varchar(128);not null;default:''"`
	Gender   int    `gorm:"type:tinyint;not null;default:1;comment:'1男, 2女'"`
	Role     int    `gorm:"type:tinyint;not null;default:1;comment:'1普通, 2管理'"`
}

func (u User) TableName() string {
	return "t_mx_user"
}
