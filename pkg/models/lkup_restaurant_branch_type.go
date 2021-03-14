package models


type LkupRestaurantBranchType struct {
	RestaurantBranchTypeID   uint                `gorm:"PRIMARY_KEY;NOT NULL"`
	RestaurantBranchTypeName string              `gorm:"type:VARCHAR(20);UNIQUE;NOT NULL"`
	RestaurantBranchs        []RestaurantBranchs `gorm:"foreign_key:RestaurantBranchTypeID"`
}
