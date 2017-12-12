package async

import (
	"github.com/armen/dp/job"
)

// JobHandler stores the state of an asynchronous job handler.
type JobHandler struct {
	confirm func(*job.Job) // Confirm handler
	process func(*job.Job) // Process handler, to process a job
	buffer  []*job.Job     // Buffer, to buffer a job

	// A multiplexer which is used to run events in a mutually exclusive way
	mux chan func()
}

// New instantiates a new asynchronous job handler.
func New() *JobHandler {
	return &JobHandler{
		process: func(*job.Job) {},
		mux:     make(chan func()),
	}
}

// Confirm registers the confirm handler.
func (jh *JobHandler) Confirm(f func(*job.Job)) {
	jh.confirm = f
}

// Process registers the process handler.
func (jh *JobHandler) Process(f func(*job.Job)) {
	jh.process = f
}

// React mutually executes events.
func (jh *JobHandler) React() {
	jh.init()

	for f := range jh.mux {
		f()

		jh.existsJob()
	}
}
