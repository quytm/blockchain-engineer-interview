package datastorage

import (
	"fmt"

	"github.com/quytm/blockchain-engineer-interview/genomicdao-server/pkg/utils"
)

type (
	IService interface {
		SampleRecording(string) string
		GetGeneticData(fileID string) (GeneticFileData, bool)
		SaveGeneticDataForSample(string, []byte, []byte, []byte) string
		SaveRiskScore(fileID string, riskScore int)
	}

	serviceImpl struct {
		salivaSamples map[string]SalivaSample    // In-memory DB, instead of `saliva_samples` table in DB
		geneticData   map[string]GeneticFileData // In-memory DB, instead of `genetic_data` table in DB
	}
)

func NewDataStorageService() IService {
	return &serviceImpl{
		salivaSamples: make(map[string]SalivaSample),
		geneticData:   make(map[string]GeneticFileData),
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func (s *serviceImpl) SampleRecording(userID string) string {
	// 1. Generate ID
	salivaSampleID := utils.RandString(10)

	// 2. Save to DB
	s.salivaSamples[salivaSampleID] = SalivaSample{
		ID:     salivaSampleID,
		UserID: userID,
	}

	// 3. Return
	return salivaSampleID
}

func (s *serviceImpl) GetGeneticData(fileID string) (GeneticFileData, bool) {
	data, existed := s.geneticData[fileID]
	return data, existed
}

func (s *serviceImpl) SaveGeneticDataForSample(salivaSampleID string,
	geneticDataRawHashed, geneticDataRawEncrypted, geneticDataComputedEncrypted []byte) string {
	// 1. Get sample
	if _, existed := s.salivaSamples[salivaSampleID]; !existed {
		// return error
		return ""
	}

	// 2. Save to DB
	fileID := utils.RandString(10)
	s.geneticData[fileID] = GeneticFileData{
		FileID:                       fileID,
		SampleID:                     salivaSampleID,
		GeneticDataHashed:            geneticDataRawHashed,
		GeneticDataEncrypted:         geneticDataRawEncrypted,
		GeneticDataComputedEncrypted: geneticDataComputedEncrypted,
	}

	// 3. Return
	return fileID
}

func (s *serviceImpl) SaveRiskScore(fileID string, riskScore int) {
	// 1. Find data
	data, existed := s.geneticData[fileID]
	if !existed {
		fmt.Printf("SaveRiskScore error because fileID %s not found", fileID)
		return
	}

	// 2. update riskScore
	data.RiskScore = riskScore

	// 3. save to storage
	s.geneticData[fileID] = data
}
