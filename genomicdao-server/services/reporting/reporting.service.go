package reporting

import (
	"fmt"
	"hash/crc32"
)

type (
	IService interface {
		Report(geneDataComputed []byte) int
	}

	serviceImpl struct {
		scoreLevel int
	}
)

func NewReportingService() IService {
	return &serviceImpl{scoreLevel: 4} // hardcode scoreLevel = 4
}

// ---------------------------------------------------------------------------------------------------------------------

func (s *serviceImpl) Report(geneDataComputed []byte) int {
	fmt.Println("Reporting ............")
	riskScore := s.calculateRiskScore(geneDataComputed)
	fmt.Println("Reporting finished!!!!")
	return riskScore
}

func (s *serviceImpl) calculateRiskScore(geneDataComputed []byte) int {
	// hash string to int
	hashValue := crc32.ChecksumIEEE(geneDataComputed)

	// mod 4 to receive data in range [1..4]
	return int(hashValue)%s.scoreLevel + 1
}

// ---------------------------------------------------------------------------------------------------------------------
