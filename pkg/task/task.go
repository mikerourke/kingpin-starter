package task

import "time"

type Task struct {
	name        string
	wasStarted  bool
	startTime   time.Time
	stopTime    time.Time
	projectName string
}

func New(name string) *Task {
	return &Task{
		name:       name,
		wasStarted: false,
	}
}

func FindWithName(name string) *Task {
	return &Task{
		name:       name,
		wasStarted: false,
	}
}

func (t *Task) Start() {
	t.startTime = time.Now()
}

func (t *Task) Stop() {
	t.stopTime = time.Now()
}

func (t *Task) SetWasStarted() {
	t.wasStarted = true
}

func (t *Task) ProjectName() string {
	return t.projectName
}

func (t *Task) SetProjectName(name string) {
	t.projectName = name
}
