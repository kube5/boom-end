package contract

import (
	"context"

	"github.com/labstack/gommon/random"
	"github.com/sirupsen/logrus"

	"github.com/Mu-Exchange/Mu-End/common/config"
	"github.com/Mu-Exchange/Mu-End/common/proto"
	"github.com/Mu-Exchange/Mu-End/common/redis"
)

type Client struct {
	pubSub   *redis.PubSub
	monitor  proto.MonitorService
	logger   *logrus.Logger
	channels map[string]struct{}
}

type Subscription struct {
	Ch      <-chan string
	channel string
	chainId uint64
}

func NewClient(redisCfg *config.Redis, monitor proto.MonitorService, logger *logrus.Logger) (*Client, error) {
	pubSub, err := redis.NewPubSub(redisCfg)
	if err != nil {
		return nil, err
	}

	return &Client{
		pubSub:   pubSub,
		monitor:  monitor,
		logger:   logger,
		channels: make(map[string]struct{}),
	}, nil
}

func (client *Client) SubscribeLogs(ctx context.Context, chainId uint64, addresses []string, start uint64) (*Subscription, error) {
	channel := random.String(32)
	ch, err := client.pubSub.Subscribe(channel)
	if err != nil {
		client.logger.Errorf("subscribe channel failed, err: %v", err)
		return nil, err
	}

	_, err = client.monitor.SubscribeContractsLogs(ctx, &proto.SubContractsLogsReq{
		ChainId:    chainId,
		Contracts:  addresses,
		StartBlock: start,
		Channel:    channel,
	})
	if err != nil {
		return nil, err
	}

	client.channels[channel] = struct{}{}

	client.logger.Infof("subscribe logs, chainId: %d, addresses: %v, start: %d, channel: %s", chainId, addresses, start, channel)

	return &Subscription{Ch: ch, channel: channel, chainId: chainId}, nil
}

func (client *Client) UnsubscribeLogs(ctx context.Context, sub *Subscription) error {
	if _, ok := client.channels[sub.channel]; !ok {
		return nil
	}

	delete(client.channels, sub.channel)

	if err := client.pubSub.Unsubscribe(sub.channel); err != nil {
		client.logger.Errorf("unsubscribe channel from pubsub failed, err: %v", err)
	}

	_, err := client.monitor.UnsubscribeLogs(ctx, &proto.UnsubscribeLogsReq{
		ChainId: sub.chainId,
		Channel: sub.channel,
	})
	if err != nil {
		return err
	}

	return nil
}

func (client *Client) Close() error {
	for channel := range client.channels {
		if err := client.pubSub.Unsubscribe(channel); err != nil {
			client.logger.Errorf("unsubscribe channel from pubsub failed, err: %v", err)
			return err
		}
	}

	if err := client.pubSub.Close(); err != nil {
		client.logger.Errorf("close pubsub failed, err: %v", err)
		return err
	}

	return nil
}
