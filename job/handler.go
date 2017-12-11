package job

import (
	"github.com/armen/dp"
)

// Interface and properties of a job handler.
//
// Properties:
// 	JH1: Guarnteed response
// 		- Every submitted job is eventually confirmed
//
type Handler interface {
	Submit(*Job)        // Requests a job to be processed
	Confirm(func(*Job)) // Confirms that the given job has been (or will be) processed
	Process(func(*Job)) // Processes a job

	dp.Reactor
}
