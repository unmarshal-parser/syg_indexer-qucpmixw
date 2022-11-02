// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"github.com/shopspring/decimal"
	"gorm.io/datatypes"
	"math/big"
	"time"
)

var (
	_ = decimal.Decimal{}
	_ = big.NewInt
	_ = datatypes.JSON{}
	_ = time.Time{}
)

func GetSwapTargetAddedEventHash() string {
	return "0xb907822409611d127ab6a64611591b98e03a6a85ade4f258bae26b7c1efdfeaf"
}

type SwapTargetAddedEvent struct {
	Target string

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:c4794f7d-635c-40b0-b2bd-a0d8a5c781e8,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:c4794f7d-635c-40b0-b2bd-a0d8a5c781e8,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	Index           uint            `gorm:"uniqueIndex:c4794f7d-635c-40b0-b2bd-a0d8a5c781e8,unique"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetSwapTargetRemovedEventHash() string {
	return "0x393b8be3e26787f19285ecd039dfd80bc6507828750f4d50367e6efe2524695c"
}

type SwapTargetRemovedEvent struct {
	Target string

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:694a6519-2403-4b83-ac6f-6d9e0fa6b473,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:694a6519-2403-4b83-ac6f-6d9e0fa6b473,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	Index           uint            `gorm:"uniqueIndex:694a6519-2403-4b83-ac6f-6d9e0fa6b473,unique"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetTokenWithdrawnEventHash() string {
	return "0x8210728e7c071f615b840ee026032693858fbcd5e5359e67e438c890f59e5620"
}

type TokenWithdrawnEvent struct {
	Token  string
	Target string
	Amount decimal.Decimal `gorm:"type:numeric"`

	TokenPriceAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedAmount float64 `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:5d635f6f-8e6a-4f56-855d-deb0dd93814a,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:5d635f6f-8e6a-4f56-855d-deb0dd93814a,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	Index           uint            `gorm:"uniqueIndex:5d635f6f-8e6a-4f56-855d-deb0dd93814a,unique"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetEthWithdrawnEventHash() string {
	return "0x8455ae6be5d92f1df1c3c1484388e247a36c7e60d72055ae216dbc258f257d4b"
}

type EthWithdrawnEvent struct {
	Target string
	Amount decimal.Decimal `gorm:"type:numeric"`

	TokenPriceAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedAmount float64 `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:c3c13441-9c1a-4fd0-8363-25689645ceaa,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:c3c13441-9c1a-4fd0-8363-25689645ceaa,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	Index           uint            `gorm:"uniqueIndex:c3c13441-9c1a-4fd0-8363-25689645ceaa,unique"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetTransferOwnershipMethodHash() string {
	return "f2fde38b"
}

type TransferOwnershipMethod struct {
	NewOwner string

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:5411cc43-164d-4148-a744-096430ae788b,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:5411cc43-164d-4148-a744-096430ae788b,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetUpdateSwapTargetsMethodHash() string {
	return "97bbda0e"
}

type UpdateSwapTargetsMethod struct {
	Target string
	Add    bool

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:ccce6ffb-2af1-4f3d-8b1d-b70cada11006,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:ccce6ffb-2af1-4f3d-8b1d-b70cada11006,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetWithdrawEthMethodHash() string {
	return "1b9a91a4"
}

type WithdrawEthMethod struct {
	To     string
	Amount decimal.Decimal `gorm:"type:numeric"`

	TokenPriceAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedAmount float64 `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:99ee54cc-4829-4eab-b860-95f21301fa64,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:99ee54cc-4829-4eab-b860-95f21301fa64,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetFillQuoteTokenToEthWithPermitMethodHash() string {
	return "b3093838"
}

type FillQuoteTokenToEthWithPermitMethod struct {
	SellTokenAddress         string
	Target                   string
	SwapCallData             []byte
	SellAmount               decimal.Decimal `gorm:"type:numeric"`
	FeePercentageBasisPoints decimal.Decimal `gorm:"type:numeric"`
	PermitData               datatypes.JSON

	TokenPriceSellAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedSellAmount float64 `gorm:"type:numeric"`

	TokenPriceFeePercentageBasisPoints      float64 `gorm:"type:numeric"`
	DecimalAdjustedFeePercentageBasisPoints float64 `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:03b512ac-f02a-4dd5-86c2-7cfd9538ea68,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:03b512ac-f02a-4dd5-86c2-7cfd9538ea68,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetFillQuoteTokenToTokenMethodHash() string {
	return "55e4b7be"
}

type FillQuoteTokenToTokenMethod struct {
	SellTokenAddress string
	BuyTokenAddress  string
	Target           string
	SwapCallData     []byte
	SellAmount       decimal.Decimal `gorm:"type:numeric"`
	FeeAmount        decimal.Decimal `gorm:"type:numeric"`

	TokenPriceSellAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedSellAmount float64 `gorm:"type:numeric"`

	TokenPriceFeeAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedFeeAmount float64 `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:edb45500-dd24-4da2-b938-c102e51c87fc,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:edb45500-dd24-4da2-b938-c102e51c87fc,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetFillQuoteTokenToTokenWithPermitMethodHash() string {
	return "b0480bbd"
}

type FillQuoteTokenToTokenWithPermitMethod struct {
	SellTokenAddress string
	BuyTokenAddress  string
	Target           string
	SwapCallData     []byte
	SellAmount       decimal.Decimal `gorm:"type:numeric"`
	FeeAmount        decimal.Decimal `gorm:"type:numeric"`
	PermitData       datatypes.JSON

	TokenPriceSellAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedSellAmount float64 `gorm:"type:numeric"`

	TokenPriceFeeAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedFeeAmount float64 `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:1292455b-be77-4438-ad83-94bddb4417fc,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:1292455b-be77-4438-ad83-94bddb4417fc,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetWithdrawTokenMethodHash() string {
	return "01e33667"
}

type WithdrawTokenMethod struct {
	Token  string
	To     string
	Amount decimal.Decimal `gorm:"type:numeric"`

	TokenPriceAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedAmount float64 `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:5af3dbc5-7475-40d7-8957-1b57dc46f487,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:5af3dbc5-7475-40d7-8957-1b57dc46f487,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetFillQuoteEthToTokenMethodHash() string {
	return "3c2b9a7d"
}

type FillQuoteEthToTokenMethod struct {
	BuyTokenAddress string
	Target          string
	SwapCallData    []byte
	FeeAmount       decimal.Decimal `gorm:"type:numeric"`

	TokenPriceFeeAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedFeeAmount float64 `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:749c9efa-6ac7-4a88-815b-de3741ef522e,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:749c9efa-6ac7-4a88-815b-de3741ef522e,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetFillQuoteTokenToEthMethodHash() string {
	return "999b6464"
}

type FillQuoteTokenToEthMethod struct {
	SellTokenAddress         string
	Target                   string
	SwapCallData             []byte
	SellAmount               decimal.Decimal `gorm:"type:numeric"`
	FeePercentageBasisPoints decimal.Decimal `gorm:"type:numeric"`

	TokenPriceSellAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedSellAmount float64 `gorm:"type:numeric"`

	TokenPriceFeePercentageBasisPoints      float64 `gorm:"type:numeric"`
	DecimalAdjustedFeePercentageBasisPoints float64 `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:2976a8d3-2fd8-47dc-859b-b601ffb3d0d0,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:2976a8d3-2fd8-47dc-859b-b601ffb3d0d0,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

type LastSyncedBlock struct {
	Contract    string `gorm:"primaryKey"`
	ChainID     string `gorm:"primaryKey"`
	SyncType    string `gorm:"primaryKey"`
	BlockNumber uint64
}

// Plugin Models

type TokenDetails struct {
	ID      int
	Address string `gorm:"uniqueIndex:address_and_chain"`
	Symbol  string
	ChainID string `gorm:"uniqueIndex:address_and_chain"`
	Decimal int
}

// Config

type PostgresConfig struct {
	ConnectionString string `mapstructure:"connection_string"`
	TablePrefix      string `mapstructure:"table_prefix"`
	CreateBatchSize  int    `mapstructure:"create_batch_size"`
}

type IndexerConfig struct {
	EthEndpoint       string `mapstructure:"eth_endpoint"`
	ContractAddress   string `mapstructure:"contract_address"`
	StartBlock        int    `mapstructure:"start_block"`
	ApiKey            string `mapstructure:"api_key"`
	PostgresConfig    `mapstructure:"postgres_config"`
	LagToHighestBlock int `mapstructure:"lag_to_highest_block"`
	StepSize          int `mapstructure:"step_size"`
	ParallelCalls     int `mapstructure:"parallel_calls_for_logs"`
}

func (i *IndexerConfig) AssignDefaults() {
	if i.PostgresConfig.CreateBatchSize == 0 {
		i.PostgresConfig.CreateBatchSize = 100
	}
	if i.StepSize == 0 {
		i.StepSize = 50
	}
	if i.LagToHighestBlock == 0 {
		i.LagToHighestBlock = 10
	}
}
