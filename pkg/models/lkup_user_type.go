package models

type LkupUserType struct {
	UserTypeID   uint   `gorm:"PRIMARY_KEY;NOT NULL"`
	UserTypeName string `gorm:"type:VARCHAR(20);UNIQUE;NOT NULL"`
	User         []User `gorm:"foreign_key:UserTypeID"`
}
