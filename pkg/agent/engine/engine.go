package engine

import "github.com/filipgorny/axion/pkg/agent/chat"

type Engine interface {
	Ask(question string) (string, error)
	Chat() (*chat.Chat, error)
}
