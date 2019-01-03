package okex

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/thrasher-/gocryptotrader/common"
	"github.com/thrasher-/gocryptotrader/config"
	"github.com/thrasher-/gocryptotrader/currency/pair"
	"github.com/thrasher-/gocryptotrader/currency/symbol"
	exchange "github.com/thrasher-/gocryptotrader/exchanges"
)

var o OKEX

// Set API data in "../../testdata/apikeys.json"
// Copy template from "../../testdata/apikeys.example.json"
const (
	canManipulateRealOrders = false
)

func TestSetDefaults(t *testing.T) {
	o.SetDefaults()
	if o.GetName() != "OKEX" {
		t.Error("Test Failed - Bittrex - SetDefaults() error")
	}
}

func TestSetup(t *testing.T) {
	cfg := config.GetConfig()
	cfg.LoadConfig("../../testdata/configtest.json")
	exchangeConfig, err := cfg.GetExchangeConfig("OKEX")
	if err != nil {
		t.Error("Test Failed - Okex Setup() init error")
	}

	exchangeConfig.AuthenticatedAPISupport = true
	apiKeyFile, err := common.ReadFile("../../testdata/apikeys.json")
	if err == nil {
		var exchangesAPIKeys []config.ExchangeConfig
		err = json.Unmarshal(apiKeyFile, &exchangesAPIKeys)
		if err != nil {
			t.Error(err)
		}
		for _, exchangeAPIKeys := range exchangesAPIKeys {
			if strings.EqualFold(exchangeAPIKeys.Name, exchangeConfig.Name) {
				exchangeConfig.APIKey = exchangeAPIKeys.APIKey
				exchangeConfig.APISecret = exchangeAPIKeys.APISecret
				exchangeConfig.Verbose = exchangeAPIKeys.Verbose
				break
			}
		}
	}

	o.Setup(exchangeConfig)
}

func TestGetSpotInstruments(t *testing.T) {
	t.Parallel()
	_, err := o.GetSpotInstruments()
	if err != nil {
		t.Errorf("Test failed - okex GetSpotInstruments() failed: %s", err)
	}
}

func TestGetContractPrice(t *testing.T) {
	t.Parallel()
	_, err := o.GetContractPrice("btc_usd", "this_week")
	if err != nil {
		t.Error("Test failed - okex GetContractPrice() error", err)
	}
	_, err = o.GetContractPrice("btc_bla", "123525")
	if err == nil {
		t.Error("Test failed - okex GetContractPrice() error", err)
	}
	_, err = o.GetContractPrice("btc_bla", "this_week")
	if err == nil {
		t.Error("Test failed - okex GetContractPrice() error", err)
	}
}

func TestGetContractMarketDepth(t *testing.T) {
	t.Parallel()
	_, err := o.GetContractMarketDepth("btc_usd", "this_week")
	if err != nil {
		t.Error("Test failed - okex GetContractMarketDepth() error", err)
	}
	_, err = o.GetContractMarketDepth("btc_bla", "123525")
	if err == nil {
		t.Error("Test failed - okex GetContractMarketDepth() error", err)
	}
	_, err = o.GetContractMarketDepth("btc_bla", "this_week")
	if err == nil {
		t.Error("Test failed - okex GetContractMarketDepth() error", err)
	}
}

func TestGetContractTradeHistory(t *testing.T) {
	t.Parallel()
	_, err := o.GetContractTradeHistory("btc_usd", "this_week")
	if err != nil {
		t.Error("Test failed - okex GetContractTradeHistory() error", err)
	}
	_, err = o.GetContractTradeHistory("btc_bla", "123525")
	if err == nil {
		t.Error("Test failed - okex GetContractTradeHistory() error", err)
	}
	_, err = o.GetContractTradeHistory("btc_bla", "this_week")
	if err == nil {
		t.Error("Test failed - okex GetContractTradeHistory() error", err)
	}
}

func TestGetContractIndexPrice(t *testing.T) {
	t.Parallel()
	_, err := o.GetContractIndexPrice("btc_usd")
	if err != nil {
		t.Error("Test failed - okex GetContractIndexPrice() error", err)
	}
	_, err = o.GetContractIndexPrice("lol123")
	if err == nil {
		t.Error("Test failed - okex GetContractTradeHistory() error", err)
	}
}

func TestGetContractExchangeRate(t *testing.T) {
	t.Parallel()
	_, err := o.GetContractExchangeRate()
	if err != nil {
		t.Error("Test failed - okex GetContractExchangeRate() error", err)
	}
}

func TestGetContractCandlestickData(t *testing.T) {
	t.Parallel()
	_, err := o.GetContractCandlestickData("btc_usd", "1min", "this_week", 1, 2)
	if err != nil {
		t.Error("Test failed - okex GetContractCandlestickData() error", err)
	}
	_, err = o.GetContractCandlestickData("btc_bla", "1min", "this_week", 1, 2)
	if err == nil {
		t.Error("Test failed - okex GetContractCandlestickData() error", err)
	}
	_, err = o.GetContractCandlestickData("btc_usd", "min", "this_week", 1, 2)
	if err == nil {
		t.Error("Test failed - okex GetContractCandlestickData() error", err)
	}
	_, err = o.GetContractCandlestickData("btc_usd", "1min", "this_wok", 1, 2)
	if err == nil {
		t.Error("Test failed - okex GetContractCandlestickData() error", err)
	}
}

func TestGetContractHoldingsNumber(t *testing.T) {
	t.Parallel()
	_, _, err := o.GetContractHoldingsNumber("btc_usd", "this_week")
	if err != nil {
		t.Error("Test failed - okex GetContractHoldingsNumber() error", err)
	}
	_, _, err = o.GetContractHoldingsNumber("btc_bla", "this_week")
	if err == nil {
		t.Error("Test failed - okex GetContractHoldingsNumber() error", err)
	}
	_, _, err = o.GetContractHoldingsNumber("btc_usd", "this_bla")
	if err == nil {
		t.Error("Test failed - okex GetContractHoldingsNumber() error", err)
	}
}

func TestGetContractlimit(t *testing.T) {
	t.Parallel()
	_, err := o.GetContractlimit("btc_usd", "this_week")
	if err != nil {
		t.Error("Test failed - okex GetContractlimit() error", err)
	}
	_, err = o.GetContractlimit("btc_bla", "this_week")
	if err == nil {
		t.Error("Test failed - okex GetContractlimit() error", err)
	}
	_, err = o.GetContractlimit("btc_usd", "this_bla")
	if err == nil {
		t.Error("Test failed - okex GetContractlimit() error", err)
	}
}

func TestGetContractUserInfo(t *testing.T) {
	t.Parallel()
	err := o.GetContractUserInfo()
	if err == nil {
		t.Error("Test failed - okex GetContractUserInfo() error", err)
	}
}

func TestGetContractPosition(t *testing.T) {
	t.Parallel()
	err := o.GetContractPosition("btc_usd", "this_week")
	if err == nil {
		t.Error("Test failed - okex GetContractPosition() error", err)
	}
}

func TestPlaceContractOrders(t *testing.T) {
	t.Parallel()
	_, err := o.PlaceContractOrders("btc_usd", "this_week", "1", 10, 1, 1, true)
	if err == nil {
		t.Error("Test failed - okex PlaceContractOrders() error", err)
	}
}

func TestGetContractFuturesTradeHistory(t *testing.T) {
	t.Parallel()
	err := o.GetContractFuturesTradeHistory("btc_usd", "1972-01-01", 0)
	if err == nil {
		t.Error("Test failed - okex GetContractTradeHistory() error", err)
	}
}

func TestGetLatestSpotPrice(t *testing.T) {
	t.Parallel()
	_, err := o.GetLatestSpotPrice("ltc_btc")
	if err != nil {
		t.Error("Test failed - okex GetLatestSpotPrice() error", err)
	}
}

func TestGetSpotTicker(t *testing.T) {
	t.Parallel()
	_, err := o.GetSpotTicker("ltc_btc")
	if err != nil {
		t.Error("Test failed - okex GetSpotTicker() error", err)
	}
}

func TestGetSpotMarketDepth(t *testing.T) {
	t.Parallel()
	_, err := o.GetSpotMarketDepth(ActualSpotDepthRequestParams{
		Symbol: "eth_btc",
		Size:   2,
	})
	if err != nil {
		t.Error("Test failed - okex GetSpotMarketDepth() error", err)
	}
}

func TestGetSpotRecentTrades(t *testing.T) {
	t.Parallel()
	_, err := o.GetSpotRecentTrades(ActualSpotTradeHistoryRequestParams{
		Symbol: "ltc_btc",
		Since:  0,
	})
	if err != nil {
		t.Error("Test failed - okex GetSpotRecentTrades() error", err)
	}
}

func TestGetSpotKline(t *testing.T) {
	t.Parallel()
	arg := KlinesRequestParams{
		Symbol: "ltc_btc",
		Type:   TimeIntervalFiveMinutes,
		Size:   100,
	}
	_, err := o.GetSpotKline(arg)
	if err != nil {
		t.Error("Test failed - okex GetSpotCandleStick() error", err)
	}
}

func TestSpotNewOrder(t *testing.T) {
	t.Parallel()

	if o.APIKey == "" || o.APISecret == "" {
		t.Skip()
	}

	_, err := o.SpotNewOrder(SpotNewOrderRequestParams{
		Symbol: "ltc_btc",
		Amount: 1.1,
		Price:  10.1,
		Type:   SpotNewOrderRequestTypeBuy,
	})
	if err != nil {
		t.Error("Test failed - okex SpotNewOrder() error", err)
	}
}

func TestSpotCancelOrder(t *testing.T) {
	t.Parallel()

	if o.APIKey == "" || o.APISecret == "" {
		t.Skip()
	}

	_, err := o.SpotCancelOrder("ltc_btc", 519158961)
	if err != nil {
		t.Error("Test failed - okex SpotCancelOrder() error", err)
	}
}

func TestGetUserInfo(t *testing.T) {
	t.Parallel()

	if o.APIKey == "" || o.APISecret == "" {
		t.Skip()
	}

	_, err := o.GetUserInfo()
	if err != nil {
		t.Error("Test failed - okex GetUserInfo() error", err)
	}
}

func setFeeBuilder() exchange.FeeBuilder {
	return exchange.FeeBuilder{
		Amount:              1,
		Delimiter:           "-",
		FeeType:             exchange.CryptocurrencyTradeFee,
		FirstCurrency:       symbol.LTC,
		SecondCurrency:      symbol.BTC,
		IsMaker:             false,
		PurchasePrice:       1,
		CurrencyItem:        symbol.USD,
		BankTransactionType: exchange.WireTransfer,
	}
}

func TestGetFee(t *testing.T) {
	o.SetDefaults()
	var feeBuilder = setFeeBuilder()
	// CryptocurrencyTradeFee Basic
	if resp, err := o.GetFee(feeBuilder); resp != float64(0.0015) || err != nil {
		t.Error(err)
		t.Errorf("Test Failed - GetFee() error. Expected: %f, Received: %f", float64(0.0015), resp)
	}

	// CryptocurrencyTradeFee High quantity
	feeBuilder = setFeeBuilder()
	feeBuilder.Amount = 1000
	feeBuilder.PurchasePrice = 1000
	if resp, err := o.GetFee(feeBuilder); resp != float64(1500) || err != nil {
		t.Errorf("Test Failed - GetFee() error. Expected: %f, Received: %f", float64(1500), resp)
		t.Error(err)
	}

	// CryptocurrencyTradeFee IsMaker
	feeBuilder = setFeeBuilder()
	feeBuilder.IsMaker = true
	if resp, err := o.GetFee(feeBuilder); resp != float64(0.001) || err != nil {
		t.Errorf("Test Failed - GetFee() error. Expected: %f, Received: %f", float64(0.001), resp)
		t.Error(err)
	}

	// CryptocurrencyTradeFee Negative purchase price
	feeBuilder = setFeeBuilder()
	feeBuilder.PurchasePrice = -1000
	if resp, err := o.GetFee(feeBuilder); resp != float64(0) || err != nil {
		t.Errorf("Test Failed - GetFee() error. Expected: %f, Received: %f", float64(0), resp)
		t.Error(err)
	}
	// CryptocurrencyWithdrawalFee Basic
	feeBuilder = setFeeBuilder()
	feeBuilder.FeeType = exchange.CryptocurrencyWithdrawalFee
	if resp, err := o.GetFee(feeBuilder); resp != float64(0.001) || err != nil {
		t.Errorf("Test Failed - GetFee() error. Expected: %f, Received: %f", float64(0.001), resp)
		t.Error(err)
	}

	// CryptocurrencyWithdrawalFee Invalid currency
	feeBuilder = setFeeBuilder()
	feeBuilder.FirstCurrency = "hello"
	feeBuilder.FeeType = exchange.CryptocurrencyWithdrawalFee
	if resp, err := o.GetFee(feeBuilder); resp != float64(0) || err != nil {
		t.Errorf("Test Failed - GetFee() error. Expected: %f, Received: %f", float64(0), resp)
		t.Error(err)
	}

	// CyptocurrencyDepositFee Basic
	feeBuilder = setFeeBuilder()
	feeBuilder.FeeType = exchange.CyptocurrencyDepositFee
	if resp, err := o.GetFee(feeBuilder); resp != float64(0) || err != nil {
		t.Errorf("Test Failed - GetFee() error. Expected: %f, Received: %f", float64(0), resp)
		t.Error(err)
	}

	// InternationalBankDepositFee Basic
	feeBuilder = setFeeBuilder()
	feeBuilder.FeeType = exchange.InternationalBankDepositFee
	if resp, err := o.GetFee(feeBuilder); resp != float64(0) || err != nil {
		t.Errorf("Test Failed - GetFee() error. Expected: %f, Received: %f", float64(0), resp)
		t.Error(err)
	}

	// InternationalBankWithdrawalFee Basic
	feeBuilder = setFeeBuilder()
	feeBuilder.FeeType = exchange.InternationalBankWithdrawalFee
	feeBuilder.CurrencyItem = symbol.USD
	if resp, err := o.GetFee(feeBuilder); resp != float64(0) || err != nil {
		t.Errorf("Test Failed - GetFee() error. Expected: %f, Received: %f", float64(0), resp)
		t.Error(err)
	}
}

func TestFormatWithdrawPermissions(t *testing.T) {
	// Arrange
	o.SetDefaults()
	expectedResult := exchange.AutoWithdrawCryptoText + " & " + exchange.NoFiatWithdrawalsText
	// Act
	withdrawPermissions := o.FormatWithdrawPermissions()
	// Assert
	if withdrawPermissions != expectedResult {
		t.Errorf("Expected: %s, Received: %s", expectedResult, withdrawPermissions)
	}
}

// Any tests below this line have the ability to impact your orders on the exchange. Enable canManipulateRealOrders to run them
// ----------------------------------------------------------------------------------------------------------------------------
func areTestAPIKeysSet() bool {
	if o.APIKey != "" && o.APIKey != "Key" &&
		o.APISecret != "" && o.APISecret != "Secret" {
		return true
	}
	return false
}

func TestSubmitOrder(t *testing.T) {
	o.SetDefaults()
	TestSetup(t)

	if areTestAPIKeysSet() && !canManipulateRealOrders {
		t.Skip("API keys set, canManipulateRealOrders false, skipping test")
	}

	var p = pair.CurrencyPair{
		Delimiter:      "",
		FirstCurrency:  symbol.BTC,
		SecondCurrency: symbol.EUR,
	}
	response, err := o.SubmitOrder(p, exchange.Buy, exchange.Market, 1, 10, "hi")
	if areTestAPIKeysSet() && (err != nil || !response.IsOrderPlaced) {
		t.Errorf("Order failed to be placed: %v", err)
	} else if !areTestAPIKeysSet() && err == nil {
		t.Error("Expecting an error when no keys are set")
	}
}

func TestCancelExchangeOrder(t *testing.T) {
	// Arrange
	o.SetDefaults()
	TestSetup(t)

	if areTestAPIKeysSet() && !canManipulateRealOrders {
		t.Skip("API keys set, canManipulateRealOrders false, skipping test")
	}

	currencyPair := pair.NewCurrencyPair(symbol.LTC, symbol.BTC)

	var orderCancellation = exchange.OrderCancellation{
		OrderID:       "1",
		WalletAddress: "1F5zVDgNjorJ51oGebSvNCrSAHpwGkUdDB",
		AccountID:     "1",
		CurrencyPair:  currencyPair,
	}

	// Act
	err := o.CancelOrder(orderCancellation)

	// Assert
	if !areTestAPIKeysSet() && err == nil {
		t.Errorf("Expecting an error when no keys are set: %v", err)
	}
	if areTestAPIKeysSet() && err != nil {
		t.Errorf("Could not cancel orders: %v", err)
	}
}

func TestCancelAllExchangeOrders(t *testing.T) {
	// Arrange
	o.SetDefaults()
	TestSetup(t)

	if areTestAPIKeysSet() && !canManipulateRealOrders {
		t.Skip("API keys set, canManipulateRealOrders false, skipping test")
	}

	currencyPair := pair.NewCurrencyPair(symbol.LTC, symbol.BTC)

	var orderCancellation = exchange.OrderCancellation{
		OrderID:       "1",
		WalletAddress: "1F5zVDgNjorJ51oGebSvNCrSAHpwGkUdDB",
		AccountID:     "1",
		CurrencyPair:  currencyPair,
	}

	// Act
	resp, err := o.CancelAllOrders(orderCancellation)

	// Assert
	if !areTestAPIKeysSet() && err == nil {
		t.Errorf("Expecting an error when no keys are set: %v", err)
	}
	if areTestAPIKeysSet() && err != nil {
		t.Errorf("Could not cancel orders: %v", err)
	}

	if len(resp.OrderStatus) > 0 {
		t.Errorf("%v orders failed to cancel", len(resp.OrderStatus))
	}
}

func TestGetAccountInfo(t *testing.T) {
	if o.APIKey != "" || o.APISecret != "" {
		_, err := o.GetAccountInfo()
		if err != nil {
			t.Error("Test Failed - GetAccountInfo() error", err)
		}
	} else {
		_, err := o.GetAccountInfo()
		if err == nil {
			t.Error("Test Failed - GetAccountInfo() error")
		}
	}
}

func TestModifyOrder(t *testing.T) {
	_, err := o.ModifyOrder(exchange.ModifyOrder{})
	if err == nil {
		t.Error("Test failed - ModifyOrder() error")
	}
}

func TestWithdraw(t *testing.T) {
	o.SetDefaults()
	TestSetup(t)
	var withdrawCryptoRequest = exchange.WithdrawRequest{
		Amount:        100,
		Currency:      "btc_usd",
		Address:       "1F5zVDgNjorJ51oGebSvNCrSAHpwGkUdDB",
		Description:   "WITHDRAW IT ALL",
		TradePassword: "Password",
		FeeAmount:     1,
	}

	if areTestAPIKeysSet() && !canManipulateRealOrders {
		t.Skip("API keys set, canManipulateRealOrders false, skipping test")
	}

	_, err := o.WithdrawCryptocurrencyFunds(withdrawCryptoRequest)
	if !areTestAPIKeysSet() && err == nil {
		t.Errorf("Expecting an error when no keys are set: %v", err)
	}
	if areTestAPIKeysSet() && err != nil {
		t.Errorf("Withdraw failed to be placed: %v", err)
	}
}

func TestWithdrawFiat(t *testing.T) {
	o.SetDefaults()
	TestSetup(t)

	if areTestAPIKeysSet() && !canManipulateRealOrders {
		t.Skip("API keys set, canManipulateRealOrders false, skipping test")
	}

	var withdrawFiatRequest = exchange.WithdrawRequest{}

	_, err := o.WithdrawFiatFunds(withdrawFiatRequest)
	if err != common.ErrFunctionNotSupported {
		t.Errorf("Expected '%v', recieved: '%v'", common.ErrFunctionNotSupported, err)
	}
}

func TestWithdrawInternationalBank(t *testing.T) {
	o.SetDefaults()
	TestSetup(t)

	if areTestAPIKeysSet() && !canManipulateRealOrders {
		t.Skip("API keys set, canManipulateRealOrders false, skipping test")
	}

	var withdrawFiatRequest = exchange.WithdrawRequest{}

	_, err := o.WithdrawFiatFundsToInternationalBank(withdrawFiatRequest)
	if err != common.ErrFunctionNotSupported {
		t.Errorf("Expected '%v', recieved: '%v'", common.ErrFunctionNotSupported, err)
	}
}
