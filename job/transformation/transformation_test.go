package transformation_test

import (
	"testing"

	"github.com/armen/dp/job/internal/test"
	"github.com/armen/dp/job/handler/sync"
	"github.com/armen/dp/job/transformation"
)

func TestGuaranteedResponse(t *testing.T) {
	jh := sync.New()
	jt := transformation.New(jh, 1)

	test.GuaranteedResponse(jt, t)
}

func TestProcess(t *testing.T) {
	jh := sync.New()
	jt := transformation.New(jh, 1)

	test.Process(jt, t)
}

func TestFailedThirdResponse(t *testing.T) {
	jh := sync.New()
	th := transformation.New(jh, 1)

	test.FailedThirdResponse(jh, th, t)
}
