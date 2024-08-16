package models

type InvItem struct {
	InvItemID    int
	InvCode      string
	InvName      string
	InvPkgSize   string
	InvStockUnit string
}

type StockIssue struct {
	StockIssueID int
	Custommer    *CustomerInfo
	Warehous     WarehouseInfo
}

type CustomerInfo struct {
	CRMID   string
	CRMName string
	CRMSnT  string
}

type WarehouseInfo struct {
	WarehouseID   int
	WarehouseName string
}
