package sync

import (
	"irdp/job"
)

type jobHandler struct {
	confirm func(*job.Job) // Confirm handler
	process func(*job.Job) // Process handler, to process a job

	// A multiplexer which is used to run events in a mutually exclusive way
	mux chan func()
}

// Instantiates a new synchronous job handler.
func New() *jobHandler {
	return &jobHandler{
		process: func(*job.Job) {},
		mux:     make(chan func()),
	}
}

// React mutually executes events.
func (jh *jobHandler) React() {
	for f := range jh.mux {
		f()
	}
}
