package structs

import (
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	"github.com/ethereum/go-ethereum/common"
)

type AllowanceRes struct {
	Allowance *lib.BigInt `json:"allowance" example:"100000000"`
}

type TxRes struct {
	Tx common.Hash `json:"tx" example:"0x1234"`
}

type ErrRes struct {
	Error error `json:"error" example:"error message"`
}

type OpenSessionRes struct {
	SessionID common.Hash `json:"sessionID" example:"0x1234"`
}

type BalanceRes struct {
	Balance *lib.BigInt `json:"balance"`
}

type ProviderRes struct {
	Provider *Provider `json:"provider"`
}

type ProvidersRes struct {
	Providers []*Provider `json:"providers"`
}

type BidRes struct {
	Bid *Bid `json:"bid"`
}

type BidsRes struct {
	Bids []*Bid `json:"bids"`
}

type ScoredBidsRes struct {
	Bids []ScoredBid `json:"bids"`
}

type ModelRes struct {
	Model *Model `json:"model"`
}

type ModelsRes struct {
	Models []*Model `json:"models"`
}

type TokenBalanceRes struct {
	MOR *lib.BigInt `json:"mor" example:"100000000"`
	ETH *lib.BigInt `json:"eth" example:"100000000"`
}

type TransactionsRes struct {
	Transactions []RawTransaction `json:"transactions"`
}

type SessionRes struct {
	Session *Session `json:"session"`
}

type SessionsRes struct {
	Sessions []*Session `json:"sessions"`
}

type BudgetRes struct {
	Budget *lib.BigInt `json:"budget" example:"100000000"`
}

type SupplyRes struct {
	Supply *lib.BigInt `json:"supply" example:"100000000"`
}

type BlockRes struct {
	Block uint64 `json:"block" example:"1234"`
}
