package model

type Page struct {
	PageNo    int         `form:"pageNo" json:"pageNo"`
	PageSize  int         `form:"pageSize" json:"pageSize"`
	PageCount int         `json:"pageCount"`
	ItemCount int         `json:"itemCount"`
	ItemList  interface{} `json:"itemList"`
}
