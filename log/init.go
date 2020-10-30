package log

import (
	"apiproject/config"
	"go.uber.org/zap"
)

/**
@author 王世彪
	个人博客: https://sofineday.com?from=apiproject
	微信: 645102170
	QQ: 645102170
*/

var Logger *zap.Logger
var HookLogger *zap.Logger

/**
初始化logger
*/
func Init() {
	Logger = GetCommonLogger()
	if config.GlobalConfig.LogKafkaHookSwitch ||
		config.GlobalConfig.LogMongoHookSwitch {
		HookLogger = GetHookLogger()
	}

	Logger.Info("初始化logger, 完成")
}
