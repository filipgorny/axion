package agent

import "github.com/tmc/langchaingo/prompts"

type Agent struct {
	initialPrompt string
}

func NewAgent() *Agent {
	agent := &Agent{
		initialPrompt: "You are an experienced software engineer with expertise in architecture. You can read the whole codebase and create a diagram using Nodes. You can also create a code from the diagram.",
	}

	system_message_prompt := prompts.NewSystemMessagePrompt(agent.initialPrompt)

	return agent
}
