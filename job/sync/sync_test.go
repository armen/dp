package sync_test

import (
	"testing"

	"github.com/armen/dp/job"
	"github.com/armen/dp/job/sync"
)

func TestGuaranteedResponse(t *testing.T) {
	jh := sync.New()
	job.GuaranteedResponseTest(jh, t)
}

func TestProcess(t *testing.T) {
	jh := sync.New()
	job.ProcessTest(jh, t)
}
