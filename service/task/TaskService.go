package s_task

import "apiproject/log"

var TaskService = &taskService{}

type taskService struct {
}

/**
执行任务1
*/
func (this *taskService) Task1() {
	log.Logger.Info("执行任务1, 完成")
}
