package db

import (
	"github.com/amrali/FOODOFFERS/pkg/models"
	"github.com/jinzhu/gorm"
)

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&models.LkupUserType{}, &models.User{}, &models.Restaurant{}, &models.LkupRestaurantBranchType{},
		&models.RestaurantBranchs{}, &models.RestaurantOffers{}, &models.SiteOffers{}, &models.LkupDiscountType{})
}

/*
func AssignForeignKeys(db *gorm.DB) {
	// Table Users
	db.Model(&models.User{}).AddForeignKey("user_type_id", "lkup_user_types(user_type_id)", "RESTRICT", "RESTRICT")

	// Restaurant Branches
	db.Model(&models.RestaurantBranchs{}).AddForeignKey("restaurant_id", "restaurants(restaurant_id)", "CASCADE", "CASCADE")
	db.Model(&models.RestaurantBranchs{}).AddForeignKey("restaurant_branch_type_id", "lkup_restaurant_branch_types(restaurant_branch_type_id)", "RESTRICT", "RESTRICT")

	// Restaurant Offers
	db.Model(&models.RestaurantOffers{}).AddForeignKey("restaurant_id", "restaurants(restaurant_id)", "CASCADE", "CASCADE")
	db.Model(&models.RestaurantOffers{}).AddForeignKey("discount_type_id", "lkup_discount_types(discount_type_id)", "RESTRICT", "RESTRICT")
	db.Model(&models.RestaurantOffers{}).AddForeignKey("create_user_id", "users(user_id)", "RESTRICT", "RESTRICT")

	// Site Offers
	db.Model(&models.SiteOffers{}).AddForeignKey("discount_type_id", "lkup_discount_types(discount_type_id)", "RESTRICT", "RESTRICT")
	db.Model(&models.SiteOffers{}).AddForeignKey("create_user_id", "users(user_id)", "RESTRICT", "RESTRICT")
}
*/
