package models

type StockDecreaseLog struct {
	Id           int64 `json:"id"`
	OrderId      int64 `json:"order_id"`
	ProductRefer int64 `json:"product_refer"`
}
