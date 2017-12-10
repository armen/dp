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

// React mutually executes events.
func (th *tfmHandler) React() {
	go th.jh.React()

	for f := range th.mux {
		f()

		if th.bottom < th.top && th.handling == false {
			go th.handleJob()
		}
	}
}
