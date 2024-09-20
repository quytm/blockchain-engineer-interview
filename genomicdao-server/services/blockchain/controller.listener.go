package blockchain

import (
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/quytm/blockchain-engineer-interview/genomicdao-server/config"
	"github.com/quytm/blockchain-engineer-interview/genomicdao-server/contracts"
	msgqueue "github.com/quytm/blockchain-engineer-interview/genomicdao-server/pkg/messagequeue"
)

type (
	listener struct {
		controllerContract *contracts.Controller
		msgQueueUploadData *msgqueue.MessageQueue
	}
)

func NewControllerListener(cfg *config.Config, msgQueueUploadData *msgqueue.MessageQueue) (*listener, error) {
	networkCfg := cfg.GetNetworkCfg()
	clientWs, err := ethclient.Dial(networkCfg.UrlWs)
	if err != nil {
		fmt.Printf("Dial error: %v", err)
		return nil, err
	}

	controllerAddress := common.HexToAddress(networkCfg.Controller.Address) // testnet
	instance, err := contracts.NewController(controllerAddress, clientWs)
	if err != nil {
		fmt.Println("Create New Controller error:", err)
		return nil, err
	}

	return &listener{
		controllerContract: instance,
		msgQueueUploadData: msgQueueUploadData,
	}, nil
}

// ---------------------------------------------------------------------------------------------------------------------

func (l *listener) WatchUploadData() {
	watcherChannel := make(chan *contracts.ControllerUploadData)
	sub, err := l.controllerContract.WatchUploadData(&bind.WatchOpts{}, watcherChannel)
	if err != nil {
		fmt.Printf("WatchUploadData: Subscribe error: %v", err)
		log.Fatal(err)
	}

	go func() { // run async
		for {
			select {
			case err := <-sub.Err():
				fmt.Printf("WatchUploadData: listen error %v", err)
				break
			case event := <-watcherChannel:
				time.Sleep(1 * time.Second)
				fmt.Printf(" ~ WatchUploadData: UploadData with fileID = %s, sessionID = %s\n", event.DocId, event.SessionId.String())
				fmt.Printf(" ~ WatchUploadData: Publishing data to queue\n")
				msg := msgqueue.UploadDataMsg{FileID: event.DocId, SessionID: event.SessionId}
				l.msgQueueUploadData.Publish(msg)
			}
		}
	}()
}

func (l *listener) WatchMintedGNFT() {
	watcherChannel := make(chan *contracts.ControllerMintedGNFT)
	sub, err := l.controllerContract.WatchMintedGNFT(&bind.WatchOpts{}, watcherChannel)
	if err != nil {
		fmt.Printf("WatchMintedGNFT: Subscribe error: %v", err)
		log.Fatal(err)
	}

	go func() { // run async
		for {
			select {
			case err := <-sub.Err():
				fmt.Printf("WatchMintedGNFT: listen error %v", err)
				break
			case event := <-watcherChannel:
				time.Sleep(1 * time.Second)
				fmt.Printf(" ~ WatchMintedGNFT: MintedGNFT with fileID = %s, sessionID = %s, tokenID = %s\n",
					event.DocId, event.SessionId.String(), event.TokenId)
			}
		}
	}()
}

func (l *listener) WatchRewardedPCSP() {
	watcherChannel := make(chan *contracts.ControllerRewardedPCSP)
	sub, err := l.controllerContract.WatchRewardedPCSP(&bind.WatchOpts{}, watcherChannel)
	if err != nil {
		fmt.Printf("WatchRewardedPCSP: Subscribe error: %v", err)
		log.Fatal(err)
	}

	go func() { // run async
		for {
			select {
			case err := <-sub.Err():
				fmt.Printf("WatchRewardedPCSP: listen error %v", err)
				break
			case event := <-watcherChannel:
				time.Sleep(1 * time.Second)
				fmt.Printf(" ~ WatchRewardedPCSP: RewardedPCSP with fileID = %s, sessionID = %s, riskScore = %s\n",
					event.DocId, event.SessionId.String(), event.RiskScore)
			}
		}
	}()
}
