package project

import (
	"github.com/mikerourke/kingpin-starter/internal/db"
	"github.com/mikerourke/kingpin-starter/pkg/task"
)

type Project struct {
	name  string
	tasks []*task.Task
	db    *db.Database
}

func New(name string) *Project {
	d := db.New("/some/path.db")

	p := &Project{
		name: name,
		db:   d,
	}
	d.AddRecord(p)

	return p
}

func FindWithName(name string) *Project {
	return &Project{
		name:  name,
		tasks: nil,
		db:    nil,
	}
}

func (p *Project) Name() string {
	return p.name
}

func (p *Project) Save() {
	// Do something that involves saving.
}

func (p *Project) Tasks() []*task.Task {
	return p.tasks
}

func (p *Project) AddTask(t *task.Task) {
	t.SetProjectName(p.name)
	p.tasks = append(p.tasks, t)
}
