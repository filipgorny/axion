package action

import "github.com/filipgorny/axion/pkg/project"

type Fact string

type Answers struct {
	data map[string]string
}

func NewAnswers() Answers {
	return Answers{
		data: make(map[string]string),
	}
}

func (a *Answers) SetAnswer(key string, value string) {
	a.data[key] = value
}

func (a *Answers) Get(key string) string {
	if value, exists := a.data[key]; exists {
		return value
	}
	return ""
}

type Questions struct {
	key      string // Key for the question, can be used to identify the question
	Question string
	Ask      func(answer string) Fact
}

func NewQuestion(key string, question string) *Questions {
	return &Questions{
		key:      key, // Using the question itself as the key
		Question: question,
		Ask:      func(answer string) Fact { return Fact(answer) },
	}
}

type ActionFactory[T Action] struct {
	Questions []*Questions
	Node      T                // Node can be any type, allowing flexibility in the action's context
	Project   *project.Project // Project can be used to associate the action with a specific project
}

func NewActionFactory[T Action](project *project.Project) *ActionFactory[T] {
	return &ActionFactory[T]{
		Questions: make([]*Questions, 0),
		Project:   project,
	}
}

func PrepareAction[T Action](p *project.Project) T {
	actionFactory := NewActionFactory[T](p)

	actionNode := actionFactory.Node.Factory().(T)

	return actionNode
}

func (af *ActionFactory[T]) Question(key string, questionContent string) *ActionFactory[T] {
	af.Questions = append(af.Questions, NewQuestion(key, questionContent))

	return af
}

func (af *ActionFactory[T]) Build(func(answers Answers) T) T
