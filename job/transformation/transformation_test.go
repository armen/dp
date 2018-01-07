package transformation_test

import (
	"fmt"
	"testing"

	"github.com/armen/dp/job"
	jh "github.com/armen/dp/job/handler/sync"
	"github.com/armen/dp/job/internal/transformation/test"
	"github.com/armen/dp/job/transformation"
)

func TestGuaranteedResponse(t *testing.T) {
	jh := jh.New()
	jt := transformation.New(jh, 1)

	test.GuaranteedResponse(jt, t)
}

func TestProcess(t *testing.T) {
	jh := jh.New()
	jt := transformation.New(jh, 1)

	test.Process(jt, t)
}

func TestFailedThirdResponse(t *testing.T) {
	jh := jh.New()
	th := transformation.New(jh, 1)

	test.FailedThirdResponse(jh, th, t)
}

// Three jobs are submitted. First one is in progress
// and the second one is queued (i.e. buffer) hence the
// queue doesn't have room for the third job, that's
// why the third job failed.
func Example() {
	var msg = make(chan string)
	var processing = make(chan struct{})
	var wait = make(chan struct{})

	jh := jh.New()
	jh.Process(func(job.Job) {
		// Signal that we started processing
		processing <- struct{}{}
		// Don't rush on processing
		<-wait
	})

	th := transformation.New(jh, 1)
	th.Confirm(func(j job.Job) {
		msg <- fmt.Sprintln("Confirmed", j)
	})
	th.Error(func(j job.Job) {
		msg <- fmt.Sprintln("Failed to execute", j)
	})

	th.Submit("Job 1")
	<-processing
	th.Submit("Job 2")
	th.Submit("Job 3")

	fmt.Print(<-msg)
	fmt.Print(<-msg)
	fmt.Print(<-msg)

	close(wait)

	// Unordered output:
	// Confirmed Job 1
	// Confirmed Job 2
	// Failed to execute Job 3
}
