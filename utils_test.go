package gbase

import (
	"fmt"
	"testing"
	"time"
)

var _ = fmt.Print

func TestLiveProcessName(t *testing.T) {
	var ctx = SimpleContext()
	LiveProcessName(ctx, time.Second, func(max int) string {
		return "CJey " + time.Now().String()
	})
}
