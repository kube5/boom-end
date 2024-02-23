package model

type GetPairDescReq struct {
	ChainId uint64 `json:"chain_id" form:"chain_id" binding:"required"`
	PairId  uint64 `json:"pair_id" form:"pair_id" binding:"omitempty"`
}

func (req GetPairDescReq) ErrMessages() map[string]string {
	return map[string]string{
		"ChainId.required": "chain_id is empty",
	}
}

type OpenTinderTradeReq struct {
	ChainId     uint64 `json:"chain_id" binding:"required"`
	PairId      uint64 `json:"pair_id" binding:"omitempty"`
	Long        bool   `json:"long" binding:"omitempty"`
	MarkedPrice string `json:"marked_price" binding:"required"`
}

func (req OpenTinderTradeReq) ErrMessages() map[string]string {
	return map[string]string{
		"ChainId.required":     "chain_id is empty",
		"MarkedPrice.required": "marked_price is empty",
	}
}

type CloseTinderTradeReq struct {
	ChainId uint64 `json:"chain_id" binding:"required"`
	OrderId string `json:"order_id" binding:"required"`
}

func (req CloseTinderTradeReq) ErrMessages() map[string]string {
	return map[string]string{
		"ChainId.required": "chain_id is empty",
		"OrderId.required": "order_id is empty",
	}
}

type SetTraderConfigReq struct {
	Margin   uint64 `json:"margin" form:"margin" binding:"required"`
	Leverage uint64 `json:"leverage" form:"leverage" binding:"required"`
}

func (req SetTraderConfigReq) ErrMessages() map[string]string {
	return map[string]string{
		"Margin.required":   "margin is empty",
		"Leverage.required": "leverage is empty",
	}
}

type ChainIdPageReq struct {
	ChainId  uint64 `json:"chain_id" form:"chain_id" binding:"required"`
	PageNo   uint64 `json:"page_no" form:"page_no" binding:"required,gt=0"`
	PageSize uint64 `json:"page_size" form:"page_size" binding:"required"`
}

func (req ChainIdPageReq) ErrMessages() map[string]string {
	return map[string]string{
		"ChainId.required":  "chain_id is empty",
		"PageNo.required":   "page_no is empty",
		"PageNo.gt":         "required page_no > 0",
		"PageSize.required": "page_size is empty",
	}
}

type OrderIdReq struct {
	ChainId uint64 `json:"chain_id" form:"chain_id" binding:"required"`
	OrderId string `json:"order_id" form:"order_id" binding:"required"`
}

func (req OrderIdReq) ErrMessages() map[string]string {
	return map[string]string{
		"ChainId.required": "chain_id is empty",
		"OrderId.required": "order_id is empty",
	}
}
