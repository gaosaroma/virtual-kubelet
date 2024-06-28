package common

import (
	"context"
	"gotest.tools/assert"
	"testing"
	"time"
)

func TestTimedTaskWithInterval(t *testing.T) {
	testList := make([]int, 0)
	go TimedTaskWithInterval("test timed task", time.Second, func(_ context.Context) {
		testList = append(testList, 0)
	})
	time.Sleep(2 * time.Second)
	assert.Assert(t, len(testList) == 2)
}
