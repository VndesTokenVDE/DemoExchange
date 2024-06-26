package entities

import (
	"github.com/google/uuid"
)

const (
	DefaultMarginType MarginType       = "isolated"
	DefaultLeverage   PositionLeverage = 10
)

type Position struct {
	AccountUID       AccountUID       `json:"account_uid" db:"account_uid"`
	PositionUID      string           `json:"position_uid" db:"position_uid"`
	Exchange         Exchange         `json:"exchange" db:"exchange"`
	Symbol           Symbol           `json:"symbol" db:"symbol"`
	Mode             PositionMode     `json:"mode" db:"position_mode"`
	MarginType       MarginType       `json:"type" db:"position_mode"`
	Leverage         PositionLeverage `json:"leverage" db:"leverage"`
	Side             PositionSide     `json:"side" db:"side"`
	Amount           float64          `json:"amount" db:"amount"`
	Price            float64          `json:"price" db:"price"`
	MarkPrice        float64          `json:"mark_price" db:"-"`
	Margin           float64          `json:"margin" db:"margin"`
	HoldAmount       float64          `json:"-" db:"hold_amount"`
	CreateTS         int64            `json:"create_ts" db:"create_ts"`
	UpdateTS         int64            `json:"update_ts" db:"update_ts"`
	IsNew            bool             `json:"-" db:"-"`
	MarginBalance    float64          `json:"margin_balance" db:"-"`
	UnrealisedPnl    float64          `json:"unrealised_pnl" db:"-"`
	LiquidationPrice float64          `json:"liquidation_price" db:"-"`
}

type PositionMode string

const (
	PositionModeOneway PositionMode = "oneway"
	PositionModeHedge  PositionMode = "hedge"
)

type MarginType string

const (
	MarginTypeIsolated MarginType = "isolated"
	MarginTypeCross    MarginType = "cross"
)

type PositionLeverage int32

func (l PositionLeverage) ToFloat64() float64 {
	return float64(l)
}

type PositionSide string

const (
	PositionSideLong  PositionSide = "long"
	PositionSideShort PositionSide = "short"
	PositionSideBoth  PositionSide = "both"
)

func (s PositionSide) PositionMode() PositionMode {
	switch s {
	case PositionSideLong, PositionSideShort:
		return PositionModeHedge
	default:
		return PositionModeOneway
	}
}

var sides = map[PositionMode][]PositionSide{
	PositionModeOneway: {PositionSideBoth},
	PositionModeHedge:  {PositionSideLong, PositionSideShort},
}

func (m PositionMode) GetSides() []PositionSide {
	return sides[m]
}

func NewPosition(account *Account, exchange Exchange, symbol Symbol, positionSide PositionSide) *Position {
	ts := TS()

	return &Position{
		PositionUID: uuid.New().String(),
		AccountUID:  account.AccountUID,
		Exchange:    exchange,
		Symbol:      symbol,
		Mode:        account.PositionMode,
		MarginType:  DefaultMarginType,
		Leverage:    DefaultLeverage,
		Side:        positionSide,
		CreateTS:    ts,
		UpdateTS:    ts,
		IsNew:       true,
	}
}

func (p *Position) CalcMarginBalance(price float64) {
	if p.Amount == 0 {
		return
	}

	leverage := float64(p.Leverage)

	p.MarkPrice = price

	if p.Mode == PositionModeHedge {
		if p.Side == PositionSideLong {
			p.UnrealisedPnl = p.Amount*p.MarkPrice - p.Margin*leverage
		} else {
			p.UnrealisedPnl = p.Margin*leverage - p.Amount*p.MarkPrice
		}
	} else {
		if p.Amount > 0 {
			p.UnrealisedPnl = p.Amount*p.MarkPrice - p.Margin*leverage
		} else {
			p.UnrealisedPnl = p.Margin*leverage + p.Amount*p.MarkPrice
		}
	}

	p.MarginBalance = p.Margin + p.UnrealisedPnl
}

func (p *Position) CalcLiquidationPrice(balance float64) {
	if p.Amount == 0 {
		return
	}

	leverage := float64(p.Leverage)

	if p.MarginType == MarginTypeIsolated {
		if p.Mode == PositionModeHedge {
			if p.Side == PositionSideLong {
				p.LiquidationPrice = p.Price * (1 - 1/leverage)
			} else {
				p.LiquidationPrice = p.Price * (1 + 1/leverage)
			}
		} else {
			if p.Amount > 0 {
				p.LiquidationPrice = p.Price * (1 - 1/leverage)
			} else {
				p.LiquidationPrice = p.Price * (1 + 1/leverage)
			}
		}
	} else {
		if p.Mode == PositionModeHedge {
			if p.Side == PositionSideLong {
				p.LiquidationPrice = p.Price - (balance + (balance-p.Margin)*(1-p.Margin/balance)*leverage)
			} else {
				p.LiquidationPrice = p.Price + (balance + (balance-p.Margin)*(1-p.Margin/balance)*leverage)
			}
		} else {
			if p.Amount > 0 {
				p.LiquidationPrice = p.Price - (balance + (balance-p.Margin)*(1-p.Margin/balance)*leverage)
			} else {
				p.LiquidationPrice = p.Price + (balance + (balance-p.Margin)*(1-p.Margin/balance)*leverage)
			}
		}
	}
}
