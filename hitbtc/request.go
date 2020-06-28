package hitbtc

import (
	"time"

	"github.com/ramezanius/crypex/hitbtc/helper"
)

const (
	// Exchange order sides
	Buy  = "buy"
	Sell = "sell"

	// Exchange order types
	Limit      = "limit"
	Market     = "market"
	StopLimit  = "stopLimit"
	StopMarket = "stopMarket"
)

// Symbol struct
type Symbol struct {
	ID    string `json:"id,required"`
	Base  string `json:"baseCurrency,required"`
	Quote string `json:"quoteCurrency,required"`

	TickSize             string `json:"tickSize,required"`
	FeeCurrency          string `json:"feeCurrency,required"`
	QuantityIncrement    string `json:"quantityIncrement,required"`
	TakeLiquidityRate    string `json:"takeLiquidityRate,required"`
	ProvideLiquidityRate string `json:"provideLiquidityRate,required"`
}

// Symbols struct
type Symbols []Symbol

// Get specific symbol
func (h *HitBTC) GetSymbol(symbol string) (response *Symbol, err error) {
	request := struct {
		Symbol string `json:"symbol,required"`
	}{Symbol: symbol}

	err = h.Request("getSymbol", &request, &response)
	if err != nil {
		return nil, err
	}

	return
}

// Get exchange symbols
func (h *HitBTC) GetSymbols() (response *Symbols, err error) {
	err = h.Request("getSymbols", nil, &response)
	if err != nil {
		return nil, err
	}

	return
}

// Balance struct
type Balance struct {
	Currency string `json:"currency,required"`

	Reserved  float64 `json:"reserved,string"`
	Available float64 `json:"available,string"`
}

// Balances struct
type Balances []Balance

// Get user balances on exchange [!Authenticate]
func (h *HitBTC) GetBalances() (response *Balances, err error) {
	err = h.Request("getTradingBalance", nil, &response)
	if err != nil {
		return nil, err
	}

	return
}

// New order request struct
type NewOrder struct {
	Side           string    `json:"side,required"`
	Type           string    `json:"type,required"`
	Price          float64   `json:"price,string"`
	Symbol         string    `json:"symbol,required"`
	Quantity       float64   `json:"quantity,string"`
	StopPrice      float64   `json:"stopPrice,required"`
	ExpireTime     time.Time `json:"expireTime,required"`
	TimeInForce    string    `json:"timeInForce,required"`
	OrderID        string    `json:"clientOrderId,required"`
	PostOnly       bool      `json:"postOnly,required"`
	StrictValidate bool      `json:"strictValidate,required"`
}

// Place a new order [!Authenticate]
func (h *HitBTC) NewOrder(request *NewOrder) (response *Report, err error) {
	if request.OrderID == "" {
		request.OrderID = helper.GenerateUUID()
	}

	err = h.Request("newOrder", &request, &response)
	if err != nil {
		return nil, err
	}

	return
}

// Cancel an order [!Authenticate]
func (h *HitBTC) CancelOrder(orderID string) (response *Report, err error) {
	request := struct {
		OrderID string `json:"clientOrderId,required"`
	}{OrderID: orderID}

	err = h.Request("cancelOrder", &request, &response)
	if err != nil {
		return nil, err
	}

	return
}

// Replace order request struct
type ReplaceOrder struct {
	Price          float64 `json:"price,string"`
	Quantity       float64 `json:"quantity,string"`
	OrderID        string  `json:"clientOrderId,required"`
	RequestOrderID string  `json:"requestClientId,required"`
	StrictValidate bool    `json:"strictValidate,required"`
}

// Replace a new order [!Authenticate]
func (h *HitBTC) ReplaceOrder(request *ReplaceOrder) (response *Report, err error) {
	if request.RequestOrderID == "" {
		request.RequestOrderID = helper.GenerateUUID()
	}

	err = h.Request("cancelReplaceOrder", &request, &response)
	if err != nil {
		return nil, err
	}

	return
}
