package models

type PersonalAccessToken struct {
	ID             int    `json:"id"`
	UserID         int    `json:"user_id"`
	User           User   `gorm:"foreignKey:UserID"`
	Token          string `json:"token"`
	ExpiryDate     string `json:"expiry_date"`
	LastAccessedOn string `json:"last_accessed_on"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}
