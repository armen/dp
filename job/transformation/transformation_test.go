package transformation_test

import (
	"testing"

	"github.com/armen/dp/job/internal/test"
	"github.com/armen/dp/job/sync"
	"github.com/armen/dp/job/transformation"
)

func TestGuaranteedResponse(t *testing.T) {
	jh := sync.New()
	jt := transformation.New(jh, 1)

	test.GuaranteedResponseTest(jt, t)
}

func TestProcess(t *testing.T) {
	jh := sync.New()
	jt := transformation.New(jh, 1)

	test.ProcessTest(jt, t)
}

func TestFailedThirdResponseTest(t *testing.T) {
	jh := sync.New()
	th := transformation.New(jh, 1)

	test.FailedThirdResponseTest(jh, th, t)
}
