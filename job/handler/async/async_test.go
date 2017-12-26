package async_test

import (
	"fmt"
	"testing"

	"github.com/armen/dp/job"
	"github.com/armen/dp/job/handler/async"
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

func Example() {
	var msg = make(chan string)

	jh := async.New()
	jh.Confirm(func(j job.Job) {
		msg <- fmt.Sprintln("Confirmed", j)
	})
	go jh.React()

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
