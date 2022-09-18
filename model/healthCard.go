package model

type HealthCard struct {
	ID       int    `json:"id"`
	Openid   string `gorm:"type:varchar(40);" json:"openid"`
	StuID    string `gorm:"type:varchar(20);" json:"stu_id"`
	Passwd   string `json:"passwd"`
	Location string `gorm:"type:varchar(1023);" json:"location"`
}
