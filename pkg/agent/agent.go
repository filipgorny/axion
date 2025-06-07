package agent

import (
	"github.com/filipgorny/axion/pkg/agent/action"
	"github.com/filipgorny/axion/pkg/agent/chat"
	"github.com/filipgorny/axion/pkg/agent/engine"
	"github.com/filipgorny/axion/pkg/project"
)

type Agent struct {
	engine  engine.Engine
	actions []action.Action
	project project.Project
}

func NewAgent(engine engine.Engine) *Agent {
	return &Agent{
		engine:  engine,
		actions: make([]action.Action, 0),
	}
}

func (a *Agent) Say(message string) chat.ChatFunction {
	chat, err := a.engine.Chat()

	if err != nil {
		panic("failed to create chat: " + err.Error())
	}

	return chat.Say(message)
}

func (a *Agent) RunLoop() {
}

func (a *Agent) AddAction(actions ...action.Action) {
	a.actions = append(a.actions, actions...)
}

func (a *Agent) AddProjectNodes(project *project.Project) {
	a.project = *project

	return func(actionType) {
		return action.Factory().(T)
	}
}
