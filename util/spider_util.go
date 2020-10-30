package util

import (
	"context"
	"github.com/chromedp/chromedp"
	"golang.org/x/text/encoding/simplifiedchinese"
)

/**
@author 王世彪
	个人博客: https://sofineday.com?from=apiproject
	微信: 645102170
	QQ: 645102170
*/

/**
获取指定动态网页的html内容
*/
func GetDynamicPageHtmlContent(url string) (htmlContent string, err error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	if err := chromedp.Run(ctx,
		//设置网页url
		chromedp.Navigate(url),
		//抽取数据
		chromedp.OuterHTML("html", &htmlContent),
	); err != nil {
		return "", err
	}

	return htmlContent, nil
}

/**
中文乱码转成中文：当爬取的网页内容出现中文乱码时,需要调用此函数做下转换
*/
func DecodeToGBK(text string) (string, error) {

	dst := make([]byte, len(text)*2)
	tr := simplifiedchinese.GB18030.NewDecoder()
	nDst, _, err := tr.Transform(dst, []byte(text), true)
	if err != nil {
		return text, err
	}

	return string(dst[:nDst]), nil
}
