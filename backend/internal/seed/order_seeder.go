package seed
import (
	"encoding/json"
	"log"
	"weel-backend/internal/domain"
	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)
type OrderSeeder struct{}
func NewOrderSeeder() Seeder {
	return &OrderSeeder{}
}
func (s *OrderSeeder) Name() string {
	return "OrderSeeder"
}
func (s *OrderSeeder) Seed(db *gorm.DB) error {
	var count int64
	db.Model(&domain.Order{}).Count(&count)
	if count > 0 {
		log.Println("Orders already exist, skipping seed")
		return nil
	}
	var users []domain.User
	if err := db.Find(&users).Error; err != nil {
		return err
	}
	if len(users) == 0 {
		log.Println("No users found, skipping order seed")
		return nil
	}
	ordersCreated := 0
	for _, user := range users {
		numOrders := gofakeit.IntRange(1, 3)
		for i := 0; i < numOrders; i++ {
			summary := gofakeit.Sentence(gofakeit.IntRange(10, 30))
			preferences := []domain.DeliveryPreference{
				domain.DeliveryPreferenceInStore,
				domain.DeliveryPreferenceDelivery,
				domain.DeliveryPreferenceCurbside,
			}
			preference := preferences[gofakeit.IntRange(0, len(preferences)-1)]
			var deliveryAddress *string
			var postalCode *string
			if preference == domain.DeliveryPreferenceDelivery {
				street := gofakeit.Street()
				city := gofakeit.City()
				state := gofakeit.State()
				zipCode := gofakeit.Zip()
				addr := street + ", " + city + ", " + state + " " + zipCode
				deliveryAddress = &addr
				if gofakeit.Bool() {
					pc := gofakeit.Zip()
					postalCode = &pc
				}
			} else if preference == domain.DeliveryPreferenceCurbside {
				if gofakeit.Bool() {
					street := gofakeit.Street()
					city := gofakeit.City()
					state := gofakeit.State()
					zipCode := gofakeit.Zip()
					addr := street + ", " + city + ", " + state + " " + zipCode
					deliveryAddress = &addr
					if gofakeit.Bool() {
						pc := gofakeit.Zip()
						postalCode = &pc
					}
				}
			}
			statuses := []domain.OrderStatus{
				domain.OrderStatusPending,
				domain.OrderStatusProcessing,
				domain.OrderStatusCompleted,
				domain.OrderStatusCancelled,
			}
			var aiProducts *string
			if gofakeit.Bool() {
				numProducts := gofakeit.IntRange(2, 5)
				products := make([]domain.AISuggestedProduct, numProducts)
				for j := 0; j < numProducts; j++ {
					products[j] = domain.AISuggestedProduct{
						Name:     gofakeit.ProductName(),
						Quantity: gofakeit.IntRange(1, 5),
						Price:    gofakeit.Float64Range(5.0, 50.0),
						Reason:   gofakeit.Sentence(gofakeit.IntRange(5, 15)),
					}
				}
				productsJSON, _ := json.Marshal(products)
				jsonStr := string(productsJSON)
				aiProducts = &jsonStr
			}
			order := &domain.Order{
				UserID:              user.ID,
				Summary:             summary,
				DeliveryPreference:  preference,
				DeliveryAddress:     deliveryAddress,
				PostalCode:          postalCode,
				AISuggestedProducts: aiProducts,
				Status:              statuses[gofakeit.IntRange(0, len(statuses)-1)],
			}
			if err := db.Create(order).Error; err != nil {
				return err
			}
			ordersCreated++
		}
	}
	log.Printf("✅ Created %d orders", ordersCreated)
	return nil
}
