package coinmarketcap

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchETHPrice(t *testing.T) {
	price, err := FetchETHPrice("e21aca65-14f0-4dcd-b1bf-7e6ba30d8e5c")
	assert.Nil(t, err)

	fmt.Println(price)
}
