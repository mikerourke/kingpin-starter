package main

import (
	"os"
	"path/filepath"

	"github.com/mikerourke/kingpin-starter/pkg/project"
	"github.com/mikerourke/kingpin-starter/pkg/task"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	app := kingpin.New(filepath.Base(os.Args[0]), "Manage projects and tasks")
	app.HelpFlag.Short('h')

	projCommand := app.Command("project", "Add a project").Alias("p")
	projNameFlag := projCommand.Flag("name",
		"Name of the project").Short('n').Required().String()

	taskCommand := app.Command("task", "Add a task").Alias("t")
	taskProjNameFlag := taskCommand.Flag("proj",
		"Name of the parent project").Short('p').Required().String()
	taskNameFlag := taskCommand.Flag("name",
		"Name of the task").Short('n').Required().String()
	taskSetWasStartedFlag := taskCommand.Flag("set", "Set if started").Bool()

	startTaskCommand := app.Command("start", "Start a task")
	startTaskNameFlag := startTaskCommand.Flag("name",
		"Name of the task").Short('n').Required().String()

	stopTaskCommand := app.Command("stop", "Start a task")
	stopTaskNameFlag := stopTaskCommand.Flag("name",
		"Name of the task").Short('n').Required().String()

	parsedCmd := kingpin.MustParse(app.Parse(os.Args[1:]))
	switch parsedCmd {
	case projCommand.FullCommand():
		p := project.New(*projNameFlag)
		p.Save()

	case taskCommand.FullCommand():
		t := task.New(*taskNameFlag)
		t.SetProjectName(*taskProjNameFlag)
		if *taskSetWasStartedFlag {
			t.SetWasStarted()
		}

	case startTaskCommand.FullCommand():
		t := task.FindWithName(*startTaskNameFlag)
		t.Start()

	case stopTaskCommand.FullCommand():
		t := task.FindWithName(*stopTaskNameFlag)
		t.Stop()
	}
}
