package config

type Subscriptions struct {
	ServiceName string `json:"service_name" gorm:"column:service_name;not null"`
	Price       int    `json:"price" gorm:"column:price;not null"`
	UserID      string `json:"user_id" gorm:"column:user_id;type:uuid;not null"`
	StartDate   string `json:"start_date" gorm:"column:start_date;type:date;not null"`
}

type FilterRequest struct {
	ServiceName string `json:"service_name" gorm:"column:service_name;not null"`
	EndDate     string `json:"end_date"`
	UserID      string `json:"user_id" gorm:"column:user_id;type:uuid;not null"`
	StartDate   string `json:"start_date" gorm:"column:start_date;type:date;not null"`
}
