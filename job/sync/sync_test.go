package sync_test

import (
	"testing"

	"github.com/armen/dp/job/internal/test"
	"github.com/armen/dp/job/sync"
)

func TestGuaranteedResponse(t *testing.T) {
	jh := sync.New()
	test.GuaranteedResponseTest(jh, t)
}

func TestProcess(t *testing.T) {
	jh := sync.New()
	test.ProcessTest(jh, t)
}
