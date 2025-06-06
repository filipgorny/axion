package tests

import (
	"testing"

	"github.com/filipgorny/axion/pkg/agent/engine"
)

func TestChatDeepSeek(t *testing.T) {
	deepSeek, err := engine.NewDeepSeekEngine()

	if err != nil {
		t.Fatalf("Failed to create DeepSeek engine: %v", err)
	}

	chat, err := deepSeek.Chat()

	if err != nil {
		t.Fatalf("Failed to chat with DeepSeek: %v", err)
	}

	if chat == nil {
		t.Fatal("Chat response is nil")
	}

	chat.Say("Hello, DeepSeek!")("What is the capital of France?")("What is the capital of Germany?")
}
