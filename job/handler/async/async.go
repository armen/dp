// Package async implements job.Handler interface
package async

import (
	"github.com/armen/dp/job"
)

func (jh *JobHandler) init() {
	jh.buffer = make([]job.Job, 0)
}

// Submit submits a job to be processed.
func (jh *JobHandler) Submit(j job.Job) {
	jh.mux <- func() {
		jh.buffer = append(jh.buffer, j)

		go jh.confirm(j)
	}
}

func (jh *JobHandler) existsJob() {
	if len(jh.buffer) > 0 {
		go jh.handleJob()
	}
}

func (jh *JobHandler) handleJob() {
	jh.mux <- func() {
		// Select a job
		j := jh.buffer[0]

		jh.process(j)
		jh.buffer = jh.buffer[1:]
	}
}
