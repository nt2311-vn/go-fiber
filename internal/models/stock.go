package models

type InvItem struct {
	InvItemID    int
	InvCode      string
	InvName      string
	InvPkgSize   string
	InvStockUnit string
	InvUnit1     string
	InvRate1     float64
	InvUnit2     string
	InvRate2     float64
	InvUnit3     string
	InvRate3     float64
	RequestQty   uint32
}

type StockIssue struct {
	StockIssueID int
	Custommer    *CustomerInfo
	Warehouse    WarehouseInfo
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
