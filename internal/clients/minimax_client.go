package clients

import (
	"context"
	"errors"

	"github.com/sashabaranov/go-openai"
)

type AIClient interface {
	ProcessInput(ctx context.Context, systemPrompt string, userInput string) (string, error)
}

type minimaxClient struct {
	client *openai.Client
}

func NewMiniMaxClient(apiKey string) AIClient {
	cfg := openai.DefaultConfig(apiKey)
	cfg.BaseURL = "https://api.minimax.io/v1"

	return &minimaxClient{
		client: openai.NewClientWithConfig(cfg),
	}
}

func (m *minimaxClient) ProcessInput(ctx context.Context, systemPrompt string, userInput string) (string, error) {
	set := openai.ChatCompletionRequest{
		Model: "MiniMax-M3.1-Flash-Preview",
		Messages: []openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleSystem, Content: systemPrompt},
			{Role: openai.ChatMessageRoleUser, Content: userInput},
		},
		Temperature: 0.2,
	}
	resp, err := m.client.CreateChatCompletion(
		ctx,
		set,
	)
	if err != nil {
		return "", err
	}

	if len(resp.Choices) == 0 {
		return "", errors.New("El modelo no fue capaz de procesar la respuesta")
	}

	return resp.Choices[0].Message.Content, nil
}
