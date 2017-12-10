package sync_test

import (
	"irdp/job"
	"irdp/job/sync"
	"testing"
)

func TestGuaranteedResponse(t *testing.T) {
	jh := sync.New()
	job.GuaranteedResponseTest(jh, t)
}

func TestProcess(t *testing.T) {
	jh := sync.New()
	job.ProcessTest(jh, t)
}
