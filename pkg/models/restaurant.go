package models


type Restaurant struct {
	RestaurantID      uint                `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	RestaurantName    string              `gorm:"type:VARCHAR(150);UNIQUE;NOT NULL"`
	RestaurantOffers  []RestaurantOffers  `gorm:"foreign_key:RestaurantID"`
	RestaurantBranchs []RestaurantBranchs `sql:"foreign_key:RestaurantID"`
}
