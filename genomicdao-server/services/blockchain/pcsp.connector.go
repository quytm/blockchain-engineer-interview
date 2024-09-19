package blockchain

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/quytm/blockchain-engineer-interview/genomicdao-server/config"
	"github.com/quytm/blockchain-engineer-interview/genomicdao-server/contracts"
)

type (
	IPcspConnector interface {
		GetBalance(address common.Address) (*big.Int, error)
	}

	pcspConnectorImpl struct {
		cfg          *config.Config
		client       *ethclient.Client
		chainID      *big.Int
		pcspContract *contracts.PostCovidStrokePrevention
		privateKey   *ecdsa.PrivateKey
	}
)

func NewPcspBlockchainConnector(cfg *config.Config) (IPcspConnector, error) {
	networkCfg := cfg.GetNetworkCfg()
	client, err := ethclient.Dial(networkCfg.Url)
	if err != nil {
		fmt.Printf("Dial error: %v", err)
		return nil, err
	}

	pcspAddress := common.HexToAddress(networkCfg.PostCovidStrokePrevention.Address) // testnet
	instance, err := contracts.NewPostCovidStrokePrevention(pcspAddress, client)
	if err != nil {
		fmt.Println("Create New PostCovidStrokePrevention error:", err)
		return nil, err
	}

	return &pcspConnectorImpl{
		cfg:          cfg,
		pcspContract: instance,
	}, nil
}

// ---------------------------------------------------------------------------------------------------------------------

// GetBalance get balance
func (s *pcspConnectorImpl) GetBalance(address common.Address) (*big.Int, error) {
	balance, err := s.pcspContract.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		return nil, err
	}
	return balance, nil
}
