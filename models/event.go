package models

type Event struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Action   string `json:"action" binding:"required"`
	Origin   string `json:"origin" binding:"required"`
	Date     string `json:"date"`
	Metadata string `json:"metadata" gorm:"type:jsonb"` // JSONB para PostgreSQL
}
