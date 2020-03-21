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

/**
执行任务2
*/
func (this *taskService) Task2() {
	log.Logger.Info("执行任务2, 完成")
}
