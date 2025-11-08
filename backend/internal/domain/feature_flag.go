package domain
import (
	"time"
	"gorm.io/gorm"
)
type FeatureFlag struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"uniqueIndex;not null"`
	Description string         `json:"description" gorm:"type:text"`
	Enabled     bool           `json:"enabled" gorm:"default:false;not null"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}
func (FeatureFlag) TableName() string {
	return "feature_flags"
}
