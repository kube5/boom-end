package redis

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/redis/go-redis/v9"

	"github.com/Mu-Exchange/Mu-End/common/config"
)

// PubSub is a simple Redis Pub/Sub library
type PubSub struct {
	client   *redis.Client
	ctx      context.Context
	mu       sync.Mutex
	unsubChs map[string]chan string
}

// NewPubSub creates a new instance of the PubSub library
func NewPubSub(cfg *config.Redis) (*PubSub, error) {
	client, err := NewRedisClient(cfg)
	if err != nil {
		return nil, err
	}

	return &PubSub{
		client:   client,
		ctx:      context.Background(),
		unsubChs: make(map[string]chan string),
	}, nil
}

// Publish publishes a message to a specific channel
func (ps *PubSub) Publish(channel string, message string) (uint64, error) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	subs, err := ps.client.Publish(ps.ctx, channel, message).Result()
	return uint64(subs), err
}

// Subscribe subscribes to a specific channel and returns a channel for receiving messages
func (ps *PubSub) Subscribe(channel string) (<-chan string, error) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	pubsub := ps.client.Subscribe(ps.ctx, channel)
	_, err := pubsub.Receive(ps.ctx)
	if err != nil {
		return nil, fmt.Errorf("subscribe error: %v", err)
	}

	ch := make(chan string)
	ps.unsubChs[channel] = make(chan string)

	go func() {
		for {
			select {
			case <-ps.unsubChs[channel]:
				close(ch)
				delete(ps.unsubChs, channel)
				return
			default:
				msg, err := pubsub.ReceiveMessage(ps.ctx)
				if err != nil {
					log.Printf("receive message error: %v", err)
					close(ch)
					return
				}
				ch <- msg.Payload
			}
		}
	}()

	return ch, nil
}

// Unsubscribe unsubscribes from a specific channel
func (ps *PubSub) Unsubscribe(channel string) error {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	unsubCh, ok := ps.unsubChs[channel]
	if !ok {
		return nil
	}

	close(unsubCh)

	return nil
}

// Close closes the Pub/Sub connection
func (ps *PubSub) Close() error {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	err := ps.client.Close()
	if err != nil {
		return fmt.Errorf("close error: %v", err)
	}

	return nil
}
