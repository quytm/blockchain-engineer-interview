package msgqueue

import "math/big"

type UploadDataMsg struct {
	FileID    string
	SessionID *big.Int
}
