package metamask

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestSign(t *testing.T) {
	hash := "dbdbf94250314de85f9540708de0952c36163ba2dd3933f1ea07a3e982618b59"
	sign := "0x2e2aa705f8a929e79ca5e0db96d69d5afa3379727bd8dddab0df7bf4e646c6866aa6d21b33d61908f85c2b7a63a890d504afcf98dc07ae1ad8fe03bfdd268b9a1b"
	address := "0x817016163775AaF0B25DF274fB4b18edB67E1F26"
	t.Run("test_hash", func(t *testing.T) {
		isVerify, err := SignVerify(
			address,
			common.Hex2Bytes(hash),
			sign,
		)
		require.Nil(t, err)
		require.Equal(t, isVerify, true)
	})
}
