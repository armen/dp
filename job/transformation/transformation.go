// Package transformation implements Job-Transformation by buffering
package transformation

import (
	"github.com/armen/dp/job"
)

func (th *TfmHandler) init() {
	th.top = 0
	th.bottom = 0
	th.handling = false
	th.buffer = make([]*job.Job, th.bound)
}

// Submit submits a job to be processed.
func (th *TfmHandler) Submit(j *job.Job) {
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

func (th *TfmHandler) existsJob() {
	if th.bottom < th.top && th.handling == false {
		go th.handleJob()
	}
}

func (th *TfmHandler) handleJob() {
	th.mux <- func() {
		j := th.buffer[th.bottom%th.bound]
		th.bottom++
		th.handling = true
		go th.jh.Submit(j)
	}
}

func (th *TfmHandler) jhConfirm(j *job.Job) {
	th.mux <- func() {
		th.handling = false
	}
}
