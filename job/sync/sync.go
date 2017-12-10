// Package sync implements job.Handler interface
package sync

import (
	"github.com/armen/irdp/job"
)

// Submits a job to be processed.
func (jh *jobHandler) Submit(j *job.Job) {
	jh.mux <- func() {
		jh.process(j)

		go jh.confirm(j)
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
