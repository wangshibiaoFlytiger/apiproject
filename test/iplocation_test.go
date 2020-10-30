package test

import (
	"github.com/kayon/iploc"
	"log"
	"testing"
)

/**
@author 王世彪
	个人博客: https://sofineday.com?from=apiproject
	微信: 645102170
	QQ: 645102170
*/

/**
测试IP归属地
*/
func TestIplocation(t *testing.T) {
	loc, err := iploc.Open("../ip_location/qqwry.utf8.dat")
	if err != nil {
		panic(err)
	}
	detail := loc.Find("8.8.8") // 补全为8.8.0.8, 参考 ping 工具
	log.Printf("IP:%s; 网段:%s - %s; %s\n", detail.IP, detail.Start, detail.End, detail)

	detail2 := loc.Find("8.8.3.1")
	log.Printf("%t %t\n", detail.In(detail2.IP.String()), detail.String() == detail2.String())

	// output
	// IP:8.8.0.8; 网段: 8.7.245.0 - 8.8.3.255; 美国 科罗拉多州布隆菲尔德市Level 3通信股份有限公司
	// true true

	detail = loc.Find("1.24.41.0")
	log.Println(detail.String())
	log.Println(detail.Country, detail.Province, detail.City, detail.County)

	// output
	// 内蒙古锡林郭勒盟苏尼特右旗 联通
	// 中国 内蒙古 锡林郭勒盟 苏尼特右旗
}
