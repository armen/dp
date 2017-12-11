package async

import (
	"github.com/armen/dp/job"
)

type jobHandler struct {
	confirm func(*job.Job) // Confirm handler
	process func(*job.Job) // Process handler, to process a job
	buffer  []*job.Job     // Buffer, to buffer a job

	// A multiplexer which is used to run events in a mutually exclusive way
	mux chan func()
}

// Instantiates a new asynchronous job handler.
func New() *jobHandler {
	return &jobHandler{
		process: func(*job.Job) {},
		mux:     make(chan func()),
	}
}

// Confirm registers the confirm handler.
func (jh *jobHandler) Confirm(f func(*job.Job)) {
	jh.confirm = f
}

// Process registers the process handler.
func (jh *jobHandler) Process(f func(*job.Job)) {
	jh.process = f
}

// React mutually executes events.
func (jh *jobHandler) React() {
	jh.init()

	for f := range jh.mux {
		f()

		jh.existsJob()
	}
}
