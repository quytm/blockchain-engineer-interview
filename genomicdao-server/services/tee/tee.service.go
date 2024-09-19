package tee

type (
	IService interface {
		ComputeEncryptAndStore(geneticDataRawBytes []byte) ([]byte, []byte, []byte)
	}

	serviceImpl struct {
	}
)

func NewTeeService() IService {
	return &serviceImpl{}
}

// ---------------------------------------------------------------------------------------------------------------------

// ComputeEncryptAndStore ... TEE computes the user’s genetic data then encrypted them, after that, it calls DataStorage service to store data
func (s *serviceImpl) ComputeEncryptAndStore(geneticDataRawBytes []byte) ([]byte, []byte, []byte) {
	// TEE computes the user’s genetic data then encrypted them
	// 1. Compute user’s genetic
	geneticDataRaw := string(geneticDataRawBytes)
	geneticDataComputed := s.computing(geneticDataRaw)

	// 2. Encrypt data
	geneticDataRawHashed := s.hashing(geneticDataRaw)
	geneticDataRawEncrypted := s.encrypting(geneticDataRaw)
	geneticDataComputedEncrypted := s.encrypting(geneticDataComputed)
	return geneticDataRawHashed, geneticDataRawEncrypted, geneticDataComputedEncrypted
}

// ---------------------------------------------------------------------------------------------------------------------

func (s *serviceImpl) computing(geneticDataRaw string) string {
	return geneticDataRaw + "-computed"
}

func (s *serviceImpl) encrypting(rawData string) []byte {
	return []byte((rawData + "-encrypted"))
}

func (s *serviceImpl) hashing(rawData string) []byte {
	return []byte((rawData + "-hashed"))
}
