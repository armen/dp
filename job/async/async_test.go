package async_test

import (
	"testing"

	"github.com/armen/dp/job/async"
	"github.com/armen/dp/job/internal/test"
)

func TestGuaranteedResponse(t *testing.T) {
	jh := async.New()
	test.GuaranteedResponse(jh, t)
}

func TestProcess(t *testing.T) {
	jh := async.New()
	test.Process(jh, t)
}
