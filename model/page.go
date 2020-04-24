package model

type Page struct {
	PageNo    int         `json:"pageNo"`
	PageSize  int         `json:"pageSize"`
	PageCount int         `json:"pageCount"`
	ItemCount int         `json:"itemCount"`
	ItemList  interface{} `json:"itemList"`
}
