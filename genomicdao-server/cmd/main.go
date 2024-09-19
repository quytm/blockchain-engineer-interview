package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"os/signal"
	"syscall"

	"github.com/quytm/blockchain-engineer-interview/genomicdao-server/config"
	msgqueue "github.com/quytm/blockchain-engineer-interview/genomicdao-server/pkg/messagequeue"
	"github.com/quytm/blockchain-engineer-interview/genomicdao-server/pkg/utils"
	"github.com/quytm/blockchain-engineer-interview/genomicdao-server/services/auth"
	"github.com/quytm/blockchain-engineer-interview/genomicdao-server/services/blockchain"
	"github.com/quytm/blockchain-engineer-interview/genomicdao-server/services/datastorage"
	"github.com/quytm/blockchain-engineer-interview/genomicdao-server/services/reporting"
	"github.com/quytm/blockchain-engineer-interview/genomicdao-server/services/tee"
)

func main() {

	cfg := config.Load()
	// 0. Init system
	fmt.Printf("0. Init system ----------------------------------------------------------------------------------\n")
	userPublicKey, err := cfg.GetUserPublicKey()
	if err != nil {
		log.Fatal(err)
		return
	}
	ctx := context.Background()
	authService := auth.NewAuthService()
	datastorageService := datastorage.NewDataStorageService()
	teeService := tee.NewTeeService()
	reportingService := reporting.NewReportingService()
	blControllerConnector, err := blockchain.NewControllerBlockchainConnector(ctx, &cfg)
	if err != nil {
		log.Fatal(err)
		return
	}
	msgQueueUploadData := msgqueue.NewMessageQueue(2)
	controllerListener, err := blockchain.NewControllerListener(&cfg, msgQueueUploadData)
	if err != nil {
		log.Fatal(err)
		return
	}
	controllerListener.WatchUploadData()
	controllerListener.WatchMintedGNFT()
	controllerListener.WatchRewardedPCSP()
	go msgQueueUploadData.Listen(func(data interface{}) {
		confirmInBlockchain(blControllerConnector, datastorageService, data)
	})

	fmt.Println("- Init services")
	fmt.Println("- Connect to Blockchain network")
	fmt.Println("- Listen event from Blockchain network")
	fmt.Println("- Message Queue for handling event")
	fmt.Print("\n")

	// 1. After purchasing service, user can register an account by using their public_key
	fmt.Printf("1. Register -------------------------------------------------------------------------------------\n")
	userID := authService.Register(userPublicKey)
	fmt.Printf(" - userID = %s", userID)
	fmt.Print("\n\n")

	// 2. Laboratory submit saliva sample information to Data Storage
	fmt.Printf("2. SampleRecording ------------------------------------------------------------------------------\n")
	salivaSampleID := datastorageService.SampleRecording(userID)
	fmt.Printf(" - salivaSampleID = %s", salivaSampleID)
	fmt.Print("\n\n")

	// 3. laboratory processes the saliva sample and generates the user’s raw genetic data in a text file format
	fmt.Printf("3. laboratory processes the saliva sample and generates the user’s raw genetic data -------------\n")
	textFile := utils.RandString(1000)
	// then send to TEE
	geneticDataRawHashed, geneticDataRawEncrypted, geneticDataComputedEncrypted := teeService.ComputeEncryptAndStore([]byte(textFile))
	// data is store in Data Storage
	fileID := datastorageService.SaveGeneticDataForSample(salivaSampleID, geneticDataRawHashed, geneticDataRawEncrypted, geneticDataComputedEncrypted)
	fmt.Printf(" - out is a text file, with fileID = %s", fileID)
	fmt.Print("\n\n")

	// 4. Report
	fmt.Printf("4. Report ---------------------------------------------------------------------------------------\n")
	riskScore := reportingService.Report(geneticDataComputedEncrypted)
	fmt.Printf(" - out is riskScore = %d\n", riskScore)
	datastorageService.SaveRiskScore(fileID, riskScore)
	fmt.Printf(" - saved riskScore to Data Storage")
	fmt.Print("\n\n")

	// 5. Upload Data to blockchain
	fmt.Printf("5. Upload Data to Blockchain --------------------------------------------------------------------\n")
	fmt.Printf("\t\t(note: after uploading, we also receive event `UploadData` from blockchain)\n")
	blControllerConnector.UploadData(ctx, fileID)
	fmt.Print("\n\n")

	// 6. Confirm in blockchain
	// run asynchronous in `confirmInBlockchain` method

	fmt.Printf("n. Finished!!! ----------------------------------------------------------------------------------\n")
	waitForKillingSign()
}

func confirmInBlockchain(
	blControllerConnector blockchain.IService, datastorageService datastorage.IService,
	data interface{}) {
	// 6. Confirm in blockchain
	fmt.Print("...\n...\n...\n")
	fmt.Printf("6. Confirm in blockchain --------------------------------------------------------------------\n")
	updateDataMsg := data.(msgqueue.UploadDataMsg)
	fileID := updateDataMsg.FileID
	sessionID := updateDataMsg.SessionID
	fmt.Printf(" - event `UploadData` contains fileID, sessionID, we will use them to confirm\n")
	geneticData, existed := datastorageService.GetGeneticData(fileID)
	if !existed {
		fmt.Printf("---> FileID %s not found!", fileID)
		return
	}

	geneticDataHashed := string(geneticData.GeneticDataHashed)
	riskScore := big.NewInt(int64(geneticData.RiskScore))
	proof := "success" // hardcode to success
	fmt.Printf(" - confirming to blockchain\n")
	fmt.Printf("\t\t(note: after confirming, we also receive event `MintedGNFT`, `RewardedPCSP` from blockchain)\n")
	blControllerConnector.Confirm(context.Background(), fileID, geneticDataHashed, proof, sessionID, riskScore)
	fmt.Print("\n\n")
}

func waitForKillingSign() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	endSignal := <-sig
	fmt.Printf("Process end due to signal: %s", endSignal.String())
}
