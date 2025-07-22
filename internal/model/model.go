package models

type User struct {
	ID        int64  `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	TenantID  uint   `gorm:"uniqueIndex:idx_tenant_user"`
	UserCode  string `gorm:"uniqueIndex:idx_tenant_user"`
}

func (User) TableName() string {
	return "users"
}
