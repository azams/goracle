package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	store "Goracle/RandomPkg"
)

var (
	client          *ethclient.Client
	contractAddress common.Address
	contractABI     string
	err             error
)

const (
	conAddress = "yourContractAddress"
	privKey    = "YourPrivateKey"
	infuraURL  = "ws://127.0.0.1:7545"
)

func main() {

	contractAddress = common.HexToAddress(conAddress)
	contractABI = readFile("./random.abi")

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

			if eventName.Name == "NewRandomRequest" {
				fmt.Println("[Info] New random number request detected!")
				fmt.Println("  [RESULT] Transaction hash:", vLog.TxHash.Hex())

				requestId := big.NewInt(0).SetBytes(vLog.Data).Uint64()
				seed := time.Now().UnixNano() // Initialize global pseudo random generator
				rand.New(rand.NewSource(seed))
				//randomNumber := rand.Uint64()
				randomNumber := rand.Intn(100) + 1
				fmt.Println("  [RESULT] Random number: ", randomNumber)

				err = setRandomNumber(requestId, randomNumber)
				if err != nil {
					log.Panic(err)
				}
				fmt.Println("  [RESULT] Random number set for request ID:", requestId)
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

func setRandomNumber(requestId uint64, randomNumber int) error {
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

	instance, err := store.NewRandomPkg(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	tx, err := instance.SetRandomNumber(auth, big.NewInt(int64(requestId)), big.NewInt(int64(randomNumber)))
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
