package util

import (
	"context"
	"github.com/chromedp/chromedp"
)

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
