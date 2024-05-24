package models

type Product struct {
	Id               int64            `json:"id"`
	Name             string           `json:"name"`
	Stock            int64            `json:"stock"`
	Price            float32          `json:"price"`
	StockDecreaseLog StockDecreaseLog `gorm:"foreignKey:ProductRefer"`
}
