package models

type User struct {
	UserID           uint               `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	UserName         string             `gorm:"type:VARCHAR(50);UNIQUE;NOT NULL"`
	Password         string             `gorm:"type:VARCHAR(100);NOT NULL"`
	UserTypeID       uint               `sql:"NOT NULL"`
	RestaurantOffers []RestaurantOffers `gorm:"foreign_key:CreateUserID"`
	SiteOffers       []SiteOffers       `gorm:"foreign_key:CreateUserID"`
}
