package entities

import (
	"github.com/google/uuid"
)

const (
	DefaultPositionType PositionType     = "isolated"
	DefaultLeverage     PositionLeverage = 10
)

type Position struct {
	AccountUID  AccountUID       `json:"account_uid" db:"account_uid"`
	PositionUID string           `json:"position_uid" db:"position_uid"`
	Exchange    Exchange         `json:"exchange" db:"exchange"`
	Symbol      Symbol           `json:"symbol" db:"symbol"`
	Mode        PositionMode     `json:"mode" db:"position_mode"`
	Type        PositionType     `json:"type" db:"position_type"`
	Leverage    PositionLeverage `json:"leverage" db:"leverage"`
	Side        PositionSide     `json:"side" db:"side"`
	Amount      float64          `json:"amount" db:"amount"`
	Price       float64          `json:"price" db:"price"`
	MarkPrice   float64          `json:"mark_price" db:"-"`
	CreateTS    int64            `json:"create_ts" db:"create_ts"`
	UpdateTS    int64            `json:"update_ts" db:"update_ts"`
	IsNew       bool             `json:"-" db:"-"`
	HoldAmount  float64          `json:"-" db:"hold_amount"`
}

type PositionMode string

const (
	PositionModeOneway PositionMode = "oneway"
	PositionModeHedge  PositionMode = "hedge"
)

type PositionType string

const (
	PositionTypeIsolated PositionType = "isolated"
	PositionTypeCross    PositionType = "cross"
)

type PositionLeverage int32

type PositionSide string

const (
	PositionSideLong  PositionSide = "long"
	PositionSideShort PositionSide = "short"
	PositionSideBoth  PositionSide = "both"
)

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
		Type:        DefaultPositionType,
		Leverage:    DefaultLeverage,
		Side:        positionSide,
		CreateTS:    ts,
		UpdateTS:    ts,
		IsNew:       true,
	}
}
