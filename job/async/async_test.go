package async_test

import (
	"irdp/job"
	"irdp/job/async"
	"testing"
)

func TestGuaranteedResponse(t *testing.T) {
	jh := async.New()
	job.GuaranteedResponseTest(jh, t)
}

func TestProcess(t *testing.T) {
	jh := async.New()
	job.ProcessTest(jh, t)
}
