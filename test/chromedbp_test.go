package test

import (
	"apiproject/util"
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"testing"
)

/**
@author 王世彪
	个人博客: https://sofineday.com?from=apiproject
	微信: 645102170
	QQ: 645102170
*/

/**
测试headless浏览器 chromedp,爬取动态网页
*/
func TestChromedp(t *testing.T) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var title, content string
	if err := chromedp.Run(ctx,
		//设置网页url
		chromedp.Navigate("https://www.toutiao.com/a6713074848530170381/"),
		//抽取数据
		chromedp.OuterHTML("h1.article-title", &title),
		chromedp.OuterHTML("div.article-content", &content),
	); err != nil {
		panic(err)
	}

	//打印抽取出的数据
	fmt.Println(title)
	fmt.Println(content)
}

/**
测试GetDynamicPageHtmlContent
*/
func TestGetDynamicPageHtmlContent(t *testing.T) {
	htmlContent, err := util.GetDynamicPageHtmlContent("https://www.toutiao.com/a6713074848530170381/")
	if err != nil {
		panic(err)
	}

	fmt.Println(htmlContent)
}
