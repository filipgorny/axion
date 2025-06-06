package chat

type Participant struct {
	Name string
}

func NewHumanParticipant(name string) *Participant {
	return &Participant{
		Name: name,
	}
}

func (p *Participant) isBot() bool {
	return false
}

type Bot struct {
	Name string
}

func NewBotParticipant(name string) *Bot {
	return &Bot{
		Name: name,
	}
}

func (b *Bot) isBot() bool {
	return true
}

type Author interface {
	isBot() bool
}
