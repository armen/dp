package sync

import (
	"github.com/armen/dp/job"
)

// JobHandler stores the state of a synchronous job handler.
type JobHandler struct {
	confirm func(job.Job) // Confirm handler
	process func(job.Job) // Process handler, to process a job

	// A multiplexer which is used to run events in a mutually exclusive way
	mux chan func()
}

// New instantiates a new synchronous job handler.
func New() *JobHandler {
	jh := &JobHandler{
		confirm: func(job.Job) {},
		process: func(job.Job) {},
		mux:     make(chan func()),
	}

	go jh.react()

	return jh
}

// Confirm registers the confirm handler.
func (jh *JobHandler) Confirm(f func(job.Job)) {
	jh.confirm = f
}

// Process registers the process handler.
func (jh *JobHandler) Process(f func(job.Job)) {
	jh.process = f
}

// react mutually executes events.
func (jh *JobHandler) react() {
	for f := range jh.mux {
		f()
	}
}
