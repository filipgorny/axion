package engine

import (
	"context"
	"log"
	"os"

	deepseek "github.com/cohesion-org/deepseek-go"
	dotenv "github.com/dsh2dsh/expx-dotenv"
	"github.com/filipgorny/axion/pkg/agent/chat"
)

type DeepSeekEngine struct {
	client *deepseek.Client
}

func NewDeepSeekEngine() (*DeepSeekEngine, error) {
	env := dotenv.New()
	if err := env.Load(); err != nil {
		log.Fatalf("error loading .env files: %v", err)
	}

	client := deepseek.NewClient(os.Getenv("DEEPSEEK_API_KEY"))

	return &DeepSeekEngine{
		client: client,
	}, nil
}

func (d *DeepSeekEngine) Chat() (*chat.Chat, error) {
	messages := []deepseek.ChatCompletionMessage{}

	ctx := context.Background()

	say := func(question string) string {
		messages = append(messages, deepseek.ChatCompletionMessage{
			Role:    deepseek.ChatMessageRoleUser,
			Content: question,
		})

		response, err := d.client.CreateChatCompletion(ctx, &deepseek.ChatCompletionRequest{
			Model:    deepseek.DeepSeekChat,
			Messages: messages,
		})

		if err != nil {
			log.Fatalf("error: %v", err)
		}

		responseMessage, err := deepseek.MapMessageToChatCompletionMessage(response.Choices[0].Message)

		if err != nil {
			log.Fatalf("Mapping to message failed: %v", err)
		}

		messages = append(messages, responseMessage)

		log.Printf("Response: %s", responseMessage.Content)

		return responseMessage.Content
	}

	return chat.NewChat(say), nil
}
