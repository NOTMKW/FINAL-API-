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

type CreateUserRequest struct {
	Firstname string `json:"Firstname" validate:"required"`
	Lastname  string `json:"Lastname" validate:"required"`
	Email     string `json:"Email" validate:"required,email"`
	Password  string `json:"Password" validate:"required,min=6"`
}
