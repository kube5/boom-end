package ens

import (
	"context"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func TestEns(t *testing.T) {
	rpcClient, err := rpc.DialContext(context.Background(), "https://rpc.ankr.com/eth/131d7e387f50ea1cb7636b78974bfe6fc7140536e2f334cd0a0c4a0d1c4a618d")
	if err != nil {
		fmt.Print(err)
	}

	ethClient := ethclient.NewClient(rpcClient)
	resolve, err := ReverseResolve(ethClient, "0xC8F952F862f781f2A16AdACD1D3649F9a9351E75")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(resolve)
}
