package config

type Subscriptions struct {
	ServiceName string `json:"service_name" gorm:"column:service_name;not null" example:"Yandex Plus"`
	Price       int    `json:"price" gorm:"column:price;not null" example:"400"`
	UserID      string `json:"user_id" gorm:"column:user_id;type:uuid;not null" example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
	StartDate   string `json:"start_date" gorm:"column:start_date;type:date;not null" example:"2025-07"`
}

type FilterRequest struct {
	ServiceName string `json:"service_name" gorm:"column:service_name;not null" example:"Yandex Plus"`
	EndDate     string `json:"end_date" example:"2025-08"`
	UserID      string `json:"user_id" gorm:"column:user_id;type:uuid;not null" example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
	StartDate   string `json:"start_date" gorm:"column:start_date;type:date;not null" example:"2024-07"`
}
