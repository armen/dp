// Package async implements job.Handler interface
package async

import (
	"irdp/job"
)

// Submits a job to be processed.
func (jh *jobHandler) Submit(j *job.Job) {
	jh.mux <- func() {
		jh.buffer = append(jh.buffer, j)

		go jh.confirm(j)
	}
}

func (jh *jobHandler) notEmptyBuffer() {
	jh.mux <- func() {
		// Select a job
		j := jh.buffer[0]

		jh.process(j)
		jh.buffer = jh.buffer[1:]
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
