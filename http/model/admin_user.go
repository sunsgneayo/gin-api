package model

type DpjAdmin struct {
	ID       uint `gorm:"primaryKey"`
	UserName string
	Password string
	//CreatedAt time.Time
	//UpdatedAt time.Time
}
