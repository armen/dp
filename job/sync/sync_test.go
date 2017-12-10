package sync_test

import (
	"testing"

	"github.com/armen/irdp/job"
	"github.com/armen/irdp/job/sync"
)

func TestGuaranteedResponse(t *testing.T) {
	jh := sync.New()
	job.GuaranteedResponseTest(jh, t)
}

func TestProcess(t *testing.T) {
	jh := sync.New()
	job.ProcessTest(jh, t)
}
