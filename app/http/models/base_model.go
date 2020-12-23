package models

type BaseModel struct {
	ID        int64 `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`
	CreatedAt int64 `gorm:"column:created_at;index" json:"created_at"`
	UpdatedAt int64 `gorm:"column:updated_at;index" json:"updated_at"`
}
