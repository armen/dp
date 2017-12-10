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
