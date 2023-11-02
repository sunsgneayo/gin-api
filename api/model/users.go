package model

type DpjUsers struct {
	ID            int `gorm:"primaryKey"`
	OpenId        string
	Nickname      string
	HeadPicture   string
	OnlineStatus  int
	Status        int
	Balance       int
	CreateTime    string
	LastLoginTime string
	InRoomId      int
	MatchNumber   int
	Region        string
}
