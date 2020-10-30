package model

/**
@author 王世彪
	个人博客: https://sofineday.com?from=apiproject
	微信: 645102170
	QQ: 645102170
*/

type Page struct {
	PageNo    int         `json:"pageNo"`
	PageSize  int         `json:"pageSize"`
	PageCount int         `json:"pageCount"`
	ItemCount int         `json:"itemCount"`
	ItemList  interface{} `json:"itemList"`
}
