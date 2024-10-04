package rocketmq

import (
	"math"
	"strings"
	"testing"
)

func Test_MaxSize(t *testing.T) {
	msg := strings.Repeat("A", int(math.Pow(2, 24))+1)
	mqMsg := CreateMqRemotingMessage(msg, 112, 1)
	if mqMsg != nil {
		t.Error("msg size should have exceeded artificial cap")
	}

	msg = strings.Repeat("A", int(math.Pow(2, 24)))
	mqMsg = CreateMqRemotingMessage(msg, 112, 1)
	if mqMsg == nil {
		t.Error("msg size should not have exceeded the artificial cap")
	}
}
