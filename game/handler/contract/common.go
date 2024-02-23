package contract

import (
	"context"
	"regexp"
	"strings"
	"time"

	"github.com/Rican7/retry"
	"github.com/Rican7/retry/strategy"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/sirupsen/logrus"
)

func dialRpcWithRetry(rpcUrls []string, logger logrus.FieldLogger) (*rpc.Client, int, error) {
	var rpcClient *rpc.Client
	var err error
	index := 0

	if err := retry.Retry(func(attempt uint) error {
		rpcClient, err = rpc.DialContext(context.Background(), rpcUrls[index])
		if err != nil {
			logger.Errorf("dial rpc url %s err: %w", rpcUrls[index], err)

			if isNetworkError(err) {
				index = int((attempt + 1) % uint(len(rpcUrls)))
			}
		}
		return err
	}, strategy.Wait(10*time.Second)); err != nil {
		return nil, 0, err
	}

	return rpcClient, index, nil
}

func isNetworkError(err error) bool {
	if err == nil {
		return false
	}

	return regexp.MustCompile("Post .* EOF").MatchString(err.Error()) ||
		strings.Contains(err.Error(), "connection reset by peer") ||
		strings.Contains(err.Error(), "TLS handshake timeout") ||
		strings.Contains(err.Error(), "Too Many Requests") ||
		strings.Contains(err.Error(), "502")
}
