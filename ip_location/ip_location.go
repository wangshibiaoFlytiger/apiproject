package ip_location

import (
	"apiproject/config"
	"apiproject/log"
	"github.com/kayon/iploc"
	"go.uber.org/zap"
)

var locator *iploc.Locator

func Init() {
	var err error
	locator, err = iploc.Open(config.GlobalConfig.IplocationQqwryPath)
	if err != nil {
		log.Logger.Error("qqwry初始化异常", zap.Error(err))
	}
}

/**
查询IP所属位置
*/
func GetIpLocationString(ip string) string {
	detail := locator.Find(ip)

	ipLocationString := ""
	if detail != nil {
		ipLocationString = detail.String()
	}

	log.Logger.Info("查询IP所属位置", zap.Any("ip", ip), zap.Any("ipLocationString", ipLocationString))

	return ipLocationString
}
