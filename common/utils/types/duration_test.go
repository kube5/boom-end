package ty

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"gotest.tools/assert"
)

type Message struct {
	Elapsed Duration `json:"elapsed"`
}

func TestDuration_MarshalJSON(t *testing.T) {
	msgEnc, err := json.Marshal(&Message{
		Elapsed: Duration{time.Second * 5},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", msgEnc)

	assert.Assert(t, string(msgEnc) == `{"elapsed":"5s"}`)

}

func TestDuration_UnmarshalJSON(t *testing.T) {
	msg := Message{}
	err := json.Unmarshal([]byte(`{"elapsed":"5s"}`), &msg)
	if err != nil {
		panic(err)
	}

	assert.Assert(t, msg.Elapsed == Duration{time.Second * 5})
}
