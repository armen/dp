package test

import (
	"testing"
	"time"

	"github.com/armen/dp/job"
)

// GuaranteedResponseTest tests GuaranteedResponse property.
func GuaranteedResponseTest(jh job.Handler, t *testing.T) {
	var c = make(chan *job.Job)
	jh.Confirm(func(j *job.Job) {
		c <- j
	})
	go jh.React()

	jh.Submit(&job.Job{ID: 1, Payload: []byte("payload")})

	select {
	case j := <-c:
		if j.ID != 1 || string(j.Payload) != "payload" {
			t.Error("The confirmed job is not the submitted job")
		}
	case <-time.After(1 * time.Millisecond):
		t.Error("Confirm indication is not called")
	}
}

// ProcessTest tests if the job is processed.
func ProcessTest(jh job.Handler, t *testing.T) {
	var p = make(chan *job.Job)
	jh.Process(func(j *job.Job) {
		p <- j
	})

	jh.Confirm(func(*job.Job) {})
	go jh.React()

	jh.Submit(&job.Job{ID: 1, Payload: []byte("payload")})

	select {
	case j := <-p:
		if j.ID != 1 || string(j.Payload) != "payload" {
			t.Error("The processed job is not the submitted job")
		}
	case <-time.After(1 * time.Millisecond):
		t.Error("The job is not processed")
	}
}

// FailedThirdResponseTest tests job-transformation.
func FailedThirdResponseTest(jh job.Handler, th job.TransformationHandler, t *testing.T) {
	var c = make(chan *job.Job)
	var processing = make(chan struct{})
	var wait = make(chan struct{})

	jh.Process(func(*job.Job) {
		// Signal that we started processing
		processing <- struct{}{}
		// Don't rush on processing
		<-wait
	})

	th.Confirm(func(*job.Job) {}) // Confirm is already tested
	th.Error(func(j *job.Job) {
		c <- j
	})
	go th.React()

	th.Submit(&job.Job{ID: 1, Payload: []byte("payload 1")})
	// Wait until the first submission started processing then submit two
	// more jobs, the latter should not be processed
	<-processing
	th.Submit(&job.Job{ID: 2, Payload: []byte("payload 2")})
	th.Submit(&job.Job{ID: 3, Payload: []byte("payload 3")})

	select {
	case j := <-c:
		if j.ID != 3 || string(j.Payload) != "payload 3" {
			t.Error("The failed job is not the submitted job")
		}
	case <-time.After(1 * time.Millisecond):
		t.Error("Error indication is not called")
	}

	// We hold the first submission, now we can let it go
	close(wait)
}
