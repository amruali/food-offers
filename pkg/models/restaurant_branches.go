package models

type RestaurantBranchs struct {
	RestaurantBranchID       uint   `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	RestaurantID             uint   `sql:"NOT NULL"`
	RestaurantBranchTypeID   uint   `sql:"NOT NULL"`
	RestaurantBranchLocation string `gorm:"type:VARCHAR(200);NOT NULL"`
}
