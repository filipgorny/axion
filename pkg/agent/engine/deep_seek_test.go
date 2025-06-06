package engine

import "testing"

func TestDeepSeekEngine(t *testing.T) {
	engine, err := NewDeepSeekEngine()
	if err != nil {
		t.Fatalf("Failed to create DeepSeek engine: %v", err)
	}

	question := "What is the capital of France?, answer in one word"
	chat, err := engine.Chat()
	if err != nil {
		t.Fatalf("Failed to get response: %v", err)
	}

	sayMethod := chat.Say(question)

	message := <-chat.GetMessages()

	if message.Content != "Paris" {
		t.Errorf("Expected response to be 'Paris', got '%s'", message.Content)
	}

	sayMethod("What is the capital of Poland?, answer in one word")

	message = <-chat.GetMessages()

	if message.Content != "Warsaw" {
		t.Errorf("Expected response to be 'Warsaw', got '%s'", message.Content)
	}
}
