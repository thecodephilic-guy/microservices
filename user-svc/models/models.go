package models

type User struct {
	ID    uint   `json:"user_id,omitempty" gorm:"primaryKey;autoIncrement"`
	Name  string `json:"name" gorm:"size:100"`
	Email string `json:"email" gorm:"unique"`
}

type Response struct {
	Message     string      `json:"message"`
	Explanation string      `json:"explanation"`
	Data        interface{} `json:"data"`
}
