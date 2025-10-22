package display

import (
	"time"

	"github.com/KrystofJan/tempus/internal/repository"
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func PrintTasks(task []repository.Task) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("ID", "Name", "StartTime", "Finished", "RecordedTime")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	// TODO: Calculate recorded time -> sum of entries
	for _, task := range task {
		startTime := time.Unix(task.StartTimestamp, 0)
		tbl.AddRow(task.ID, task.Name, startTime.Format(time.Kitchen), task.Finished, task.RecordedTime)
	}

	tbl.Print()

}
