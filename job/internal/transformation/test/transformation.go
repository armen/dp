// Package test implements end to end tests for job transformation interface.
package test

import (
	"testing"
	"time"

	"github.com/armen/dp/job"
	"github.com/armen/dp/job/internal/handler/test"
)

// GuaranteedResponse tests GuaranteedResponse property.
func GuaranteedResponse(jh job.Handler, t *testing.T) {
	test.GuaranteedResponse(jh, t)
}

// Process tests if the job is processed.
func Process(jh job.Handler, t *testing.T) {
	test.Process(jh, t)
}

// FailedThirdResponse tests job-transformation.
func FailedThirdResponse(jh job.Handler, th job.TransformationHandler, t *testing.T) {
	var c = make(chan job.Job)
	var processing = make(chan struct{})
	var wait = make(chan struct{})

	jh.Process(func(job.Job) {
		// Signal that we started processing
		processing <- struct{}{}
		// Don't rush on processing
		<-wait
	})

	th.Confirm(func(job.Job) {}) // Confirm is already tested
	th.Error(func(j job.Job) {
		c <- j
	})

	th.Submit("job1")
	// Wait until the first submission started processing then submit two
	// more jobs, the latter should not be processed
	<-processing
	th.Submit("job2")
	th.Submit("job3")

	select {
	case j := <-c:
		if payload, ok := j.(string); !ok || payload != "job3" {
			t.Error("The failed job is not the submitted job")
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("Error indication is not called")
	}

	// We hold the first submission, now we can let it go
	close(wait)
}
