package job

import (
	"irdp"
)

// Interface and properties of a job transformation and processing abstraction.
//
// Properties:
// 		TH1: Guarnteed response
// 			- Every submitted job is eventually confirmed or its transformation fails
//      TH2: Soundness
// 			- A submitted job whose transformation fails is not processed
//
type TransformationHandler interface {
	Submit(*Job)        // Requests a job for transformation and for processing
	Confirm(func(*Job)) // Confirms that the given job has been (or will be) transformed and processed
	Error(func(*Job))   // Indicates that the transformation of the given job failed
	Process(func(*Job)) // Processes a job

	irdp.Reactor
}
