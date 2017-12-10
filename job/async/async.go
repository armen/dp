// Package async implements job.Handler interface
package async

import (
	"github.com/armen/irdp/job"
)

func (jh *jobHandler) init() {
	jh.buffer = make([]*job.Job, 0)
}

// Submits a job to be processed.
func (jh *jobHandler) Submit(j *job.Job) {
	jh.mux <- func() {
		jh.buffer = append(jh.buffer, j)

		go jh.confirm(j)
	}
}

func (jh *jobHandler) existsJob() {
	if len(jh.buffer) > 0 {
		go jh.handleJob()
	}
}

func (jh *jobHandler) handleJob() {
	jh.mux <- func() {
		// Select a job
		j := jh.buffer[0]

		jh.process(j)
		jh.buffer = jh.buffer[1:]
	}
}
