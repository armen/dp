package transformation

import (
	"github.com/armen/irdp/job"
)

type tfmHandler struct {
	jh job.Handler

	buffer   []*job.Job // Buffer
	bound    int
	top      int
	bottom   int
	handling bool

	confirm func(*job.Job) // Confirm handler
	error   func(*job.Job) // Error handler

	// A multiplexer which is used to run events in a mutually exclusive way
	mux chan func()
}

// Instantiates a new asynchronous job handler.
func New(jh job.Handler, bound int) *tfmHandler {
	th := &tfmHandler{
		jh: jh,

		buffer:   make([]*job.Job, bound),
		bound:    bound,
		top:      0,
		bottom:   0,
		handling: false,

		mux: make(chan func()),
	}

	jh.Confirm(th.jhConfirm)

	return th
}

// Confirm registers the confirm handler.
func (th *tfmHandler) Confirm(f func(*job.Job)) {
	th.confirm = f
}

// Error registers the error handler.
func (th *tfmHandler) Error(f func(*job.Job)) {
	th.error = f
}

// Process registers the process handler.
func (th *tfmHandler) Process(f func(*job.Job)) {
	th.jh.Process(f)
}

// React mutually executes events.
func (th *tfmHandler) React() {
	go th.jh.React()

	for f := range th.mux {
		f()

		th.existsJob()
	}
}
