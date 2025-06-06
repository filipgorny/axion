package chat

type SayMethod func(question string)

type LLMSayFunction func(question string) string

type Chat struct {
	messages       chan Message
	sayFunction    SayMethod
	Speaker        Author
	llmSayFunction LLMSayFunction
	chatFunction   ChatFunction
	GetMessages    func() <-chan Message
}

func NewChat(sayToLLM LLMSayFunction) *Chat {
	chat := &Chat{
		messages:       make(chan Message, 1),
		llmSayFunction: sayToLLM,
		Speaker:        NewHumanParticipant("User"),
	}

	chat.sayFunction = func(question string) {
		chat.messages <- *NewMessage(sayToLLM(question), chat.Speaker)

		go func(question string) {
			chat.messages <- *NewMessage(sayToLLM(question), chat.Speaker)
		}(question)
	}

	chat.chatFunction = func(question string) ChatFunction {
		chat.sayFunction(question)

		return chat.chatFunction
	}

	chat.GetMessages = func() <-chan Message {
		return chat.messages
	}

	return chat
}

func (c *Chat) Say(question string) ChatFunction {
	return c.chatFunction(question)
}
