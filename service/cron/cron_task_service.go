package s_cron

import (
	"apiproject/cron"
	"apiproject/dao"
	d_video "apiproject/dao/cron"
	"apiproject/log"
	m_cron "apiproject/model/cron"
	s_task "apiproject/service/task"
	"errors"
	"github.com/bwmarrin/snowflake"
	cron2 "github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

var CronTaskService = &cronTaskService{}

type cronTaskService struct {
}

/**
查询定时任务列表
*/
func (this *cronTaskService) FindCronTaskList() (cronTaskList []*m_cron.CronTask, err error) {
	cronTaskList = []*m_cron.CronTask{}
	if err := d_video.CronTaskDao.FindList(dao.Db, &cronTaskList); err != nil {
		log.Logger.Error("查询定时任务列表, 异常", zap.Error(err))
		return nil, err
	}

	return cronTaskList, nil
}

/**
加载定时任务列表: 需要系统启动时自动执行
*/
func (this *cronTaskService) LoadCronTaskList() (err error) {
	cronTaskList := []*m_cron.CronTask{}
	if err = d_video.CronTaskDao.FindList(dao.Db.Where("status = ?", 1), &cronTaskList); err != nil {
		log.Logger.Error("加载定时任务列表, 查询定时任务列表, 异常", zap.Any("err", err))
		return err
	}

	if len(cronTaskList) == 0 {
		log.Logger.Info("加载定时任务列表, 没有需要启动的定时任务")
		return
	}

	for _, cronTask := range cronTaskList {
		if err = this.LoadCronTask(cronTask); err != nil {
			log.Logger.Error("加载定时任务列表, 加载当前任务, 异常", zap.Any("cronTask", cronTask), zap.Error(err))
			return err
		}
	}

	log.Logger.Info("加载定时任务列表, 完成", zap.Any("count", len(cronTaskList)))
	return nil
}

/**
加载定时任务
*/
func (this *cronTaskService) LoadCronTask(cronTask *m_cron.CronTask) (err error) {
	//注册定时任务
	cronTask.EntryId, err = this.RegisterCronTask(cronTask)
	if err != nil {
		log.Logger.Error("加载定时任务, 注册异常", zap.Any("cronTask", cronTask), zap.Error(err))
		return err
	}

	//更新定时任务: 为了更新任务的entryId
	if err := d_video.CronTaskDao.Update(dao.Db.Model(cronTask).Where("id = ?", cronTask.ID), &cronTask); err != nil {
		log.Logger.Error("加载定时任务, 更新定时任务, 异常", zap.Any("cronTask", cronTask), zap.Error(err))
		return err
	}

	return nil
}

/**
启用定时任务
*/
func (this *cronTaskService) EnableCronTask(cronTaskId snowflake.ID) (err error) {
	cronTask := &m_cron.CronTask{}
	if err = d_video.CronTaskDao.Get(dao.Db.Where("id = ?", cronTaskId), cronTask); err != nil {
		log.Logger.Error("启用定时任务, 查询定时任务, 异常", zap.Any("cronTaskId", cronTaskId), zap.Error(err))
		return err
	}

	if cronTask.Status == 1 {
		log.Logger.Error("启用定时任务, 当前已启用", zap.Any("cronTask", cronTask))
		return errors.New("当前任务已启用, 不允许重复启用")
	}

	//注册定时任务
	entryId, err := this.RegisterCronTask(cronTask)
	if err != nil {
		log.Logger.Error("启用定时任务, 注册定时任务, 异常", zap.Any("cronTask", cronTask), zap.Error(err))
		return err
	}

	//更新数据库
	cronTask.Status = 1
	cronTask.EntryId = entryId
	if err = d_video.CronTaskDao.Update(dao.Db.Model(cronTask).Where("id = ?", cronTask.ID), cronTask); err != nil {
		log.Logger.Error("启用定时任务, 更新定时任务, 异常", zap.Any("cronTask", cronTask), zap.Error(err))
		return err
	}

	return nil
}

/**
禁用定时任务
*/
func (this *cronTaskService) DisableCronTask(cronTaskId snowflake.ID) (err error) {
	cronTask := &m_cron.CronTask{}
	if err = d_video.CronTaskDao.Get(dao.Db.Where("id = ?", cronTaskId), cronTask); err != nil {
		log.Logger.Error("禁用定时任务, 查询定时任务, 异常", zap.Any("cronTaskId", cronTaskId), zap.Error(err))
		return err
	}

	if cronTask.Status == 2 {
		log.Logger.Error("禁用定时任务, 当前已禁用", zap.Any("cronTask", cronTask))
		return errors.New("当前任务已禁用, 不允许重复禁用")
	}

	//反注册定时任务
	cron.UnRegisterTask(cronTask.EntryId)

	//更新数据库
	cronTask.Status = 2
	if err = d_video.CronTaskDao.Update(dao.Db.Model(cronTask).Where("id = ?", cronTask.ID), cronTask); err != nil {
		log.Logger.Error("启用定时任务, 更新定时任务, 异常", zap.Any("cronTask", cronTask), zap.Error(err))
		return err
	}

	return nil
}

/**
添加定时任务
*/
func (this *cronTaskService) AddCronTask(cronTask *m_cron.CronTask) (err error) {
	cronTask.EntryId, err = this.RegisterCronTask(cronTask)
	if err != nil {
		log.Logger.Error("添加定时任务, 异常", zap.Any("cronTask", cronTask), zap.Error(err))
		return err
	}

	//插入定时任务
	if err = d_video.CronTaskDao.Insert(dao.Db, cronTask); err != nil {
		log.Logger.Error("添加定时任务, 插入异常", zap.Any("cronTask", cronTask), zap.Error(err))
		return err
	}

	log.Logger.Info("添加定时任务, 完成", zap.Any("cronTask", cronTask))

	return nil
}

/**
删除定时任务
*/
func (this *cronTaskService) DeleteCronTask(cronTaskId snowflake.ID) (err error) {
	cronTask := &m_cron.CronTask{}
	if err = d_video.CronTaskDao.Get(dao.Db.Where("id = ?", cronTaskId), cronTask); err != nil {
		log.Logger.Error("删除定时任务, 查询异常", zap.Any("cronTaskId", cronTaskId), zap.Error(err))
		return err
	}

	//反注册定时任务
	cron.UnRegisterTask(cronTask.EntryId)

	//从数据库中删除定时任务
	if err = d_video.CronTaskDao.Delete(dao.Db.Where("id = ?", cronTaskId), cronTask); err != nil {
		log.Logger.Error("删除定时任务, 删除异常", zap.Any("cronTaskId", cronTaskId), zap.Error(err))
		return err
	}

	log.Logger.Info("删除定时任务, 完成", zap.Any("cronTask", cronTask))

	return nil
}

/**
注册定时任务
*/
func (this *cronTaskService) RegisterCronTask(cronTask *m_cron.CronTask) (entryId cron2.EntryID, err error) {
	if cronTask.Type == "task1" {
		//注册定时任务
		entryId, err = cron.RegisterTask(cronTask.Title, cronTask.Spec, s_task.TaskService.Task1, cronTask)
		if err != nil {
			log.Logger.Error("添加类型为task1的定时任务, 注册异常", zap.Any("cronTask", cronTask), zap.Error(err))
			return 0, err
		}
	}

	return entryId, nil
}
