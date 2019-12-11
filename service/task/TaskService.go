package s_task

import "apiproject/log"

type TaskService struct {
}

/**
执行任务1
*/
func (this *TaskService) Task1() {
	log.Logger.Info("执行任务1, 完成")
}
