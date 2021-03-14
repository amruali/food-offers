package models

import "time"

type SiteOffers struct {
	SiteOfferID            uint      `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL" json:"site_offer_id"`
	DiscountTypeID         uint      `sql:"NOT NULL" json:"discount_type_id"`
	SiteOfferDiscountValue uint      `sql:"NOT NULL" json:"site_offer_discount_value"`
	SiteOfferCode          string    `gorm:"type:VARCHAR(200);UNIQUE;NOT NULL" json:"site_offer_code"`
	Valid                  string    `gorm:"type:VARCHAR(1)" json:"valid"`
	CreateUserID           uint      `sql:"NOT NULL" json:"create_user_id"`
	StartTimeStamp         time.Time `json:"start_time_stamp"`
	EndTimestamp           time.Time `json:"end_timestamp"`
	CreateTimestamp        time.Time `json:"create_timestamp"`
	UpdateTimestamp        time.Time `json:"update_timestamp"`
}
