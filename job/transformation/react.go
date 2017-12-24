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
		jh:    jh,
		bound: bound,
		mux:   make(chan func()),
	}

	jh.Confirm(th.jhConfirm)

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

// React mutually executes events.
func (th *TfmHandler) React() {
	th.init()
	go th.jh.React()

	for f := range th.mux {
		f()
	}
}
