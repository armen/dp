// Package sync implements job.Handler interface
package sync

import (
	"github.com/armen/dp/job"
)

// Submit submits a job to be processed.
func (jh *JobHandler) Submit(j job.Job) {
	jh.mux <- func() {
		jh.process(j)

		go jh.confirm(j)
	}
}
