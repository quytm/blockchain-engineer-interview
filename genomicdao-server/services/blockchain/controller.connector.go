package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/quytm/blockchain-engineer-interview/genomicdao-server/config"
	"github.com/quytm/blockchain-engineer-interview/genomicdao-server/contracts"
	"github.com/quytm/blockchain-engineer-interview/genomicdao-server/pkg/utils"
)

type (
	IService interface {
		UploadData(ctx context.Context, fileID string)
		Confirm(ctx context.Context, fileID, contentHash, proof string, sessionID, riskScore *big.Int)
	}

	serviceImpl struct {
		cfg                *config.Config
		client             *ethclient.Client
		chainID            *big.Int
		controllerContract *contracts.Controller
		privateKey         *ecdsa.PrivateKey
	}
)

func NewControllerBlockchainConnector(ctx context.Context, cfg *config.Config) (IService, error) {
	networkCfg := cfg.GetNetworkCfg()
	client, err := ethclient.Dial(networkCfg.Url)
	if err != nil {
		fmt.Printf("Dial error: %v", err)
		return nil, err
	}

	// Get the chain ID
	chainID, err := client.NetworkID(ctx)
	if err != nil {
		fmt.Println("Getting chain ID error:", err)
		return nil, err
	}

	controllerAddress := common.HexToAddress(networkCfg.Controller.Address) // testnet
	instance, err := contracts.NewController(controllerAddress, client)
	if err != nil {
		fmt.Println("Create New Controller error:", err)
		return nil, err
	}

	ecdsaPrivateKey, err := utils.HexToECDSAPrivateKey(cfg.UserPrivateKey)
	if err != nil {
		fmt.Println("Converting private key hex to ECDSA error:", err)
		return nil, err
	}

	return &serviceImpl{
		cfg:                cfg,
		client:             client,
		chainID:            chainID,
		controllerContract: instance,
		privateKey:         ecdsaPrivateKey,
	}, nil
}

// ---------------------------------------------------------------------------------------------------------------------

func (c *serviceImpl) UploadData(ctx context.Context, fileID string) {
	// 1. Create a new transactor with chain ID
	txnOpts, err := bind.NewKeyedTransactorWithChainID(c.privateKey, c.chainID)
	if err != nil {
		fmt.Println("Creating keyed transactor error:", err)
		return
	}

	// 2. Send the transaction to upload data
	txn, err := c.controllerContract.UploadData(txnOpts, fileID)
	if err != nil {
		fmt.Println("UploadData error:", err)
		return
	}

	// 3. Wait for the transaction is processed
	_, err = bind.WaitMined(ctx, c.client, txn)
	if err != nil {
		fmt.Println("Error WaitMined:", err)
		return
	}

	// 4. Log the transaction hash
	fmt.Printf(" - UploadData: txn hash = %s\n", txn.Hash())
}

func (c *serviceImpl) Confirm(ctx context.Context, fileID, contentHash, proof string, sessionID, riskScore *big.Int) {
	// 1. Create a new transactor with chain ID
	txnOpts, err := bind.NewKeyedTransactorWithChainID(c.privateKey, c.chainID)
	if err != nil {
		fmt.Println("Creating keyed transactor error:", err)
		return
	}

	// 2. Send the transaction to upload data
	txn, err := c.controllerContract.Confirm(txnOpts, fileID, contentHash, proof, sessionID, riskScore)
	if err != nil {
		fmt.Println("Confirm error:", err)
		return
	}

	// 3. Wait for the transaction is processed
	_, err = bind.WaitMined(ctx, c.client, txn)
	if err != nil {
		fmt.Println("Error WaitMined:", err)
		return
	}

	// 4. Log the transaction hash
	fmt.Printf(" - Confirm: txn hash = %s\n", txn.Hash())
}
