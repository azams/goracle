package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	store "Goracle/PricePkg"
)

var (
	client          *ethclient.Client
	contractAddress common.Address
	contractABI     string
	err             error
)

const (
	conAddress   = "0xa7E3069287463109D943d3c7DAF6cDAE0e27791a"
	privKey      = "4a12c1c1071ae2469af0eb01d801f9d277255de47b51b5356597f9b5b0cfbe71"
	marketcapAPI = "4dece0c3-8e91-47b3-8549-9c85c1264045"
	infuraURL    = "ws://127.0.0.1:7545"
)

type coinMarketCapResponse struct {
	Data map[string]coinData `json:"data"`
}

type coinData struct {
	Quote quoteData `json:"quote"`
}

type quoteData struct {
	USD priceData `json:"USD"`
}

type priceData struct {
	Price float64 `json:"price"`
}

/*
solc contracts/Price.sol --abi | tail -n1 > contract.abi
solc --abi --bin contracts/Price.sol -o build/ --overwrite
abigen --pkg=PricePkg --out=PricePkg/PricePkg.go --abi=build/Price.abi
*/

func main() {
	// Replace the below variables with your own configuration
	contractAddress = common.HexToAddress(conAddress)
	contractABI = readFile("./contract.abi")

	client, err = ethclient.Dial(infuraURL)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("[Info] Connected to the server")

	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("[Info] ABI parsed successfully")

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("[Info] Subscribed successfully")

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			eventName, err := parsedABI.EventByID(vLog.Topics[0])
			if err != nil {
				log.Fatal(err)
			}

			if eventName.Name == "NewTransaction" {
				fmt.Println("[Info] New transaction detected!")
				fmt.Println("  [RESULT] Transaction hash:", vLog.TxHash.Hex())
				//transactionHash := vLog.TxHash
				length := int(vLog.Data[63])
				fmt.Println(vLog.Data[64 : 64+length])
				coin := string(vLog.Data[64 : 64+length])

				// Retrieve the transaction receipt
				price, err := fetchCoinPrice(coin, marketcapAPI)
				if err != nil {
					log.Panic(err)
				}
				fmt.Println("  [RESULT] ", coin, "->", price)
				err = updateCoin(coin, fmt.Sprintf("%f", price))
				if err != nil {
					log.Panic(err)
				}
				fmt.Println("  [RESULT] Coin price updated")
			}
		}
	}
}

func readFile(filename string) string {
	res, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(res)
}

func fetchCoinPrice(coin string, apiKey string) (float64, error) {
	apiURL := fmt.Sprintf("https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest?symbol=%s&convert=USD", coin)

	// Create an HTTP client
	client := &http.Client{}

	// Create an HTTP GET request
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return 0, err
	}

	// Set the API key in the request header
	req.Header.Set("X-CMC_PRO_API_KEY", apiKey)

	// Send the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	// Parse the response JSON
	var data coinMarketCapResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return 0, err
	}

	// Get the price from the response
	price, ok := data.Data[coin]
	if !ok {
		return 0, fmt.Errorf("price data not found for coin: %s", coin)
	}

	return price.Quote.USD.Price, nil
}

func updateCoin(coin string, price string) error {
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = gasPrice

	instance, err := store.NewPricePkg(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	tx, err := instance.UpdateCoinPrice(auth, coin, price)
	if err != nil {
		return err
	}

	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return err
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		return fmt.Errorf("transaction failed. status: %v", receipt.Status)
	}
	return nil
}
