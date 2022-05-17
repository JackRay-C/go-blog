package encrypt

import (
	"crypto"
	"encoding/hex"
)

func Sha256(s string) string{
	hash := crypto.SHA256.New()
	hash.Write([]byte(s))
	return hex.EncodeToString(hash.Sum(nil))
}
