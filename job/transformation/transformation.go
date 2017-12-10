// Package transformation implements Job-Transformation by buffering
package transformation

import (
	"irdp/job"
)

// Submits a job to be processed.
func (th *tfmHandler) Submit(j *job.Job) {
	th.mux <- func() {
		if th.bottom+th.bound == th.top {
			go th.error(j)

			return
		}
		th.buffer[th.top%th.bound] = j
		th.top++
		go th.confirm(j)
	}
}

func (th *tfmHandler) handleJob() {
	th.mux <- func() {
		j := th.buffer[th.bottom%th.bound]
		th.bottom++
		th.handling = true
		go th.jh.Submit(j)
	}
}

func (th *tfmHandler) jhConfirm(j *job.Job) {
	th.mux <- func() {
		th.handling = false
	}
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
