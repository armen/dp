package transformation

import (
	"github.com/armen/dp/job"
)

// TfmHandler stores the state of job-transformation handler.
type TfmHandler struct {
	jh job.Handler

	buffer   []job.Job // Buffer
	bound    int
	top      int
	bottom   int
	handling bool

	confirm func(job.Job) // Confirm handler
	error   func(job.Job) // Error handler

	// A multiplexer which is used to run events in a mutually exclusive way
	mux chan func()
}

// New instantiates a new asynchronous job handler.
func New(jh job.Handler, bound int) *TfmHandler {
	th := &TfmHandler{
		jh:      jh,
		bound:   bound,
		confirm: func(job.Job) {},
		error:   func(job.Job) {},
		mux:     make(chan func()),
	}

	th.init()
	go th.react()

	return th
}

// Confirm registers the confirm handler.
func (th *TfmHandler) Confirm(f func(job.Job)) {
	th.confirm = f
}

// Error registers the error handler.
func (th *TfmHandler) Error(f func(job.Job)) {
	th.error = f
}

// Process registers the process handler.
func (th *TfmHandler) Process(f func(job.Job)) {
	th.jh.Process(f)
}

// react mutually executes events.
func (th *TfmHandler) react() {
	for f := range th.mux {
		f()
	}
}
