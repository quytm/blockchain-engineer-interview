package datastorage

type SalivaSample struct {
	ID     string
	UserID string
}

type GeneticFileData struct {
	FileID                       string
	SampleID                     string
	GeneticDataHashed            []byte // Genetic Data is hashed
	GeneticDataEncrypted         []byte // Genetic Data is encrypted
	GeneticDataComputedEncrypted []byte // Genetic Data is computed and encrypted
	RiskScore                    int
}
