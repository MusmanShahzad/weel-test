package domain
import (
	"encoding/json"
	"time"
	"gorm.io/gorm"
)
type OrderStatus string
const (
	OrderStatusPending    OrderStatus = "pending"
	OrderStatusProcessing OrderStatus = "processing"
	OrderStatusCompleted  OrderStatus = "completed"
	OrderStatusCancelled  OrderStatus = "cancelled"
)
type DeliveryPreference string
const (
	DeliveryPreferenceInStore  DeliveryPreference = "IN_STORE"
	DeliveryPreferenceDelivery DeliveryPreference = "DELIVERY"
	DeliveryPreferenceCurbside DeliveryPreference = "CURBSIDE"
)
type AISuggestedProduct struct {
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
	Reason   string  `json:"reason,omitempty"`
}
type Order struct {
	ID                  uint               `json:"id" gorm:"primaryKey"`
	UserID              uint               `json:"user_id" gorm:"not null;index"`
	User                User               `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Summary             string             `json:"summary" gorm:"type:text;not null"`
	DeliveryPreference  DeliveryPreference `json:"delivery_preference" gorm:"type:varchar(20);not null"`
	DeliveryAddress     *string            `json:"delivery_address,omitempty" gorm:"type:text"`
	PostalCode          *string            `json:"postal_code,omitempty" gorm:"type:varchar(20)"`
	AISuggestedProducts *string            `json:"ai_suggested_products,omitempty" gorm:"type:jsonb"`
	Status              OrderStatus        `json:"status" gorm:"type:varchar(20);default:'pending';not null"`
	CreatedAt           time.Time          `json:"created_at"`
	UpdatedAt           time.Time          `json:"updated_at"`
	DeletedAt           gorm.DeletedAt     `json:"-" gorm:"index"`
}
func (o *Order) GetAISuggestedProducts() ([]AISuggestedProduct, error) {
	if o.AISuggestedProducts == nil || *o.AISuggestedProducts == "" {
		return []AISuggestedProduct{}, nil
	}
	var products []AISuggestedProduct
	err := json.Unmarshal([]byte(*o.AISuggestedProducts), &products)
	return products, err
}
func (o *Order) SetAISuggestedProducts(products []AISuggestedProduct) error {
	if len(products) == 0 {
		o.AISuggestedProducts = nil
		return nil
	}
	data, err := json.Marshal(products)
	if err != nil {
		return err
	}
	jsonStr := string(data)
	o.AISuggestedProducts = &jsonStr
	return nil
}
func (Order) TableName() string {
	return "orders"
}
