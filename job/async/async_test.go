package async_test

import (
	"testing"

	"github.com/armen/dp/job"
	"github.com/armen/dp/job/async"
)

func TestGuaranteedResponse(t *testing.T) {
	jh := async.New()
	job.GuaranteedResponseTest(jh, t)
}

func TestProcess(t *testing.T) {
	jh := async.New()
	job.ProcessTest(jh, t)
}
