// Package test implements end to end tests for job handler interface.
package test

import (
	"testing"
	"time"

	"github.com/armen/dp/job"
)

// GuaranteedResponse tests GuaranteedResponse property.
func GuaranteedResponse(jh job.Handler, t *testing.T) {
	var c = make(chan job.Job)
	jh.Confirm(func(j job.Job) {
		c <- j
	})

	jh.Submit("job1")

	select {
	case j := <-c:
		if payload, ok := j.(string); !ok || payload != "job1" {
			t.Error("The confirmed job is not the submitted job")
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("Confirm indication is not called")
	}
}

// Process tests if the job is processed.
func Process(jh job.Handler, t *testing.T) {
	var p = make(chan job.Job)
	jh.Process(func(j job.Job) {
		p <- j
	})

	jh.Confirm(func(job.Job) {})

	jh.Submit("job1")

	select {
	case j := <-p:
		if payload, ok := j.(string); !ok || payload != "job1" {
			t.Error("The processed job is not the submitted job")
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("The job is not processed")
	}
}
