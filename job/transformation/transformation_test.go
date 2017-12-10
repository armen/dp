package transformation_test

import (
	"irdp/job"
	"irdp/job/sync"
	"irdp/job/transformation"
	"testing"
)

func TestGuaranteedResponse(t *testing.T) {
	jh := sync.New()
	jt := transformation.New(jh, 1)

	job.GuaranteedResponseTest(jt, t)
}

func TestProcess(t *testing.T) {
	jh := sync.New()
	jt := transformation.New(jh, 1)

	job.ProcessTest(jt, t)
}

func TestFailedSecondResponse(t *testing.T) {
	jh := sync.New()
	th := transformation.New(jh, 1)

	job.FailedSecondResponse(jh, th, t)
}
