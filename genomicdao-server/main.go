package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/quytm/blockchain-engineer-interview/genomicdao-server/contracts"
)

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:9650/ext/bc/LIFENetwork/rpc")
	//client, err := ethclient.Dial("https://api.avax-test.network/ext/bc/C/rpc")
	if err != nil {
		log.Fatal(err)
	}

	// Get the chain ID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		fmt.Println("Error getting chain ID:", err)
		return
	}

	controllerAddress := common.HexToAddress("0xb99441DA4A918604d4ee7D1CB20B362Bd3c12705")
	instance, err := contracts.NewController(controllerAddress, client)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("contract is loaded")

	userPrivateKeyHex := "56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027"
	// Convert the private key hex string to an ECDSA private key
	ecdsaPrivateKey, err := hexToECDSAPrivateKey(userPrivateKeyHex)
	if err != nil {
		fmt.Println("Error converting private key hex to ECDSA:", err)
		return
	}

	// Create a new transactor with chain ID
	auth, err := bind.NewKeyedTransactorWithChainID(ecdsaPrivateKey, chainID)
	if err != nil {
		fmt.Println("Error creating keyed transactor:", err)
		return
	}

	// Send the transaction to upload data
	tx, err := instance.UploadData(auth, "doc1")
	if err != nil {
		fmt.Println("Error UploadData:", err)
		return
	}

	// Wait for the transaction receipt to ensure it's processed
	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		fmt.Println("Error WaitMined:", err)
		return
	}

	// Log and return the transaction hash
	fmt.Println(tx.Hash())

	_ = instance
}

// hexToECDSAPrivateKey converts a hexadecimal string to an ECDSA private key.
func hexToECDSAPrivateKey(hexKey string) (*ecdsa.PrivateKey, error) {
	privateKey, err := crypto.HexToECDSA(hexKey)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}
