package sync_test

import (
	"fmt"
	"testing"

	"github.com/armen/dp/job"
	jh "github.com/armen/dp/job/handler/sync"
	"github.com/armen/dp/job/internal/handler/test"
)

func TestGuaranteedResponse(t *testing.T) {
	jh := jh.New()
	test.GuaranteedResponse(jh, t)
}

func TestProcess(t *testing.T) {
	jh := jh.New()
	test.Process(jh, t)
}

func Example() {
	var msg = make(chan string)

	jh := jh.New()
	jh.Confirm(func(j job.Job) {
		msg <- fmt.Sprintln("Confirmed", j)
	})

	jh.Submit("Job 1")
	jh.Submit("Job 2")
	jh.Submit("Job 3")

	fmt.Print(<-msg)
	fmt.Print(<-msg)
	fmt.Print(<-msg)

	// Unordered output:
	// Confirmed Job 1
	// Confirmed Job 2
	// Confirmed Job 3
}
