package models

import "time"

type RestaurantOffers struct {
	RestaurantOfferID            uint      `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL" json:"restaurant_offer_id"`
	RestaurantID                 uint      `sql:"NOT NULL" json:"restaurant_id"`
	DiscountTypeID               uint      `sql:"NOT NULL" json:"discount_type_id"`
	RestaurantOfferDiscountValue uint      `sql:"NOT NULL" json:"restaurant_offer_discount_value"`
	Valid                        string    `gorm:"type:VARCHAR(1) json:"valid"`
	RestaurantOfferCode          string    `gorm:"type:VARCHAR(200);UNIQUE;NOT NULL"  json:"restaurant_offer_code"`
	CreateUserID                 uint      `sql:"NOT NULL" json:"create_user_id"`
	StartTimeStamp               time.Time `json:"start_timestamp"`
	EndTimestamp                 time.Time `json:"end_timestamp"`
	CreateTimestamp              time.Time `json:"create_timestamp"`
	UpdateTimestamp              time.Time `json:"update_timestamp"`
}
