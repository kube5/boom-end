package metamask

import (
	"fmt"

	"github.com/Mu-Exchange/Mu-End/common/utils/hexutil"
	"github.com/ethereum/go-ethereum/accounts"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func SignVerify(address string, originRaw []byte, sign string) (bool, error) {
	sig := hexutil.CopyBytes(hexutil.Decode(sign))
	if len(sig) != crypto.SignatureLength {
		return false, fmt.Errorf("signature must be %d bytes long", crypto.SignatureLength)
	}
	if sig[crypto.RecoveryIDOffset] != 27 && sig[crypto.RecoveryIDOffset] != 28 {
		return false, fmt.Errorf("invalid Ethereum signature (V is not 27 or 28)")
	}
	sig[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1

	rpk, err := crypto.SigToPub(accounts.TextHash(originRaw), sig)
	if err != nil {
		return false, err
	}
	return crypto.PubkeyToAddress(*rpk) == ethcommon.HexToAddress(address), nil

}
