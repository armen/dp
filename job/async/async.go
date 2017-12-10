// Package async implements job.Handler interface
package async

import (
	"github.com/armen/irdp/job"
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
