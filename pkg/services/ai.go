package services

import (
	"context"
	"fmt"
	"sort"

	"github.com/Arash-Afshar/pagoda-tailwindcss/config"
)

type AIClients map[string]AIClient

type AIClient interface {
	GenerateText(ctx context.Context, prompt string) (string, error)
}

var _ AIClient = &OllamaClient{}
var _ AIClient = &TestClient{}

func NewAIClient(ais []config.AIConfig) (AIClients, error) {
	aisMap := make(AIClients)
	for _, ai := range ais {
		switch ai.Name {
		case "ollama":
			aisMap[ai.Name] = NewOllamaClient(ai.Key, ai.URL)
		case "test":
			aisMap[ai.Name] = NewTestClient(ai.Key, ai.URL)
		default:
			return nil, fmt.Errorf("invalid ai name: %s", ai.Name)
		}
	}

	return aisMap, nil
}

func (a AIClients) GetClient(name string) AIClient {
	return a[name]
}

func (a AIClients) GetClientList() []string {
	keys := make([]string, 0, len(a))
	for k := range a {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// ----------------------------- Ollama Client -----------------------------
type OllamaClient struct {
}

func NewOllamaClient(key, url string) *OllamaClient {
	return &OllamaClient{}
}

func (o *OllamaClient) GenerateText(ctx context.Context, prompt string) (string, error) {
	return fmt.Sprintf("By Ollama\nPrompt: %s\nResponse: %s", prompt, "This is a test response"), nil
}

// ----------------------------- Test Client -----------------------------
type TestClient struct {
}

func NewTestClient(key, url string) *TestClient {
	return &TestClient{}
}

func (t *TestClient) GenerateText(ctx context.Context, prompt string) (string, error) {
	return fmt.Sprintf("By Test\nPrompt: %s\nResponse: %s", prompt, "This is a test response"), nil
}
