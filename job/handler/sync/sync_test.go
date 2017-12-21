package sync_test

import (
	"testing"

	"github.com/armen/dp/job/internal/test"
	"github.com/armen/dp/job/handler/sync"
)

func TestGuaranteedResponse(t *testing.T) {
	jh := sync.New()
	test.GuaranteedResponse(jh, t)
}

func TestProcess(t *testing.T) {
	jh := sync.New()
	test.Process(jh, t)
}
