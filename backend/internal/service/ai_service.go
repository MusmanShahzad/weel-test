package service
import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"weel-backend/config"
	"weel-backend/internal/domain"
	"github.com/sashabaranov/go-openai"
)
type AIService interface {
	SuggestProducts(summary string, address *string) ([]domain.AISuggestedProduct, error)
}
type aiService struct {
	client *openai.Client
}
func NewAIService(cfg *config.Config) AIService {
	if cfg.OpenAI.Secret == "" {
		log.Println("⚠️  Warning: OPEN_AI_SECRET not set, AI service will return empty suggestions")
		log.Println("   To enable AI suggestions, add OPEN_AI_SECRET=sk-... to your .env file")
		return &aiService{client: nil}
	}
	log.Println("✅ OpenAI service initialized successfully")
	return &aiService{
		client: openai.NewClient(cfg.OpenAI.Secret),
	}
}
func (s *aiService) SuggestProducts(summary string, address *string) ([]domain.AISuggestedProduct, error) {
	if s.client == nil {
		log.Println("⚠️  AI client not initialized, returning empty suggestions")
		return []domain.AISuggestedProduct{}, nil
	}
	log.Printf("🤖 Getting AI suggestions from OpenAI for: %s", summary)
	addressContext := ""
	if address != nil && *address != "" {
		addressContext = fmt.Sprintf("\nDelivery Address: %s", *address)
	}
	prompt := fmt.Sprintf(`You are a professional pharmacy receptionist. Your role is to help customers with medicines and health-related products only. 
Based on the customer's request below, suggest appropriate medicines and health products. Consider:
1. Any specific medicines mentioned in the request
2. Diseases or symptoms mentioned
3. Location/address context (if provided) - consider local availability and common health needs in that area
4. Only suggest medicines, supplements, medical supplies, and health-related products
5. Do NOT suggest non-medical items like groceries, electronics, etc.
Customer Request: %s%s
Please respond ONLY with a valid JSON array of suggested products in this exact format:
[
  {
    "name": "Product Name",
    "quantity": 1,
    "price": 0.00,
    "reason": "Brief explanation why this product is suggested"
  }
]
Important:
- Return ONLY the JSON array, no other text
- Include 2-5 relevant products
- Use realistic prices (in USD)
- Be specific with product names (use actual medicine names if mentioned)
- If no medicines or health-related items are mentioned, return an empty array: []`, summary, addressContext)
	req := openai.ChatCompletionRequest{
		Model: "gpt-4o-mini",
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "You are a professional pharmacy receptionist. You only handle medicines and health-related products. Respond with valid JSON arrays only.",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
		Temperature: 0.7,
		MaxTokens:   500,
	}
	ctx := context.Background()
	resp, err := s.client.CreateChatCompletion(ctx, req)
	if err != nil {
		log.Printf("❌ Error calling OpenAI API: %v", err)
		return []domain.AISuggestedProduct{}, fmt.Errorf("failed to get AI suggestions: %w", err)
	}
	if len(resp.Choices) == 0 {
		log.Println("⚠️  OpenAI returned no choices")
		return []domain.AISuggestedProduct{}, nil
	}
	log.Printf("✅ OpenAI responded with %d choice(s)", len(resp.Choices))
	content := strings.TrimSpace(resp.Choices[0].Message.Content)
	content = strings.TrimPrefix(content, "```json")
	content = strings.TrimPrefix(content, "```")
	content = strings.TrimSuffix(content, "```")
	content = strings.TrimSpace(content)
	var products []domain.AISuggestedProduct
	if err := json.Unmarshal([]byte(content), &products); err != nil {
		log.Printf("❌ Error parsing AI response: %v", err)
		log.Printf("   Raw content: %s", content)
		return []domain.AISuggestedProduct{}, fmt.Errorf("failed to parse AI response: %w", err)
	}
	log.Printf("✅ AI suggested %d products", len(products))
	return products, nil
}
