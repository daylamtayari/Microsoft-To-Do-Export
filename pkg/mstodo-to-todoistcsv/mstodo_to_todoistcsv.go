package mstodo_to_todoistcsv

import (
	"github.com/daylamtayari/Microsoft-To-Do-Export/pkg/mstodo"
	"github.com/daylamtayari/Microsoft-To-Do-Export/pkg/todoistcsv"
)

func MSToDoToTodoistCsv(taskLists []mstodo.List) *todoistcsv.Export {
	export := todoistcsv.CreateExport(todoistcsv.ListLayout, nil)

	// Convert lists
	for i := range taskLists {
		section := todoistcsv.Section{}
		section.Title = taskLists[i].DisplayName
		for j := range taskLists[i].Tasks {
			section.AddTask(convertTask(taskLists[i].Tasks[j]))
		}
		export.AddSection(section)
	}

	return export
}

func convertTask(msTask mstodo.Task) todoistcsv.Task {
	task := todoistcsv.CreateTask(msTask.Title, msTask.Body.Content, todoistcsv.Priority1, todoistcsv.Indent1, nil)
	if msTask.DueDateTime != nil {
		task.Date = (*todoistcsv.Date)(&msTask.DueDateTime.Datetime.Time)
	}

	for _, cTask := range msTask.ChecklistItems {
		childTask := todoistcsv.CreateTask(cTask.DisplayName, "", todoistcsv.Priority1, todoistcsv.Indent2, nil)
		task.ChildTasks = append(task.ChildTasks, childTask)
	}

	return task
}
