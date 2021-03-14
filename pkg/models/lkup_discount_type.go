package models

type LkupDiscountType struct {
	DiscountTypeID   uint               `gorm:"PRIMARY_KEY;NOT NULL"`
	DiscountTypeName string             `gorm:"type:VARCHAR(20);UNIQUE;NOT NULL"`
	SiteOffers       []SiteOffers       `gorm:"foreign_key:DiscountTypeID"`
	RestaurantOffers []RestaurantOffers `gorm:"foreign_key:DiscountTypeID"`
}
